package vbuy

import (
    "time"
    "unicode"
)

var (
    ProductCutOffDate = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
)

type Product struct {
    ID           string    `json:"id"`
    Name         string    `json:"name"`
    ReleasedDate *Date     `json:"releasedDate"`
    Model        string    `json:"model"`
    Price        *Money    `json:"price"`
    Manufacturer string    `json:"manufacturer"`
    IsActive     bool      `json:"isActive"`
    CreatedAt    time.Time `json:"createdAt"`
    UpdatedAt    time.Time `json:"updatedAt"`
    CreatedBy    string    `json:"createdBy"`
    UpdatedBy    string    `json:"updatedBy"`
}

type ProductCreate struct {
    Name         string `json:"name"`
    ReleasedDate *Date  `json:"releasedDate"`
    Model        string `json:"model"`
    Price        Money  `json:"price"`
    Manufacturer string `json:"manufacturer"`
    IsActive     bool   `json:"isActive"`
}

func (pc *ProductCreate) Validate() []ErrorDetail {
    var errorDetail []ErrorDetail
    nameError := ErrorDetail{
        Field: "name",
    }
    if len(pc.Name) == 0 {
        nameError.AppendMessage("must not be empty")
    } else if len(pc.Name) > 0 {
        if !unicode.IsLetter(rune(pc.Name[0])) {
            nameError.AppendMessage("first letter must be an alphabet")
        }
        if len(pc.Name) > 200 {
            nameError.AppendMessage("must not exceed 200 characters")
        }
    }
    if nameError.HasMessages() {
        errorDetail = append(errorDetail, nameError)
    }
    rDateError := ErrorDetail{
        Field: "releasedDate",
    }
    if pc.ReleasedDate == nil {
        rDateError.AppendMessage("must not be empty")
    } else if time.Time(*pc.ReleasedDate).Before(ProductCutOffDate) {
        rDateError.AppendMessage("must be on or after product cutoff date 2000-01-01")
    }
    if rDateError.HasMessages() {
        errorDetail = append(errorDetail, rDateError)
    }
    return errorDetail
}

type ProductUpdate struct {
    Price    *Money `json:"price"`
    IsActive bool   `json:"isActive"`
}
