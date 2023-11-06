package assignment3

import (
	"fmt"
)

func Sli() {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:3]
	fmt.Println(s1)

	s2 := []int{2, 4, 6, 8, 10, 12}
	s2 = s2[2:5]
	fmt.Println(s2)
	s2 = s2[:3]
	fmt.Println(s2)

	n := []string{"a", "b", "c", "d"}
	x := n[0:2]
	y := n[3:4]
	fmt.Println(x, y)

	x[1] = "z"
	fmt.Println(x)
	fmt.Println(n)

	g := map[string]float64{
		"a": 90.0,
		"b": 91.1,
		"c": 92.2,
	}

	g["d"] = 93.3

	delete(g, "d")
	fmt.Println(g)

	for student, grade := range g {
		fmt.Println(student, grade)
	}

	if grade, ok := g["b"]; ok {
		fmt.Println(grade, ok)
	} else {
		fmt.Println("invalid")
	}

	e := make(map[string]int)

	e["x"] = 61
	e["y"] = 62
	e["z"] = 63

	for student, id := range e {
		fmt.Println(student, id)
	}

}
