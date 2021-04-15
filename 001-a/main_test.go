package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	x := sum(2, 2)
	if x != 4 {
		t.Error("expecting: ", 4, " Got: ", x)
	}
}
