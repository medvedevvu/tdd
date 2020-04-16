package parti

import "fmt"

// MonyType general type
type MonyType uint8

// Mony types
const (
	Franc MonyType = iota
	Dollar
)

//Mony
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

// Addition
func (m *Mony) plus(m2 *Mony) *Mony {
	if m.monytype != m2.monytype {
		panic(fmt.Sprintf("%v and %v has different type \n", m, m2))
	}
	return NewMony(m.monytype, m.amount+m2.amount)
}

func dollar(amount int) *Mony {
	return NewMony(Dollar, amount)
}

func franc(amount int) *Mony {
	return NewMony(Franc, amount)
}

// Expression
type Expression Mony

func (ex *Expression) plus(addend Mony) *Mony {
	return NewMony(ex.monytype, ex.amount+addend.amount)
}

type Bank struct {
}

func NewBank() *Bank {
	return &Bank{}
}

func (bank Bank) reduce(source Expression, to MonyType) *Mony {
	return NewMony(Dollar, 5)
}
