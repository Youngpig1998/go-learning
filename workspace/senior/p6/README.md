# stop channel

​	当在执行某个goroutine的同时，需要等待该goroutine结束后再执行下面的逻辑代码，但我们又需要同时执行一些特定操作时，可以考虑使用for select范式。

​	然而，这样一来大量的fo select语句会充斥在我们的业务代码中。所以，我们可以考虑使用stopCh

​	我们可以定义一个BackOffUntil函数，该函数接收两个参数，一个名为stopCh的空结构体channel和一个名为fn的匿名函数。stopCh仅仅只是用来传递信号，所以我们选用[空结构体](../p5)，而fn则是在等待期间可以做的一些业务逻辑。

​	有了这个BackOffUntil函数，我们就可以将这些共性的代码抽离出来，写入框架层，业务代码只需要在合适的时候调用BackOffUntil函数即可。

```go
package main

import (
	"fmt"
	"time"
)


// BackOffUntil：  goroutine 启停函数           fn的函数参数可以根据需求进行改变
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


//定义一个公用的业务逻辑，在BackOffUntil中执行
func fn()  {
	fmt.Println("fn run")
}


func main() {



	stopCh := make(chan struct{})

	//模拟3s后停止
	go func(chan struct{}) {
		select {
		case <-time.After(3 * time.Second):
			close(stopCh)
			fmt.Println("stopCh closed")
		}
	}(stopCh)

	//会阻塞
	BackOffUntil(stopCh,fn)


	stopCh2 := make(chan struct{})

	//模拟5s后停止
	go func(chan struct{}) {
		select {
		case <-time.After(5 * time.Second):
			close(stopCh2)
			fmt.Println("stopCh closed")
		}
	}(stopCh2)


	BackOffUntil(stopCh2,fn)
}
```

