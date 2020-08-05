package copycalculator

import "testing"

func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calclulate(150)
	if amount != 130 {
		t.Error("received amt-", amount)
		t.Fail()
	}
}

func TestDiscountNotApplied(t *testing.T) {
	calculator := NewDiscountCalculator(100, 20)
	amount := calculator.Calclulate(50)
	if amount != 50 {
		t.Fail()
	}
}
