package fatal

import "errors"

//DiscountCalculator is
type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountAmount        int
}

//NewDiscountCalculator is
func NewDiscountCalculator(minimumPurchaseAmount int, discountAmount int) (*DiscountCalculator, error) {

	if minimumPurchaseAmount <= 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount should be > 0")
	}
	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountAmount:        discountAmount,
	}, nil
}

//Calculator is
func (c DiscountCalculator) Calculator(purchaseAmount int) int {

	if purchaseAmount > c.minimumPurchaseAmount {
		multiplier := purchaseAmount / c.minimumPurchaseAmount
		return purchaseAmount - (c.discountAmount * multiplier)
	}
	return purchaseAmount
}
