package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DiscountRepositoryMock struct {
	mock.Mock
}

//FindCurrentDiscount is
func (drm DiscountRepositoryMock) FindCurrentDiscount() int {
	args := drm.Called()
	return args.Int(0)
}

func TestNewDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		purchaseAmount        int
		discount              int
		expectedAmount        int
	}

	testCases := []testCase{
		{
			name:                  "apply 20",
			minimumPurchaseAmount: 100,
			purchaseAmount:        150,
			discount:              20,
			expectedAmount:        130,
		},
		{
			name:                  "zero discount",
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              0,
			expectedAmount:        50,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			discountRepositoryMock := DiscountRepositoryMock{}
			discountRepositoryMock.On("FindCurrentDiscount").Return(tc.discount)
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepositoryMock)

			if err != nil {
				t.Fatalf("could not instantiate the calculator - %s", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmount)

			assert.Equal(t, tc.expectedAmount, amount)
		})
	}
}
