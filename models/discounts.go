package models

import (
	"database/sql"
	"fmt"

	"example.com/shoppingCart-api/db"
)

type Discount struct {
	ID       int     `json:"id" binding:"required"`
	Code     string  `json:"code" binding:"required"`
	Discount float64 `json:"discount" binding:"required"`
	Type     string  `json:"type" binding:"required"`
}

func GetDiscountValue(code string, totalPrice float64) (float64, error) {

	var discount Discount

	if err := db.DB.QueryRow("SELECT * FROM discounts where code = ?",
		code).Scan(&discount.ID, &discount.Code, &discount.Discount, &discount.Type); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("Discount code not found")
		}
		return 0, fmt.Errorf("could not fetch discount: %w", err)
	}

	if discount.Type == "percentage" {
		return (discount.Discount / 100) * totalPrice, nil
	}
	return discount.Discount, nil
}
