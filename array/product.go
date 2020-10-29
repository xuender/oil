package array

import "errors"

var (
	errNum    = errors.New("num should be greater then 0")
	errRepeat = errors.New("len should be greater then 0")
)

// Product is carthesian product
// https://en.wikipedia.org/wiki/Cartesian_product
type Product struct {
	currPositions []int
	num           int
	repeat        int
	init          bool
}

// NewProduct is constructor
func NewProduct(num, repeat int) (*Product, error) {
	if repeat < 1 {
		return nil, errRepeat
	}

	if num < 1 {
		return nil, errNum
	}

	currPosition := make([]int, repeat)

	for i := 0; i < repeat; i++ {
		currPosition[i] = 0
	}

	return &Product{currPosition, num, repeat, false}, nil
}

// Next generates the next value for product
func (p *Product) Next() bool {
	if !p.init {
		p.init = true
		return true
	}

	maxIndex := p.num - 1
	numberMaxIndexes := 0

	for i := p.repeat - 1; i >= 0; i-- {
		if p.currPositions[i] == maxIndex {
			numberMaxIndexes++
		}
	}

	if numberMaxIndexes == p.repeat {
		return false
	}

	for i := p.repeat - 1; i >= 0; i-- {
		if p.currPositions[i] < maxIndex {
			p.currPositions[i]++

			break
		}

		p.currPositions[i] = 0
	}

	return true
}

// Value gets the current value
func (p *Product) Value() []int {
	return p.currPositions
}
