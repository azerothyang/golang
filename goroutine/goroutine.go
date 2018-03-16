package main

import (
	"./policy"
	"sync"
	"log"
	"os"
	"fmt"
)

//开5个协程获取接口
const goNumber  = 1

func main() {
	retData, err := policy.Request()
	fmt.Println(err, retData)
	os.Exit(1)
	var waitGroup sync.WaitGroup
	var chRetData chan *policy.RetData
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
	for data := range chRetData {
		log.Printf("%s: \n", data.Message)
	}
}
