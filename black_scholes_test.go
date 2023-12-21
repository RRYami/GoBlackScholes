package black_scholes_test

import (
	"black_scholes"
	"math"
	"testing"
)

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a - b) < tolerance
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
	gotc, errC := option.CallPrice()
	if errC != nil {
		t.Fatalf("want no error for invalid input parameter, got %v", errC)
	}
	if !closeEnough(gotc, wantc, 0.00001) {
		t.Errorf("want %v, got %v", wantc, gotc)
	}
	gotp, errP := option.PutPrice()
	if errP != nil {
		t.Fatalf("want no error for invalid input parameter, got %v", errP)
	}
	if !closeEnough(gotp, wantp, 0.00001) {
		t.Errorf("want %v, got %v", wantc, gotc)
	}
}

func TestOptionDelta(t *testing.T) {
	t.Parallel()
	option := &black_scholes.Option{
		S: 100.0,
		K: 100.0,
		T: 1.0,
		R: 0.05,
		V: 0.2,
	}
	wantCD := 0.6368306511756191
	wantPD := -0.3631693488243809

	gotCD, errCD := option.CallDelta()
	if errCD != nil {
		t.Fatalf("want no error for invalid input parameter, got %v", errCD)
	}
	if !closeEnough(gotCD, wantCD, 0.00001) {
		t.Errorf("want %v, got %v", wantCD, gotCD)
	}
	gotPD, errPD := option.PutDelta()
	if errPD != nil {
		t.Fatalf("want no error for invalid input parameter, got %v", errPD)
	}
	if !closeEnough(gotPD, wantPD, 0.00001) {
		t.Errorf("want %v, got %v", wantPD, gotPD)
	}
}

func TestOptionGamma(t *testing.T){
	t.Parallel()
	option := &black_scholes.Option{
		S: 100.0,
		K: 100.0,
		T: 1.0,
		R: 0.05,
		V: 0.2,
	}
	wantG := 0.018762017345846895

	gotG, errG := option.Gamma()
	if errG != nil {
		t.Fatalf("want no error for invalid input parameter, got %v", errG)
	}
	if !closeEnough(gotG, wantG, 0.00001) {
		t.Errorf("want %v, got %v", wantG, gotG)
	}
}

func TestOptionVega(t *testing.T){
	t.Parallel()
	option := &black_scholes.Option{
		S: 100.0,
		K: 100.0,
		T: 1.0,
		R: 0.05,
		V: 0.2,
		}
		wantVeg := 0.3752403469169379
		gotVeg, errVeg := option.Vega()
		if errVeg != nil {
			t.Fatalf("want no error for invalid input parameter, got %v", errVeg)
		}
		if !closeEnough(gotVeg, wantVeg, 0.00001) {
			t.Errorf("want %v, got %v", wantVeg, gotVeg)
		}
}

func TestOptionCallTheta(t *testing.T){
	t.Parallel()
	option := &black_scholes.Option{
		S: 100.0,
		K: 100.0,
		T: 1.0,
		R: 0.05,
		V: 0.2,
		}
		wantCT := -0.01757267820941972
		gotCT, errCT := option.CallTheta()
		if errCT != nil {
			t.Fatalf("want no error for invalid input parameter, got %v", errCT)
		}
		if !closeEnough(gotCT, wantCT, 0.00001) {
			t.Errorf("want %v, got %v", wantCT, gotCT)
		}
}

func TestOptionPutTheta(t *testing.T){
	t.Parallel()
	option := &black_scholes.Option{
		S: 100.0,
		K: 100.0,
		T: 1.0,
		R: 0.05,
		V: 0.2,
		}
		wantPT := -0.004542138147766099
		gotPT, errPT := option.PutTheta()
		if errPT != nil {
			t.Fatalf("want no error for invalid input parameter, got %v", errPT)
		}
		if !closeEnough(gotPT, wantPT, 0.00001) {
			t.Errorf("want %v, got %v", wantPT, gotPT)
		}
}

func TestOption_CallRho(t *testing.T){
	t.Parallel()
	option := &black_scholes.Option{
		S: 100.0,
		K: 100.0,
		T: 1.0,
		R: 0.05,
		V: 0.2,
		}
		wantCR := 0.5323248154537634
		gotCR, errCR := option.CallRho()
		if errCR != nil {
			t.Fatalf("want no error for invalid input parameter, got %v", errCR)
		}
		if !closeEnough(gotCR, wantCR, 0.00001) {
			t.Errorf("want %v, got %v", wantCR, gotCR)
		}
}

func TestOption_PutRho(t *testing.T){
	t.Parallel()
	option := &black_scholes.Option{
		S: 100.0,
		K: 100.0,
		T: 1.0,
		R: 0.05,
		V: 0.2,
		}
		wantPR := -0.4189046090469506
		gotPR, errPR := option.PutRho()
		if errPR != nil {
			t.Fatalf("want no error for invalid input parameter, got %v", errPR)
		}
		if !closeEnough(gotPR, wantPR, 0.00001) {
			t.Errorf("want %v, got %v", wantPR, gotPR)
		}
}