package main

import (
	"fmt"
)

func main() {
	income := 66505.00
	flag := 1 //without tax

	//income := 93500.00
	//flag := 0 //with tax
	r := new(report)
	r.tax2011((income-3500)*12, flag)
	fmt.Printf(`
China Individual Tax Version 2011
Before Tax: %f	
After Tax: %f	
Tax: %f	
`, r.Tax2011.IncomeWithTax, r.Tax2011.IncomeWithoutTax, r.Tax2011.Tax)

}

type report struct {
	Income  float64
	Tax2011 *tax
	Tax2018 *tax
}
type tax struct {
	IncomeWithTax    float64
	IncomeWithoutTax float64
	Tax              float64
	Rate             float64
}

func (r *report) tax2011(income float64, flag int) {
	beforeTax := []float64{
		18000,
		54000,
		108000,
		420000,
		660000,
		960000,
		-1,
	}

	rate := tax2011Rate()
	var base float64
	if flag == 0 {
		r.Tax2011 = &tax{
			IncomeWithTax: income,
		}
		_, qdWithTax := quickDeduction(flag, beforeTax, rate)
		for _, v := range beforeTax {
			if v >= income && income > base && base > 0 {
				t := qdWithTax[v]
				r.Tax2011.Tax = t.Tax - (v-income)*t.Rate
				r.Tax2011.IncomeWithoutTax = r.Tax2011.IncomeWithTax - r.Tax2011.Tax
				break
			}
			base = v
		}
		if base == -1 {
			base = beforeTax[len(beforeTax)-2]
			t := qdWithTax[base]
			r.Tax2011.Tax = t.Tax + rate[-1]*(income-base)
			r.Tax2011.IncomeWithoutTax = r.Tax2011.IncomeWithTax - r.Tax2011.Tax
		}
	} else if flag == 1 {
		r.Tax2011 = &tax{
			IncomeWithoutTax: income,
		}
		afterTax, qdWithoutTax := quickDeduction(flag, beforeTax, rate)
		for _, v := range afterTax {
			if v >= income && income > base {
				t := qdWithoutTax[v]
				r.Tax2011.Tax = t.Tax - (v-income)*t.Rate/(1-t.Rate)
				r.Tax2011.IncomeWithTax = r.Tax2011.Tax + income
				break
			}
			base = v
		}
		if base == -1 {
			base = afterTax[len(afterTax)-2]
			t := qdWithoutTax[base]
			r.Tax2011.Tax = t.Tax + (income-base)*rate[-1]/(1-rate[-1])
			r.Tax2011.IncomeWithTax = r.Tax2011.IncomeWithoutTax + r.Tax2011.Tax
		}

	}
	r.Tax2011.IncomeWithTax = r.Tax2011.IncomeWithTax/12 + 3500
	r.Tax2011.IncomeWithoutTax = r.Tax2011.IncomeWithoutTax/12 + 3500
	r.Tax2011.Tax = r.Tax2011.Tax / 12
}

func tax2018ByYear() map[int]float64 {
	return map[int]float64{
		36000:  0.03,
		144000: 0.10,
		300000: 0.20,
		420000: 0.25,
		660000: 0.30,
		960000: 0.35,
		-1:     0.45,
	}
}

func quickDeduction(flag int, beforeTax []float64, rate map[float64]float64) ([]float64, map[float64]tax) {
	qdTax := make(map[float64]tax)
	var taxAmount, amount float64
	var afterTax []float64
	for _, base := range beforeTax {
		if base > 0 {
			taxAmount += (base - amount) * rate[base]
			v := tax{
				IncomeWithoutTax: (base - taxAmount),
				IncomeWithTax:    base,
				Tax:              taxAmount,
				Rate:             rate[base],
			}
			if flag == 0 {
				qdTax[v.IncomeWithTax] = v
			} else {
				qdTax[v.IncomeWithoutTax] = v
				afterTax = append(afterTax, v.IncomeWithoutTax)
			}
			amount = base
		} else {
			afterTax = append(afterTax, -1)

		}

	}
	qdTax[-1] = tax{
		Rate: 0.45,
	}

	return afterTax, qdTax
}

func tax2018QuickDeduction() map[int]float64 {
	//rates := tax2011RateByYear()
	baseline := []int{
		36000,
		144000,
		300000,
		420000,
		660000,
		960000,
		-1,
	}
	rate := tax2018RateByYear()
	tqd2018 := make(map[int]float64)
	var tax, base float64
	for _, v := range baseline {
		if v > 0 {
			tax += (float64(v) - base) * rate[v]
			tqd2018[v] = tax
			base = float64(v)
		}

	}
	tqd2018[-1] = 0.45
	return tqd2018
}
func tax2018RateByYear() map[int]float64 {
	return map[int]float64{
		36000:  0.03,
		144000: 0.10,
		300000: 0.20,
		420000: 0.25,
		660000: 0.30,
		960000: 0.35,
		-1:     0.45,
	}
}

func tax2011Rate() map[float64]float64 {
	return map[float64]float64{
		18000:  0.03,
		54000:  0.10,
		108000: 0.20,
		420000: 0.25,
		660000: 0.30,
		960000: 0.35,
		-1:     0.45,
	}
}
