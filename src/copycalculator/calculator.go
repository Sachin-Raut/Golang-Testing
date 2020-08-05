package copycalculator

//DiscountCalculator is
type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountAmount        int
}

//NewDiscountCalculator is
func NewDiscountCalculator(minimumPurchaseamount int, discountAmount int) *DiscountCalculator {
	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseamount,
		discountAmount:        discountAmount,
	}
}

//Calclulate is
func (c *DiscountCalculator) Calclulate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchaseAmount {
		return purchaseAmount - c.discountAmount
	}
	return purchaseAmount
}
