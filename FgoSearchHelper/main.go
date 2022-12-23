package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"runtime/debug"
	"time"

	"Demos/FgoSearchHelper/util"
)

var configFile = flag.String("C", "./conf/fgoSearchHelper.toml", "config file")

func main() {
	defer func() {
		if err := recover(); err != nil {
			msg := fmt.Sprintf(util.MsgPanic, "intention_server", err, debug.Stack(), util.GetHostName())
			fmt.Println(msg)
		}
	}()

	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().Unix()) // 随机初始化

}
