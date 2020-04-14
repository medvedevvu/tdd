package parti

//Franc
type Franc struct {
	amount int
}

//Dollar constructor
func NewFranc(amount int) *Franc {
	return &Franc{amount: amount}
}

// Multyplication
func (f *Franc) times(tm int) *Franc {
	return NewFranc(f.amount * tm)
}

//equals
func (f *Franc) equals(f1 *Franc) bool {
	return (*f).amount == (*f1).amount
}

//Dollar
type Dollar struct {
	amount int
}

//Dollar constructor
func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}

// Multyplication
func (d *Dollar) times(tm int) *Dollar {
	return NewDollar(d.amount * tm)
}

//equals
func (d *Dollar) equals(d1 *Dollar) bool {
	return (*d).amount == (*d1).amount
}
