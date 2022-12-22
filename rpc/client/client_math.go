package main

import (
	"fmt"
	"github.com/siddontang/go/log"
	"net/rpc"
)

type Params struct {
	A, B int
}

func main() {
	conn, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Error(err)
	}

	ret := 0
	err = conn.Call("Rect.Multi", Params{A: 4, B: 3}, &ret)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(ret)

	err = conn.Call("Rect.Division", Params{4, 3}, &ret)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(ret)

	err = conn.Call("Rect.Remainder", Params{4, 3}, &ret)
	if err != nil {
		log.Error(err)
	}
	fmt.Println(ret)
}
