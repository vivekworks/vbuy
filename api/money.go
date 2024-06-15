package vbuy

import "math"

type Money float64

func (m Money) Rounded() Money  {
    return Money(math.Round(float64(m)*100) / 100)
}