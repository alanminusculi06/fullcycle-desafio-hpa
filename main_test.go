package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	number := 4.0
	res := Calc(number)
	if res != 2 {
		t.Errorf("Sqrt(4) = %f; expected 2", res)
	}
}
