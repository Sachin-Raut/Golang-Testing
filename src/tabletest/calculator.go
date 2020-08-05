package tabletest

//DiscountCalculator is
type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountAmount        int
}

//NewDiscountCalculator is
func NewDiscountCalculator(minimumPurchaseAmount int, discountAmount int) *DiscountCalculator {
	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountAmount:        discountAmount,
	}
}

//Calculator is
func (c DiscountCalculator) Calculator(purchaseAmount int) int {

	if purchaseAmount > c.minimumPurchaseAmount {
		multiplier := purchaseAmount - c.minimumPurchaseAmount
		return purchaseAmount - (c.discountAmount * multiplier)
	}
	return purchaseAmount
}
