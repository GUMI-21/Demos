package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	A, B int
}

type Rect struct{}

func (r *Rect) Multi(p Params, ret *int) error {
	*ret = p.A * p.B
	return nil
}

func (r *Rect) Division(p Params, ret *int) error {
	*ret = p.A / p.B
	return nil
}

func (r *Rect) Remainder(p Params, ret *int) error {
	*ret = p.A % p.B
	return nil
}

func main() {
	rect := new(Rect)
	rpc.Register(rect)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln(err)
	}

}
