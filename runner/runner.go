package main

import (
	"os"
	"time"
	"errors"
	"os/signal"
)

type Runner struct {
	//interrupt通道报告，从操作系统
	//发送的信号
	interrupt chan os.Signal

	//complete通道报告处理任务已经完成
	complete chan error

	//报告处理任务已经超时
	timeout <-chan time.Time

	//tasks持有一组以索引顺序依次执行的函数
	tasks []func(int)

}

//会在任务超时时会犯
var ErrTimeout = errors.New("received timeout")

//会在任务中断时返回
var ErrInterrupt = errors.New("received interrupt")

//NEW返回一个新的准备试用的runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete: make(chan error),
		timeout: time.After(d),
	}
}

func (r *Runner)Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner)Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {

	}
}

func (r *Runner)run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt(){
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

//验证是否收到中断信号
func (r *Runner)gotInterrupt()bool {
	select {
	//当中断事件被触发时发出的信号
	case<-r.interrupt:
		//停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
		//继续正常运行
	default:
		return false
	}

}



func main() {
}
