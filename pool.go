package main

import (
	"fmt"
	"time"
)

// goroutine pool

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{
		f:f,
	}
}

func(t *Task)Execute() {
	t.f()
}

type Pool struct {
	EntryChannel chan *Task
	worker_num int
	JobsChannel chan *Task
}

func NewPool(caption int) *Pool {
	return &Pool{
		EntryChannel:make(chan *Task),
		worker_num:caption,
		JobsChannel:make(chan *Task),
	}
}

func(p *Pool)worker(id int) {
	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("工作id ", id)
	}
}

func(p *Pool)Run() {
	for i:=0; i < p.worker_num; i++ {
		go p.worker(i)
	}

	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
	close(p.JobsChannel)
	close(p.EntryChannel)
}

func main() {
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})
	p := NewPool(3)
	go func() {
		for {
			p.EntryChannel <- t
		}
	}()
	p.Run()
}
