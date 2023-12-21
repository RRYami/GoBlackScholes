package black_scholes

import (
	"errors"
	"math"
)

type Option struct {
	S float64 // Underlying price
	K float64 // Strike price
	T float64 // Time to maturity
	R float64 // Risk-free interest rate
	V float64 // Volatility
}

func standardNormalCDF(x float64) float64 {
	return 0.5 * (1 + math.Erf(x/math.Sqrt(2)))
}

func gaussianPDF(x, mean, stddev float64) float64 {
	const invSqrt2Pi = 0.3989422804014327 // 1 / sqrt(2 * pi)
	a := (x - mean) / stddev
	return invSqrt2Pi / stddev * math.Exp(-0.5*a*a)
}

func (o *Option) CallPrice() (float64, error) {
	if o.K == 0 {
		return 0, errors.New("strike price must be greater than 0")
	}
	d1 := (math.Log(o.S/o.K) + (o.R+0.5*math.Pow(o.V, 2))*o.T) / (o.V * math.Sqrt(o.T))
	d2 := d1 - o.V*math.Sqrt(o.T)
	return o.S*standardNormalCDF(d1) - o.K*math.Exp(-o.R*o.T)*standardNormalCDF(d2), nil
}

func (o *Option) PutPrice() (float64, error) {
	if o.K == 0 {
		return 0, errors.New("strike price must be greater than 0")
	}
	d1 := (math.Log(o.S/o.K) + (o.R+0.5*math.Pow(o.V, 2))*o.T) / (o.V * math.Sqrt(o.T))
	d2 := d1 - o.V*math.Sqrt(o.T)
	return o.K*math.Exp(-o.R*o.T)*standardNormalCDF(-d2) - o.S*standardNormalCDF(-d1), nil
}

func (o *Option) CallDelta() (float64, error) {
	if o.K == 0 {
		return 0, errors.New("strike price must be greater than 0")
	}
	d1 := (math.Log(o.S/o.K) + (o.R+0.5*math.Pow(o.V, 2))*o.T) / (o.V * math.Sqrt(o.T))
	return standardNormalCDF(d1), nil
}

func (o *Option) PutDelta() (float64, error) {
	if o.K == 0 {
		return 0, errors.New("strike price must be greater than 0")
	}
	d1 := (math.Log(o.S/o.K) + (o.R+0.5*math.Pow(o.V, 2))*o.T) / (o.V * math.Sqrt(o.T))
	return standardNormalCDF(d1) - 1, nil
}

func (o *Option) Gamma() (float64, error) {
	if o.K == 0 || o.S == 0 || o.T == 0 {
		return 0, errors.New("strike price, spot price and time to maturity must be greater than 0")
	}
	d1 := (math.Log(o.S/o.K) + (o.R+0.5*math.Pow(o.V, 2))*o.T) / (o.V * math.Sqrt(o.T))
	return gaussianPDF(d1, 0, 1) / (o.S * o.V * math.Sqrt(o.T)), nil
}

func (o *Option) Vega() (float64, error) {
	if o.K == 0 {
		return 0, errors.New("strike price must be greater than 0")
	}
	d1 := (math.Log(o.S/o.K) + (o.R+0.5*math.Pow(o.V, 2))*o.T) / (o.V * math.Sqrt(o.T))
	return o.S * standardNormalCDF(d1) * math.Sqrt(o.T), nil
}

func (o *Option) CallTheta() (float64, error) {
	if o.K == 0 {
		return 0, errors.New("strike price must be greater than 0")
	}
	d1 := (math.Log(o.S/o.K) + (o.R+0.5*math.Pow(o.V, 2))*o.T) / (o.V * math.Sqrt(o.T))
	d2 := d1 - o.V*math.Sqrt(o.T)
	return (-o.S*o.V*standardNormalCDF(d1))/(2*math.Sqrt(o.T)) - o.R*o.K*math.Exp(-o.R*o.T)*standardNormalCDF(d2), nil
}
