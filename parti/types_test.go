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

func TestWithNewTimesMultiplication(t *testing.T) {
	five := NewDollar(5.0)
	product := five.times(2)
	if product.amount != 10.0 {
		t.Fatalf("%v not equal 10 ", product.amount)
	}
	product = five.times(3)
	if product.amount != 15.0 {
		t.Fatalf(" %v not equal 15 ", product.amount)
	}
}
