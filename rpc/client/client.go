package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

func main() {
	//连接rpc
	conn, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	ret := 0
	// 面积
	err = conn.Call("Rect.Area", Params{Width: 50, Height: 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("面积", ret)
	// 周长
	err = conn.Call("Rect.Perimeter", Params{Width: 50, Height: 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("周长", ret)
}
