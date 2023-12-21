package black_scholes_test

import (
	"black_scholes"
	"testing"
)

func closeEnough(a, b, tolerance float64) bool {
	return a-b < tolerance
}

func TestOptionPricing(t *testing.T) {
	t.Parallel()
	option := &black_scholes.Option{
		S: 100.0, // Underlying price
		K: 100.0, // Strike price
		T: 1.0,   // Time to maturity
		R: 0.05,  // Risk-free interest rate
		V: 0.2,   // Volatility
	}
	wantc := 10.45058
	wantp := 5.573526
	gotc, err_c := option.CallPrice()
	if err_c != nil {
		t.Fatalf("want no error for invalid input parameter, got %v", err_c)
	}
	if !closeEnough(gotc, wantc, 0.00001) {
		t.Errorf("want %v, got %v", wantc, gotc)
	}
	gotp, err_p := option.PutPrice()
	if err_p != nil {
		t.Fatalf("want no error for invalid input parameter, got %v", err_p)
	}
	if !closeEnough(gotp, wantp, 0.00001) {
		t.Errorf("want %v, got %v", wantc, gotc)
	}
}
