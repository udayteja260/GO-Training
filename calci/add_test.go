package calci

import "testing"

func TestAdd(t *testing.T) {
	i := Inputs{Input1: 1, Input2: 5}
	got := i.Add()
	want := 6

	if got != want {
		t.Errorf("Sum(1,5) = %d; want %d", got, want)
	}

}
