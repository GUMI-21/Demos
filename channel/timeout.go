package main

import (
	"fmt"
	"runtime"
	"time"
)

func doThing(t chan bool) {
	time.Sleep(time.Second)
	select {
	case t <- true:
	default:
		return
	}
}

// time.After实现超时
func timeoutCheck(f func(chan bool)) {
	done := make(chan bool)
	go f(done)
	select {
	case <-done:
		fmt.Println("success")
	case <-time.After(time.Millisecond):
		fmt.Println("timeout")
	}
}

func main() {
	for i := 0; i <= 1000; i++ {
		timeoutCheck(doThing)
	}
	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}
