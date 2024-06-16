package vbuy

import (
    "encoding/json"
    "time"
)

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) error {
    var dateStr string
    err := json.Unmarshal(data, &dateStr)
    if err != nil {
        return ErrInternalServer
    }
    t, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        return ErrInternalServer
    }
    *d = Date(t)
    return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
    return json.Marshal(time.Time(*d).Format("2006-01-02"))
}
