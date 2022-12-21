package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type ParamsMath struct {
	a, b int
}

type RectMath struct{}

func (r *RectMath) multi(p ParamsMath, ret *int) error {
	*ret = p.a * p.b
	return nil
}

func (r *RectMath) Division(p ParamsMath, ret *int) error {
	*ret = p.a / p.b
	return nil
}

func (r *RectMath) Remainder(p ParamsMath, ret *int) error {
	*ret = p.a / p.b
	return nil
}

func main() {
	rect := new(RectMath)
	rpc.Register(rect)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln(err)
	}

}
