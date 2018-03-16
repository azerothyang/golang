package main

import (
	"./policy"
	"sync"
	"log"
)

//开5个协程获取接口
const goNumber  = 1000

func main() {
	var waitGroup sync.WaitGroup
	chRetData := make(chan *policy.RetData) //这里注意需要make创建, 不能只声明, 不初始化
	waitGroup.Add(goNumber)
	for i:=0; i<goNumber; i++ {
		//启动请求协程
		go func() {
			policy.RequestAndAnalyze(chRetData)
			waitGroup.Done()
		}()
	}

	////启动一个协程等待完成任务的监控协程, 完成后关闭通道
	go func() {
		waitGroup.Wait()
		close(chRetData)
	}()

	Display(chRetData)
}

func Display(chRetData chan *policy.RetData){
	count := 1
	for data := range chRetData {
		log.Printf("当前第%d次, 结果为: %d \n", count, data.Data.Total)
		count++
	}
}
