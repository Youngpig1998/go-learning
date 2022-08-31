# Learning notes for golang

​	此项目为Go语言学习中的知识总结、项目源码等，未完待续。

## 基础篇

* [lesson1: Go程序结构](./workspace/lesson1)

* [lesson2: 变量与常量](./workspace/lesson2)

* [lesson3: 数据类型：数字，字符串，bool](./workspace/lesson3)

* [lesson4: 指针pointer](./workspace/lesson4)

* [lesson5: 运算操作符](./workspace/lesson5)

* [lesson6: 控制语句if/switch](./workspace/lesson6)

* [lesson7: 循环语句for/goto/break/continue](./workspace/lesson7)

* [lesson8: 数组：一维数组和多维数组](./workspace/lesson8)

* [lesson9: 切片Slice](./workspace/lesson9)

* [lesson10: 映射Map](./workspace/lesson10)

* [lesson11: range迭代](./workspace/lesson11)

* [lesson12: 字符串String](./workspace/lesson12)

* [lesson13: 函数，闭包和方法](./workspace/lesson13)

* [lesson14: 变量作用域](./workspace/lesson14)

* [lesson15: defer语义](./workspace/lesson15)

* [lesson16: 结构体struct](./workspace/lesson16)

* [lesson17: 接口interface](./workspace/lesson17)

* [lesson18: 协程goroutine和管道channel](./workspace/lesson18)

* [lesson19: 并发编程之sync.WaitGroup](./workspace/lesson19)

* [lesson20:  并发编程之sync.Once](./workspace/lesson20)

* [lesson21:  并发编程之sync.Mutex和sync.RWMutex](./workspace/lesson21)

* [lesson22: 并发编程之sync.Cond](./workspace/lesson22)

* [lesson23: 并发编程之sync.Map](./workspace/lesson23)

* [lesson24: 并发编程之原子操作sync/atomic](./workspace/lesson24)

* [lesson25: 包Package和模块Module](./workspace/lesson25)

* [lesson26: 上下文Context](./workspace/lesson26)

* [lesson27: go单元测试](./workspace/lesson27)

* [lesson28: panic, recover运行期错误处理](./workspace/lesson28)

* [lesson29: select语义](./workspace/lesson29)

* [lesson30: 反射reflection](./workspace/lesson30)

* [lesson31: go相关命令](./workspace/lesson31)

  

## 进阶篇

- 常用关键字
  - [new和make的使用区别和最佳实践是什么？](./workspace/senior/p18)
  - [被defer的函数一定会执行么？](./workspace/senior/p3)


* 语言基础
  * [Go有引用变量和引用传递么？map,channel和slice作为函数参数是引用传递么？](./workspace/senior/p1)
  * [一文读懂Go空结构体及其使用场景](./workspace/senior/p5)
  * [一文读懂Go匿名结构体的使用场景](./workspace/senior/p15)
  * [Go语言中命名函数参数和命名函数返回值的注意事项](./workspace/senior/p16)
  * [golang协程与线程区别简要介绍](./workspace/senior/p22)
* 设计模式
  * Idioms
    * [Functional Options](./workspace/senior/p8)
  * Creational Patterns
    * [Singleton单例模式](./workspace/senior/p9)
    * [Factory Method & Abstract Factory](./workspace/senior/p10)
  * Behavioral Patterns
    * [Strategy](./workspace/senior/p14)
    * [Template](./workspace/senior/p12)
  * Structual Patterns
    * [Proxy](./workspace/senior/p13)

- Workspace mode工作区模式

  * [Go 1.18：工作区模式workspace mode简介](./workspace/senior/p21)

  * [Go 1.18：工作区模式最佳实践](./workspace/senior/p21/go1.18-workspace-best-practice.md)



## 最佳实践

- [stop channel](./workspace/senior/p6)
- [functional函数集合](./workspace/senior/p7)
- [全局变量global](./workspace/senior/p11)



## Go Quiz

1. [Go Quiz: 从Go面试题看slice的底层原理和注意事项](./workspace/senior/p2)
2. [Go Quiz: 从Go面试题搞懂slice range遍历的坑](./workspace/senior/p4)
3. [Go Quiz: 从Go面试题看channel的注意事项](./workspace/senior/p19)
4. [Go Quiz: 从Go面试题看channel在select场景下的注意事项](./workspace/senior/p20)



## 实战篇

### 代码规范

* [Go编码规范规范](./workspace/style)
* [Go项目目录结构](./workspace/senior/p17)



### Git

[Git使用手册](./workspace/git)



### Protobuf

[Protobuf](./workspace/protobuf)



### Docker/K8s

* [Docker教程](https://github.com/Youngpig1998/docker-tutorial)
* [Kubernetes教程](https://github.com/Youngpig1998/KuberneteCluster-built)

### Prometheus

- [Prometheus教程](https://www.qikqiak.com/k8s-book/docs/52.Prometheus%E5%9F%BA%E6%9C%AC%E4%BD%BF%E7%94%A8.html)
- [Prometheus入门实践](./workspace/prometheus)

### Go Operator

- [cp4d-audit-webhook-operator](https://github.com/Youngpig1998/cp4d-audit-webhook-operator)
- [HCA-Operator](https://github.com/Youngpig1998/HCA-Operator)

## Go Book

* [The Go Programming Language-Go语言圣经](http://www.gopl.io/)
* [Go语言高级编程-chai2010.gitbooks.io](https://chai2010.gitbooks.io/advanced-go-programming-book/content/)
* [Go语言设计与实现-draveness.me](https://draveness.me/golang/)
* [Go设计模式-Tamer Tas@google](https://github.com/tmrts/go-patterns)
* [深入解析Go-tiancaiamao.gitbooks.io](https://tiancaiamao.gitbooks.io/go-internals/content/zh/)
* [码农桃花源-qcrao91.gitbook.io](https://qcrao91.gitbook.io/go/)
* [Go语言高性能编程-geektutu](https://geektutu.com/post/hpg-benchmark.html)
* [Go Under The Hood-golang.design](https://golang.design/under-the-hood/)

## Go Video

### B站

* [神奇代码在哪里Go系列](https://space.bilibili.com/1557732/channel/collectiondetail?sid=464543)

* [Go最佳实践](https://space.bilibili.com/1897036286/channel/collectiondetail?sid=569368)

  


|      Topics      |              Source Code               |
| :--------------: | :------------------------------------: |
|     Packages     |         [get-bili](./get-bili)         |
|    go routine    |       [go-routine](./go-routine)       |
| buffered channel | [buffered channel](./buffered-channel) |

