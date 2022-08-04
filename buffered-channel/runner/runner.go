package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var (
	ErrTimeout   = errors.New("cannot finish tasks within the timeout")
	ErrInterrupt = errors.New("erceived interrupt from OS")
)

// give some tasks, the Runner need to finish them in time, or an ErrTimeout will be triggerd
type Runner struct {
	tasks     []func(int) // list of tasks
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time // used for timing
}

func New(t time.Duration) *Runner {
	return &Runner{
		tasks:     make([]func(int), 0),
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(t), // relay the channel of time when time is out
	}
}

func (r *Runner) AddTasks(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) run() error {

	//range the slice loop
	for id, task := range r.tasks {
		select {
		case <-r.interrupt: // if there is something in r.interrupt, go here
			signal.Stop(r.interrupt)
			return ErrInterrupt
		default: // else if r.interrupt is empty, go here
			task(id)
		}
	}

	return nil
}

func (r *Runner) Start() error {

	//if there is a signal of Interrupt occurs, r.interrupt will relay interrupt from OS, equivalent to the assignment operation
	signal.Notify(r.interrupt, os.Interrupt)

	//run the tasks
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout

	}
}
