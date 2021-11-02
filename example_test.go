package lev_test

import (
	"fmt"

	"github.com/meagar/lev-go"
)

func ExampleDistance() {
	a := "fire"
	b := "tires"
	fmt.Println(lev.Distance(a, b))
	// Output: 2
}

func ExampleDistanceD() {
	a := "ab"
	b := "ba"
	fmt.Println(lev.DistanceD(a, b))
	// Output: 1
}
