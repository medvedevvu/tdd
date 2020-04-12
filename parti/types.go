package parti

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
