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
		callPrice, errCall := option.CallPrice()
		if errCall != nil {
			println("Option number:", i, errCall.Error())
		} else {
			println("Option number:", i, "Call price:", callPrice)
		}
		putPrice, errPut := option.PutPrice()
		if errPut != nil {
			println("Option number:", i, errPut.Error())
		} else {
			println("Option number:", i, "Put price:", putPrice)
		}
	}
}
