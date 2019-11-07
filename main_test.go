package main

import (
	"fmt"
	"testing"
)
func TestSum(t *testing.T)  {
	var (
		a = 2
		b = 3
	)

	res := sum(a, b)
	if res != 5 {
		fmt.Errorf("fail sum function")
	}
}
