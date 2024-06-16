package vbuy

import (
    "encoding/json"
    "math"
)

type Money float64

func (m *Money) UnmarshalJSON(data []byte) error {
    var moneyFl float64
    err := json.Unmarshal(data, &moneyFl)
    if err != nil {
        return ErrInternalServer
    }
    *m = Money(math.Round(moneyFl*100) / 100)
    return nil
}

func (m *Money) MarshalJSON() ([]byte, error) {
    return json.Marshal(float64(*m))
}
