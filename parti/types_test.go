package parti

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.times(2)
	if five.amount != 10 {
		t.Fatalf(" not equal 10 ")
	}
}

func TestWithNewTimesMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	if product.amount != 10 {
		t.Fatalf("%v not equal 10 ", product.amount)
	}
	product = five.times(3)
	if product.amount != 15 {
		t.Fatalf(" %v not equal 15 ", product.amount)
	}
}

func TestEquality(t *testing.T) {
	d5 := NewDollar(5)
	dOther5 := NewDollar(5)
	if !d5.equals(dOther5) {
		t.Fatalf("%v not equal %v", d5, dOther5)
	}
	dOther6 := NewDollar(6)
	if d5.equals(dOther6) {
		t.Fatalf("%v is equal %v", d5, dOther5)
	}

}
