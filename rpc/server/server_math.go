package main

type ParamsMath struct {
	a, b int
}

type RectMath struct{}

func (r *RectMath) multi(p ParamsMath, ret *int) error {
	*ret = p.a * p.b
	return nil
}
