package main

import (
	"testing"
)

func TestPlus(t *testing.T) {
	x := 1
	y := 2
	result := plus(x,y)

	if result != x + y {
		t.Error("足し算間違えてるよ〜")
	}
}
