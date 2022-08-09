package main

import (
	"fmt"
	"time"
)

// BackOffUntil goroutine 启停函数
func BackOffUntil2(stopCh chan struct{}, fn func()) {
	//经典 for select范式
	for {
		select {
		// 只有stopCh被close掉，才会读到值
		case <-stopCh:
			return
		//在此处设置了一个计时器，意思每过一秒就做一些事情
		case <-time.After(1 * time.Second):
			fmt.Println("begin running fn()")
			//可以在这里调用一些函数，根据自己的需求对函数做更改
			fn()
		}
	}
}

//定义一个公用的业务逻辑，在BackOffUntil中执行
func fn2()  {
	fmt.Println("fn2 run")
}

func main() {
	stopCh := make(chan struct{},2)


	//模拟3s后停止
	go func(chan struct{}) {
		select {
		case <-time.After(3 * time.Second):
			stopCh<- struct{}{}
			fmt.Println("stopCh receives a empty struct")
		}
	}(stopCh)

	BackOffUntil2(stopCh,fn2)





	//模拟10s后停止
	go func(chan struct{}) {
		select {
		case <-time.After(10 * time.Second):
			stopCh<- struct{}{}
			fmt.Println("stopCh receives a empty struct")
		}
	}(stopCh)


	BackOffUntil2(stopCh,fn2)
}