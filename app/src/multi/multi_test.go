package multiple

import (
	"testing"
)

// 掛け算のテスト
func TestMultiple(t *testing.T) {
	x := 2
	y := 1
	result := multiple(x, y)

	if result != x * y {
		t.Error("掛け算間違えてるよ〜")
	}
}
