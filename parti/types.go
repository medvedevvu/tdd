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

func (d *Dollar) equals(d1 *Dollar) bool {
	return (*d).amount == (*d1).amount
}
