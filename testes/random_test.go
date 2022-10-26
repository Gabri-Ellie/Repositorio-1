package main

import "testing"

func TestSoma(t *testing.T) {
	x := Soma(3, 2, 1)
	res := 6
	if x != res {
		t.Error("Expected:", res, "Got:", x)
	}
}
