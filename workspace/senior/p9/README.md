# Singleton Pattern

​	单例模式单例模式（Singleton Pattern），是**最简单的一个模式**。在 Go 中，单例模式指的是全局只有一个实例，并且它负责创建自己的对象。单例模式不仅有利于减少内存开支，还有减少系统性能开销、防止多个实例产生冲突等优点。

​	因为单例模式保证了实例的全局唯一性，而且只被初始化一次，所以比较适合**全局共享一个实例，且只需要被初始化一次的场景**，例如数据库实例、全局配置、全局任务池等。

​	单例模式又分为**饿汉方式**和**懒汉方式**。饿汉方式指全局的单例实例在包被加载时创建，而懒汉方式指全局的单例实例在第一次被使用时创建。你可以看到，这种命名方式非常形象地体现了它们不同的特点。

​	接下来，我就来分别介绍下这两种方式。先来看**饿汉方式**。

​	下面是一个饿汉方式的单例模式代码：

```go
package singleton

type singleton struct {
}

var ins *singleton = &singleton{}

func GetInsOr() *singleton {
    return ins
}
```

​	你需要注意，因为实例是在包被导入时初始化的，所以如果初始化耗时，会导致程序加载时间比较长。

​	**懒汉方式是开源项目中使用最多的**，但它的缺点是非并发安全，在实际使用时需要加锁。以下是懒汉方式不加锁的一个实现：

```go
package singleton

type singleton struct {
}

var ins *singleton

func GetInsOr() *singleton {
    if ins == nil {
        ins = &singleton{}
    }
    
    return ins
}
```

​	可以看到，在创建 ins 时，如果 ins==nil，就会再创建一个 ins 实例，这时候单例就会有多个实例。为了解决懒汉方式非并发安全的问题，需要对实例进行加锁，下面是带检查锁的一个实现：

```go
import "sync"

type singleton struct {
}

var ins *singleton
var mu sync.Mutex

func GetIns() *singleton {
  if ins == nil {
    mu.Lock()
    if ins == nil {
      ins = &singleton{}
    }
        mu.Unlock()
  }
  return ins
}
```

​	上述代码只有在创建时才会加锁，既提高了代码效率，又保证了并发安全。除了饿汉方式和懒汉方式，在 Go 开发中，还有一种更优雅的实现方式，我建议你采用这种方式，代码如下：

```go
package singleton

import (
    "sync"
)

type singleton struct {
}

var ins *singleton
var once sync.Once

func GetInsOr() *singleton {
    once.Do(func() {
        ins = &singleton{}
    })
    return ins
}
```

​	使用once.Do可以确保 ins 实例全局只被创建一次，once.Do 函数还可以确保当同时有多个创建动作时，只有一个创建动作在被执行。另外，IAM 应用中大量使用了单例模式，如果你想了解更多单例模式的使用方式，可以直接查看 IAM 项目代码。

​	IAM 中单例模式有 [GetStoreInsOr](https://github.com/colin404test/iam/blob/IAMTAG/internal/authzserver/store/store.go#L45)、[GetEtcdFactoryOr](https://github.com/colin404test/iam/blob/IAMTAG/internal/apiserver/store/etcd/etcd.go#L83)等。



------

Singleton creational design pattern restricts the instantiation of a type to a single object.

## Implementation

```go
package singleton

type singleton map[string]string

var (
    once sync.Once

    instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}
```

## Usage

```go
s := singleton.New()

s["this"] = "that"

s2 := singleton.New()

fmt.Println("This is ", s2["this"])
// This is that
```

## Rules of Thumb

- Singleton pattern represents a global state and most of the time reduces testability.
