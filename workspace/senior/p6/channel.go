package main

import (
	"fmt"
	"time"
)

// BackOffUntil goroutine 启停函数
func BackOffUntil(stopCh chan struct{}, fn func()) {
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

func main() {
	stopCh := make(chan struct{})

	//定义一个业务逻辑
	fn := func() {
		fmt.Println("fn run")
	}

	//模拟3s后停止
	go func(chan struct{}) {
		select {
		case <-time.After(3 * time.Second):
			close(stopCh)
			fmt.Println("stopCh closed")
		}
	}(stopCh)


	BackOffUntil(stopCh,fn)

	//stopCh2 := make(chan struct{})
	//
	////模拟10s后停止
	//go func(chan struct{}) {
	//	select {
	//	case <-time.After(10 * time.Second):
	//		close(stopCh2)
	//		fmt.Println("stopCh closed")
	//	}
	//}(stopCh2)
	//
	//
	//BackOffUntil(stopCh2,fn)
}