package main

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//Pool管理一组可以安全的在多个goroutine间共享的资源, 被管理的资源必须实现io.Closer接口
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

//创建一个用来管理资源的池
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size too small")
	}
	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil
}

//从池中获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	//检查是否有空闲的资源
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
		//没有空闲资源, 生成新资源
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

//释放资源
func (p *Pool) Release(r io.Closer) {
	//保证本操作和close操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	//如果池已经关闭, 销毁这个资源
	if p.closed {
		r.Close()
		return
	}

	select {
	//试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:", "In Queue")

		//如果队列满了, 则关闭这个资源
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

//让资源池停止工作, 释放内部所有资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	//如果pool已经关闭了, 就什么都不做
	if p.closed {
		return
	}

	p.closed = true
	//清空管道里的资源, 如果不这么做会发生死锁
	close(p.resources) //关闭管道
	for r := range p.resources {
		r.Close()
	}
}

const (
	maxGoroutines   = 25 //要是用的goroutine
	pooledResources = 2  //池中的资源数量
)

type dbConnection struct {
	ID int32
}

var idCounter int32

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)
	p, err := New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
		return
	}
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown Program")
	p.Close()
}

func performQueries(query int, p *Pool) {
	conn, err := p.Acquire()
	defer p.Release(conn)
	if err != nil {
		log.Println(err)
		return
	}
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create:New Connection", id)
	return &dbConnection{
		ID: id,
	}, nil
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close:Connection", dbConn.ID)
	return nil
}
