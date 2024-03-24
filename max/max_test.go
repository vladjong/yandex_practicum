package max_test

import (
	"testing"

	"github.com/vladjong/yandex_practicum/max"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	res := max.MaxInt(a, b)

	if res != b {
		t.Errorf("expected %d, got %d", b, res)
	}
}
