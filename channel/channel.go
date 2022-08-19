package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGroutines = 4
	taskLoad        = 10
)

var wg1 sync.WaitGroup

func main() {
	rand.Seed(time.Now().Unix())
	tasks := make(chan string, taskLoad)
	wg1.Add(numberGroutines)
	for gr := 1; gr <= numberGroutines; gr++ {
		go worker(tasks, gr)
	}
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task:%d", post)
	}
	close(tasks)
	wg1.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg1.Done()
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("worker:%d:shutting down\n", worker)
			return
		}
		//显示我们要开始工作了
		fmt.Printf("Worker:%d:Started %s\n", worker, task)
		//随机等待一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		//显示工作完成
		fmt.Printf("Worker: %d : Completed %s \n", worker, task)
	}
}
