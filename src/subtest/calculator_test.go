package subtest

import "testing"

func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		expectedAmount        int
	}

	testCases := []testCase{
		{name: "expected amount 130", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 150, expectedAmount: 130},
		{name: "expected amount 160", minimumPurchaseAmount: 100, discount: 2, purchaseAmount: 200, expectedAmount: 160},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calculator := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)

			amount := calculator.Calculator(tc.purchaseAmount)

			if amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
			}
		})
	}
}
