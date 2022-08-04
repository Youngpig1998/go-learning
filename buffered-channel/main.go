package main

import (
	"buffered-channel/pool"
	"buffered-channel/runner"
	"fmt"
	"io"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Printf("Task complete #%d\n", id)
	}
}

type DBConnection struct {
	id int32
}

func (D DBConnection) Close() error {
	fmt.Println("database closed, #" + fmt.Sprint(D.id))
	return nil
}

var counter int32

func Factory() (io.Closer, error) {
	atomic.AddInt32(&counter, 1)
	return DBConnection{id: counter}, nil
}

var wg sync.WaitGroup

func performQuery(query int, pool *pool.Pool) {
	defer wg.Done()

	resource, err := pool.AcquireResource()
	if err != nil {
		fmt.Println(err)
	}
	defer pool.ReleaseResource(resource)

	time.Sleep(time.Second)
	fmt.Println("finish query" + fmt.Sprint(query))
}

func main() {

	ch := make(chan int, 5)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	//For buffered channel, we have to close it first, then the range loop will be executed successfully
	close(ch)
	//We can fetch the value from the buffered channel even if the channel was closed
	for val := range ch {
		fmt.Println(val)
	}

	//test Runner
	r := runner.New(5 * time.Second)

	r.AddTasks(createTask(), createTask(), createTask())

	err := r.Start()
	switch err {
	case runner.ErrInterrupt:
		fmt.Println("tasks interrupted")
	case runner.ErrTimeout:
		fmt.Println("tasks timeout")
	default:
		fmt.Println("all tasks finished")

	}

	//test Pool
	p, err2 := pool.New(Factory, 5)
	if err2 != nil {
		log.Fatalln(err2)
	}

	num := 10
	wg.Add(num)

	for id := 0; id < num; id++ {
		go performQuery(id, p)
	}

	wg.Wait()

	p.Close()
}
