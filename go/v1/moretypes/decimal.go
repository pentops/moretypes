package moretypes

import "math/big"

func (d *Decimal) AsBig() *big.Rat {
	val, ok := new(big.Rat).SetString(d.String_)
	if !ok {
		panic("Invalid Decimal value")
	}
	return val
}

func DecimalFromBig(val *big.Rat) (*Decimal, error) {
	denom := val.Denom()

	return &Decimal{String_: val.FloatString(0)}
}
