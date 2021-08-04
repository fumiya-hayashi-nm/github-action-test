package sub

import (
	"testing"
)

// 引き算のテスト
func TestSubSubtraction(t *testing.T) {
	x := 2
	y := 1
	result := subtraction( x, y)

	if result != x - y {
		t.Error("引き算間違えてるよ〜")
	}
}
