package vbuy

import (
	"time"
)

type Product struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	ReleasedDate Date      `json:"releasedDate"`
	Model        string    `json:"model"`
	Price        Money     `json:"price"`
	Manufacturer string    `json:"manufacturer"`
	IsActive     bool      `json:"isActive"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
}

type ProductCreate struct {
	Name         string `json:"name"`
	ReleasedDate Date   `json:"releasedDate"`
	Model        string `json:"model"`
	Price        Money  `json:"price"`
	Manufacturer string `json:"manufacturer"`
	IsActive     bool   `json:"isActive"`
}

type ProductUpdate struct {
	Price    Money `json:"price"`
	IsActive bool  `json:"isActive"`
}
