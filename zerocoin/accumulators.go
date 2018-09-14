package zerocoin

import (
	"errors"
	"math/big"
)

// Accumulator represents an RSA-based accumulator.
type Accumulator struct {
	params       *AccumulatorAndProofParams
	value        *big.Int
	denomination Denomination
}

// NewAccumulator initializes a new empty accumulator with given parameters
// and denomination.
func NewAccumulator(params *AccumulatorAndProofParams, d Denomination) (*Accumulator, error) {
	if !params.Initialized {
		return nil, errors.New("accumulator and proof params must be initialized")
	}
	return &Accumulator{
		denomination: d,
		value:        params.AccumulatorBase,
	}, nil
}

// NewAccumulatorWithValue initializes an accumulator with given zerocoin params,
// denomination, and with a preset value.
func NewAccumulatorWithValue(params *AccumulatorAndProofParams, d Denomination, value *big.Int) (*Accumulator, error) {
	a := &Accumulator{}

	a.params = params
	a.denomination = d

	if !a.params.Initialized {
		return nil, errors.New("zerocoin parameters must be initialized")
	}

	if value.Cmp(big.NewInt(0)) != 0 {
		a.value = value
	} else {
		a.value = params.AccumulatorBase
	}
	return a, nil
}

// Increment adds a value to the accumulator
func (a Accumulator) Increment(value *big.Int) {
	a.value = new(big.Int).Exp(a.value, value, a.params.AccumulatorModulus)
}

// Accumulate a given coin if it is valid and the denomination matches.
func (a Accumulator) Accumulate(coin *PublicCoin) error {
	if a.value.Cmp(big.NewInt(0)) == 0 {
		return errors.New("accumulator is not initialized")
	}

	if a.denomination != coin.denomination {
		return errors.New("accumulator does not match the coin being accumulated")
	}

	if coin.Validate() {
		a.Increment(coin.value)
	} else {
		return errors.New("coin is not valid")
	}
	return nil
}

// AccumulatorWitness is a witness that a public coin is
// in the accumulation of a set of coins.
type AccumulatorWitness struct {
	witness Accumulator
	element PublicCoin
}

// ResetValue resets the value of the accumulator witness to a
// given checkpoint and public coin.
func (a AccumulatorWitness) ResetValue(checkpoint *Accumulator, coin PublicCoin) {
	a.witness.value = checkpoint.value
	a.element = coin
}

// AddElement adds a public coin to the accumulator.
func (a AccumulatorWitness) AddElement(coin PublicCoin) {
	if a.element.value != coin.value {
		a.witness.Accumulate(&coin)
	}
}

// VerifyWitness verifies that a witness matches the accumulator.
func (a AccumulatorWitness) VerifyWitness(acc *Accumulator, p *PublicCoin) (bool, error) {
	temp, err := NewAccumulatorWithValue(a.witness.params, a.witness.denomination, a.witness.value)
	if err != nil {
		return false, err
	}

	temp.Accumulate(&a.element)

	return temp.value == acc.value && a.element.Equal(*p), nil
}

// Accumulators each denomination of accumulator.
type Accumulators map[Denomination]Accumulator
