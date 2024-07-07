package vbuy

import (
    "time"
    "unicode"
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
    if len(pc.Name) <= 0 || len(pc.Name) > 200 {
        nameError.AppendMessage("name must have between 1 and 200 characters")
    }
    if len(pc.Name) > 0 {
        if !unicode.IsLetter(rune(pc.Name[0])) {
            nameError.AppendMessage("first letter must be an alphabet")
        }
    }
    if nameError.HasMessages() {
        errorDetail = append(errorDetail, nameError)
    }
    return errorDetail
}

type ProductUpdate struct {
    Price    *Money `json:"price"`
    IsActive bool   `json:"isActive"`
}
