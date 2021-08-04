package main

import (
	"testing"
)

func TestAdditional(t *testing.T) {
	x := 1
	y := 2
	result := addition(x,y)

	if result == x + y {
		t.Error("足し算間違えてるよ〜")
	}
}
