package parti

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewMony(Dollar, 5)
	if five.equals(five.times(2)) {
		t.Fatalf(" %v equal %v ", five.amount, five.times(2).amount)
	}
}

// We have an equals operation now and can change previous tests
func TestWithNewTimesMultiplication(t *testing.T) {
	five := NewMony(Dollar, 5)
	product := five.times(2)
	if product.equals(five) {
		t.Fatalf("%v equal %v ", product.amount, five.amount)
	}
	if product.equals(five.times(3)) {
		t.Fatalf(" %v equal %v ", product.amount, five.amount)
	}
}
func TestEquality(t *testing.T) {
	d5 := NewMony(Dollar, 5)
	dOther5 := NewMony(Dollar, 5)
	if !d5.equals(dOther5) {
		t.Fatalf("%v not equal %v", d5, dOther5)
	}
	dOther6 := NewMony(Dollar, 6)
	if d5.equals(dOther6) {
		t.Fatalf("%v is equal %v", d5, dOther5)
	}
}

func TestFrancMultiplication(t *testing.T) {
	five := NewMony(Franc, 5)
	if !five.times(2).equals(NewMony(Franc, 10)) {
		t.Fatalf(" %v not equal %v ", five.times(2), NewMony(Franc, 10))
	}
	if !five.times(3).equals(NewMony(Franc, 15)) {
		t.Fatalf(" %v not equal %v ", five.times(3), NewMony(Franc, 15))
	}
}

func TestMonyEquals(t *testing.T) {
	fivefranc := NewMony(Franc, 5)
	fivedollar := NewMony(Dollar, 5)

	if fivefranc.equals(fivedollar) {
		t.Fail()
	}

	if !fivefranc.equals(fivefranc) {
		t.Fail()
	}

	if !fivefranc.times(3).equals(NewMony(Franc, 15)) {
		t.Fail()
	}

}

func TestSimpleAddition(t *testing.T) {
	fiveDollar := NewMony(Dollar, 5)
	moreFiveDollar := NewMony(Dollar, 5)
	sum := fiveDollar.plus(moreFiveDollar)
	if !sum.equals(NewMony(Dollar, 10)) {
		t.Fatalf("%v != %v \n", sum, NewMony(Dollar, 10))
	}
}

func TestOtherSimpleAddition(t *testing.T) {
	sum := dollar(5).plus(dollar(5))
	if !sum.equals(NewMony(Dollar, 10)) {
		t.Fatalf("%v  not equal %v \n", sum, NewMony(Dollar, 10))
	}
}

func TestWithBankSimpleAddition(t *testing.T) {
	var five *Mony = dollar(5)
	var sum Expression = Expression(*five.plus(five))
	var bank *Bank = NewBank()                   // создадим объект банк
	var reduced *Mony = bank.reduce(sum, Dollar) // слабое связывание
	if !dollar(10).equals(reduced) {
		t.Fatalf("%v  not equal %v \n", sum, NewMony(Dollar, 10))
	}
}

func TestPlusReturnsSum(t *testing.T) {
	five := dollar(5)
	var result Expression = Expression(*five.plus(five))
	var sum Sum = Sum(result)
	if !five.equals(sum.augend) {
		t.Fatalf("%v  not equal %v \n", five, sum.augend)
	}
	if !five.equals(sum.addend) {
		t.Fatalf("%v  not equal %v \n", five, sum.addend)
	}

}
