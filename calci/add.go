package calci

type Inputs struct {
	Input1 int
	Input2 int
}

func (i Inputs) Add() int {
	return i.Input1 + i.Input2
}
