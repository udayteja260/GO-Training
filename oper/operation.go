package oper

import (
	"fmt"
)

type Movie struct {
	Title string
	Hero  string
	year  int
}

func Operation1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	x := 15
	if x%5 == 0 {
		fmt.Println("x is multiple of 5")
	} else {
		fmt.Println("x is not multiple of 5")
	}
	movie := "BhagavanthKesari"
	switch movie {
	case "Leo":
		fmt.Println("Block Buster")
	case "BhagavanthKesari":
		fmt.Println("BlockBuster")
	case "Tiger Nageswara Rao":
		fmt.Println("Average")
	default:
		fmt.Println("Flop")
	}
	defer fmt.Println("Prints last")
	fmt.Println("Prints First")
	q := 36
	p := &q
	fmt.Println(*p)
	movie1 := Movie{
		Title: "Bahubali",
		Hero:  "Prabhas",
		year:  2017,
	}
	fmt.Println(movie1.Title, "is", movie1.Hero, "Movie Released in year", movie1.year)
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[0])
}
