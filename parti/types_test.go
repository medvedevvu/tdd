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

// We have an equals operation now and can change previous tests
func TestWithNewTimesMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.times(2)
	if product.equals(five) {
		t.Fatalf("%v equal %v ", product.amount, five.amount)
	}
	if product.equals(five.times(3)) {
		t.Fatalf(" %v equal %v ", product.amount, five.amount)
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
