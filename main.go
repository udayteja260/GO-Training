package main

import (
	// "assignment2/oper"
	// "assignment2/assignment3"
	"assignment2/calci"
	"fmt"
)

func main() {
	// oper.Operation1()
	// assignment3.Sli()
	i1 := calci.Inputs{Input1: 1, Input2: 6}
	fmt.Println(i1.Add())
}
