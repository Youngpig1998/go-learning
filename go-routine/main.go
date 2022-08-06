package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func say(id string) {
	time.Sleep(time.Second)

	fmt.Println("I am done! id: " + id)

	wg.Done()
}

func player(name string, ch chan int) {
	defer wg.Done()

	for {
		ball, ok := <-ch // fetch the value from channel

		if !ok {
			fmt.Printf("channel is closed! %s wins!\n", name)
			return
		}

		n := rand.Intn(100)

		if n%10 == 0 {
			close(ch)
			fmt.Printf("%s misses the ball! %s loses!\n", name, name)
			return
		}

		ball++
		fmt.Printf("%s receives ball %d\n", name, ball)
		ch <- ball

		
	}
}

var wg sync.WaitGroup

func main() {

	wg.Add(4) //  2 jobs to finish

	//go func(id string) {
	//	fmt.Println(id)
	//	wg.Done()
	//}("hello")
	//go say("world")

	//channel
	ch := make(chan int, 0) // unbuffered channel

	//ch <- 0 //This will cause the deadlock in main, because nobody will fetch the value from the channel

	go player("wang", ch)

	go player("chong", ch)

	ch <- 0

	raceCondition()

	wg.Wait() // wait for all jobs to be finished

	fmt.Println(counter)
	fmt.Println(<-ch2)
}

var counter int32
var mtx sync.Mutex // method one
var ch2 = make(chan int, 1)

func MutexIncCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		mtx.Lock()
		counter++
		mtx.Unlock()
	}
}

func UnsafeIncCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		counter++
	}
}

func AtomicIncCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&counter, 1)
	}
}

func ChannelIncCounter() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		count := <-ch2
		count++
		ch2 <- count
	}
}

func raceCondition() {

	//race condition
	//go UnsafeIncCounter()
	//go UnsafeIncCounter()

	//go MutexIncCounter()
	//go MutexIncCounter()

	//go AtomicIncCounter()
	//go AtomicIncCounter()

	ch2 <- 0

	go ChannelIncCounter()
	go ChannelIncCounter()
}
