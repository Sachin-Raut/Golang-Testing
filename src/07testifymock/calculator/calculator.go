package calculator

import (
	"errors"

	"github.com/Sachin-Raut/Golang-Testing/src/databasemock/database"
)

//DiscountCalculator is
type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountRepository    database.Discount
}

//NewDiscountCalculator is
func NewDiscountCalculator(minimumPurchaseAmount int, discountRepository database.Discount) (*DiscountCalculator, error) {
	if minimumPurchaseAmount <= 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount can't be <= 0")
	}
	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountRepository:    discountRepository,
	}, nil
}

//Calculate is
func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchaseAmount {
		multiplier := purchaseAmount / c.minimumPurchaseAmount
		discount := c.discountRepository.FindCurrentDiscount()
		return purchaseAmount - (discount * multiplier)
	}
	return purchaseAmount
}
