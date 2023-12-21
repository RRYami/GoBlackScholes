package main

import (
	"black_scholes"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	option := []*black_scholes.Option{
		{
			S: 100.0, // Underlying price
			K: 100.0, // Strike price
			T: 1.0,   // Time to maturity
			R: 0.05,  // Risk-free interest rate
			V: 0.2,   // Volatility}
		},
		{
			S: 110.0, // Underlying price
			K: 100.0, // Strike price
			T: 1.0,   // Time to maturity
			R: 0.05,  // Risk-free interest rate
			V: 0.2,   // Volatility}
		},
		{
			S: 120.0, // Underlying price
			K: 100.0, // Strike price
			T: 1.0,   // Time to maturity
			R: 0.05,  // Risk-free interest rate
			V: 0.2,   // Volatility}
		},
	}
	for i, option := range option {
		callPrice, err_call := option.CallPrice()
		if err_call != nil {
			println("Option number:", i, err_call.Error())
		} else {
			println("Option number:", i, "Call price:", callPrice)
		}
		putPrice, err_put := option.PutPrice()
		if err_put != nil {
			println("Option number:", i, err_put.Error())
		} else {
			println("Option number:", i, "Put price:", putPrice)
		}
	}
}
