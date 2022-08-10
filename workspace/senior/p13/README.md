# 代理模式

​	代理模式 (Proxy Pattern)，可以为另一个对象提供一个替身或者占位符，以控制对这个对象的访问。以下代码是一个代理模式的实现：

```go
package proxy

import "fmt"

type Seller interface {
  sell(name string)
}

// 火车站
type Station struct {
  stock int //库存
}

func (station *Station) sell(name string) {
  if station.stock > 0 {
    station.stock--
    fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, station.stock)
  } else {
    fmt.Println("票已售空")
  }

}

// 火车代理点
type StationProxy struct {
  station *Station // 持有一个火车站对象
}

func (proxy *StationProxy) sell(name string) {
  if proxy.station.stock > 0 {
    proxy.station.stock--
    fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, proxy.station.stock)
  } else {
    fmt.Println("票已售空")
  }
}
```

​	上述代码中，StationProxy 代理了 Station，代理类中持有被代理类对象，并且和被代理类对象实现了同一接口。



------



The [proxy pattern](https://en.wikipedia.org/wiki/Proxy_pattern) provides an object that controls access to another object, intercepting all calls.

## Implementation

The proxy could interface to anything: a network connection, a large object in memory, a file, or some other resource that is expensive or impossible to duplicate.

Short idea of implementation:

```go
    // To use proxy and to object they must implement same methods
    type IObject interface {
        ObjDo(action string)
    }

    // Object represents real objects which proxy will delegate data
    type Object struct {
        action string
    }

    // ObjDo implements IObject interface and handel's all logic
    func (obj *Object) ObjDo(action string) {
        // Action behavior
        fmt.Printf("I can, %s", action)
    }

    // ProxyObject represents proxy object with intercepts actions
    type ProxyObject struct {
        object *Object
    }

    // ObjDo are implemented IObject and intercept action before send in real Object
    func (p *ProxyObject) ObjDo(action string) {
        if p.object == nil {
            p.object = new(Object)
        }
        if action == "Run" {
            p.object.ObjDo(action) // Prints: I can, Run
        }
    }
```

## Usage

More complex usage of proxy as example: User creates "Terminal" authorizes and PROXY send execution command to real Terminal object
See [proxy/main.go](proxy/main.go) or [view in the Playground](https://play.golang.org/p/mnjKCMaOVE).