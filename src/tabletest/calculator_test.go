package tabletest

import "testing"

func TestCalculator(t *testing.T) {

	type testCase struct {
		minimumPurchaseAmount int
		discountAmount        int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{minimumPurchaseAmount: 100, discountAmount: 20, purchaseAmount: 150, expectedAmount: 130},
		{minimumPurchaseAmount: 100, discountAmount: 20, purchaseAmount: 200, expectedAmount: 160},
		{minimumPurchaseAmount: 100, discountAmount: 20, purchaseAmount: 350, expectedAmount: 290},
		{minimumPurchaseAmount: 100, discountAmount: 20, purchaseAmount: 30, expectedAmount: 30},
	}

	for _, tc := range testCases {
		calculator := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discountAmount)

		amount := calculator.Calculator(tc.purchaseAmount)

		if amount != tc.expectedAmount {
			t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
		}
	}
}
