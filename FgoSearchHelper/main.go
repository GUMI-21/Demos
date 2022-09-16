package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var configFile = flag.String("C", "./conf/fgoSearchHelper.toml", "config file")

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix()) // 随机初始化

	fmt.Println("test")
}
