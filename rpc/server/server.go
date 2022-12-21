package main

import (
	"log"
	"net/http"
	"net/rpc"
)

// golang实现rpc求矩形面积和周长

type Params struct {
	Width, Height int
}

type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

func main() {
	// 注册服务
	rect := new(Rect)
	// 注册
	rpc.Register(rect)
	rpc.HandleHTTP()
	// 监听服务
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln(err)
	}
}
