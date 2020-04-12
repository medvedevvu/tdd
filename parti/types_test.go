package parti

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5.0)
	five.times(2)
	if five.amount != 10.0 {
		t.Fatalf(" not equal 10 ")
	}
}
