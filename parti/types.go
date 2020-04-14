package parti

// MonyType general type
type MonyType uint8

// Mony types
const (
	Franc MonyType = iota
	Dollar
)

type Mony struct {
	monytype MonyType
	amount   int
}

func NewMony(monytype MonyType, amount int) *Mony {
	return &Mony{monytype: monytype, amount: amount}
}

// Multyplication
func (f *Mony) times(tm int) *Mony {
	return NewMony(f.monytype, f.amount*tm)
}

//equals
func (f *Mony) equals(f1 *Mony) bool {
	return (*f).amount == (*f1).amount && f.monytype == f1.monytype
}
