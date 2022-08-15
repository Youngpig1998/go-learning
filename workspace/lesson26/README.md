# 上下文Context

​	Context主要用来在Goroutine之间传递上下文信息，包括：取消信号，超时时间，截止时间，以及一些key-value pair等。



## 将context传入函数中使用

​	Go官方建议我们把context作为函数的第一个参数使用，命名为ctx。

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // 5 seconds pass
		fmt.Println("finish doing something")
	case <-ctx.Done(): // ctx is cancelled
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func main() {
	// 创建空context的两种方法
	ctx := context.Background() // 返回一个空的context，不能被cancel，kv为空

	// todoCtx := context.TODO() // 和Background类似，当你不确定的时候使用
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	doSomething(ctx)

}

```



## 监听Context是否被取消

`Context.Done()`返回一个只读的通道，当context被取消时，这个通道里会传入一个值表示它已经被关闭了。因此，我们只需要用select来监听`context.Done()`是否有值，就可以判断context是否被取消。







## 缺陷

​	如果函数接收 context 参数，确保检查它是如何处理取消通知的。例如，`exec.CommandContext`不会关闭读取管道，直到命令执行了进程创建的所有分支（Github 问题：https://github.com/golang/go/issues/23019 ），这意味着如果等待 `cmd.Wait()` 直到外部命令的所有分支都已完成，则 context 取消不会使该函数立即返回。如果您使用超时或截止日期，您可能会发现这不能按预期运行。如果遇到任何此类问题，可以使用 time.After 实现超时。

## 最佳实践

1. context.Background 只应用在最高等级，作为所有派生 context 的根。
2. context.TODO 应用在不确定要使用什么的地方，或者当前函数以后会更新以便使用 context。
3. context 取消是建议性的，这些函数可能需要一些时间来清理和退出。
4. context.Value 应该很少使用，它不应该被用来传递可选参数。这使得 API 隐式的并且可以引起错误。取而代之的是，这些值应该作为参数传递。
5. 不要将 context 存储在结构中，在函数中显式传递它们，最好是作为第一个参数。
6. 永远不要传递不存在的 context 。相反，如果您不确定使用什么，使用一个 ToDo context。
7. Context 结构没有取消方法，因为只有派生 context 的函数才应该取消 context。

