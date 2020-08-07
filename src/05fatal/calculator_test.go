package fatal

import "testing"

func TestNewDiscountCalculatorShouldFailWithZero(t *testing.T) {
	_, err := NewDiscountCalculator(0, 20)
	if err == nil {
		t.Fatalf("couldn't instantiate the calculator-%s", err.Error())
	}
}

func TestNewDiscountCalculator(t *testing.T) {
	dc, err := NewDiscountCalculator(100, 20)
	if err != nil {
		t.Fatalf("couldn't instantiate the calculator-%s", err.Error())
	}
	if dc == nil {
		t.Fatalf("was not expecting nil response")
	}
	if dc.minimumPurchaseAmount != 100 {
		t.Fatalf("was expecting minimum purchase amount = 100")
	}
	if dc.discountAmount != 20 {
		t.Fatalf("was expecting discount amount = 20")
	}
}

func TestCalculatorDiscountApplied(t *testing.T) {
	calculator, _ := NewDiscountCalculator(100, 20)
	amount := calculator.Calculator(200)

	if amount != 160 {
		t.Fatalf("was expecting amount = 160")
	}
}

func TestCalculatorDiscountNotApplied(t *testing.T) {
	calculator, _ := NewDiscountCalculator(100, 20)
	amount := calculator.Calculator(50)

	if amount != 50 {
		t.Fatalf("was expecting amount = 50")
	}
}
