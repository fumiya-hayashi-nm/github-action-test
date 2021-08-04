package minus

import (
	"testing"
)

// 引き算のテスト
func TestMinus(t *testing.T) {
	x := 2
	y := 1
	result := minus(x, y)

	if result != x - y {
		t.Error("引き算間違えてるよ〜")
	}
}
