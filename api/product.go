package vbuy

import (
    "fmt"
    "slices"
    "time"
    "unicode"
)

var (
    ProductCutOffDate = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
)

type Category string

var (
    Books       Category = "Books"
    Smartphones Category = "Smartphones"
    Televisions Category = "Televisions"
)

var Categories = []Category{Books, Smartphones, Televisions}

type Product struct {
    ID           string    `json:"id"`
    Name         string    `json:"name"`
    ReleasedDate *Date     `json:"releasedDate"`
    Model        string    `json:"model"`
    Price        *Money    `json:"price"`
    Manufacturer string    `json:"manufacturer"`
    Category     Category  `json:"category"`
    IsActive     bool      `json:"isActive"`
    CreatedAt    time.Time `json:"createdAt"`
    UpdatedAt    time.Time `json:"updatedAt"`
    CreatedBy    string    `json:"createdBy"`
    UpdatedBy    string    `json:"updatedBy"`
}

type ProductCreate struct {
    Name         string   `json:"name"`
    ReleasedDate *Date    `json:"releasedDate"`
    Model        string   `json:"model"`
    Price        Money    `json:"price"`
    Manufacturer string   `json:"manufacturer"`
    IsActive     bool     `json:"isActive"`
    Category     Category `json:"category"`
}

func (pc *ProductCreate) Validate() []*ErrorDetail {
    var errorDetail []*ErrorDetail
    if err := validateName(pc.Name); err != nil {
        errorDetail = append(errorDetail, err)
    }
    if err := validateReleasedDate(pc.ReleasedDate); err != nil {
        errorDetail = append(errorDetail, err)
    }
    if err := validateCategory(pc.Category); err != nil {
        errorDetail = append(errorDetail, err)
    }
    return errorDetail
}

type ProductUpdate struct {
    Price    *Money `json:"price"`
    IsActive bool   `json:"isActive"`
}

func validateCategory(category Category) *ErrorDetail {
    categoryError := &ErrorDetail{
        Field: "category",
    }
    if !slices.Contains(Categories, category) {
        categoryError.AppendMessage(fmt.Sprintf("must be one of %v", Categories))
    }
    if categoryError.HasMessages() {
        return categoryError
    }
    return nil
}

func validateName(name string) *ErrorDetail {
    nameError := &ErrorDetail{
        Field: "name",
    }
    if len(name) == 0 {
        nameError.AppendMessage("must not be empty")
    } else if len(name) > 0 {
        if !unicode.IsLetter(rune(name[0])) {
            nameError.AppendMessage("first letter must be an alphabet")
        }
        if len(name) > 200 {
            nameError.AppendMessage("must not exceed 200 characters")
        }
    }
    if nameError.HasMessages() {
        return nameError
    }
    return nil
}

func validateReleasedDate(d *Date) *ErrorDetail {
    rDateError := &ErrorDetail{
        Field: "releasedDate",
    }
    if d == nil {
        rDateError.AppendMessage("must not be empty")
    } else if time.Time(*d).Before(ProductCutOffDate) {
        rDateError.AppendMessage("must be on or after product cutoff date 2000-01-01")
    }
    if rDateError.HasMessages() {
        return rDateError
    }
    return nil
}
