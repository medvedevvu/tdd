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

func (ex *Mony) plus2(addend *Mony) *Sum {
	return NewSum(ex, addend)
}

type Bank struct {
	rate int
}

func NewBank() *Bank {
	return &Bank{rate: 1}
}

func (bank Bank) reduce(source *Sum, to MonyType) *Mony {
	return NewMony(to, source.addend.amount+source.augend.amount)
}

func (bank *Bank) AddRate(frm MonyType, to MonyType, rate int) {
	bank.rate = rate
}

//Sum
type Sum struct {
	augend *Mony
	addend *Mony
}

//Sum constructor
func NewSum(augend *Mony, addend *Mony) *Sum {
	return &Sum{augend: augend, addend: addend}
}
