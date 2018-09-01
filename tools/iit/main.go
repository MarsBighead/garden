package main

import (
	"fmt"
)

func main() {
	income := 12000.00
	flag := 1 //without tax

	/*income := 5000.00
	flag := 0 //with tax*/
	r := new(report)
	r.tax2011((income-3500)*12, flag)
	fmt.Printf(`
China Individual Tax Version 2011
Before Tax: %f	
After Tax: %f	
Tax: %f	
`, r.Tax2011.IncomeWithTax, r.Tax2011.IncomeWithoutTax, r.Tax2011.Tax)
	income = r.Tax2011.IncomeWithTax
	r.tax2018((income-5000)*12, 0)
	fmt.Printf(`
China Individual Tax Version 2018
Before Tax: %f	
After Tax: %f	
Tax: %f	
`, r.Tax2018.IncomeWithTax, r.Tax2018.IncomeWithoutTax, r.Tax2018.Tax)

	fmt.Printf(`
China Individual Tax Version 2018
Tax reduction amount: %f
`, r.Tax2011.Tax-r.Tax2018.Tax)

}

type report struct {
	Income    float64
	Tax2011   *summary
	Tax2018   *summary
	rates     map[float64]float64
	beforeTax []float64
	baseline  float64
}
type summary struct {
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

	rates := tax2011Rate()
	r.beforeTax = beforeTax
	r.rates = rates
	r.baseline = 3500
	if income <= 0 {
		r.Tax2011 = &summary{
			IncomeWithoutTax: income/12 + r.baseline,
			IncomeWithTax:    income/12 + r.baseline,
			Tax:              0,
		}
	} else {
		//fmt.Printf("income: %f\n", income)
		r.Tax2011 = r.compute(flag, income)
	}
}
func (r *report) tax2018(income float64, flag int) {
	beforeTax := []float64{
		36000,
		144000,
		300000,
		420000,
		660000,
		960000,
		-1,
	}
	rate := tax2018Rate()
	r.beforeTax = beforeTax
	r.rates = rate
	r.baseline = 5000
	if income <= 0 {
		r.Tax2018 = &summary{
			IncomeWithoutTax: income/12 + r.baseline,
			IncomeWithTax:    income/12 + r.baseline,
			Tax:              0,
		}
	} else {
		r.Tax2018 = r.compute(flag, income)
	}
}

func (r *report) compute(flag int, income float64) *summary {
	beforeTax := r.beforeTax
	rate := r.rates
	var base float64
	s := new(summary)
	if flag == 0 {
		s.IncomeWithTax = income
		_, qdWithTax := quickDeduction(flag, beforeTax, rate)
		for _, v := range beforeTax {
			//fmt.Printf("income:%f, v:%f, base: %f\n", income, v, base)
			if v >= income && income > base && base >= 0 {
				t := qdWithTax[v]
				s.Tax = t.Tax - (v-income)*t.Rate
				s.IncomeWithoutTax = s.IncomeWithTax - s.Tax
				break
			}
			base = v
		}
		if base == -1 {
			base = beforeTax[len(beforeTax)-2]
			t := qdWithTax[base]
			s.Tax = t.Tax + rate[-1]*(income-base)
			s.IncomeWithoutTax = s.IncomeWithTax - s.Tax
		}
	} else if flag == 1 {
		s.IncomeWithoutTax = income
		afterTax, qdWithoutTax := quickDeduction(flag, beforeTax, rate)
		for _, v := range afterTax {
			if v >= income && income > base {
				t := qdWithoutTax[v]
				s.Tax = t.Tax - (v-income)*t.Rate/(1-t.Rate)
				s.IncomeWithTax = s.Tax + income
				break
			}
			base = v
		}
		if base == -1 {
			base = afterTax[len(afterTax)-2]
			t := qdWithoutTax[base]
			s.Tax = t.Tax + (income-base)*rate[-1]/(1-rate[-1])
			s.IncomeWithTax = s.IncomeWithoutTax + s.Tax
		}

	}
	s.IncomeWithTax = s.IncomeWithTax/12 + r.baseline
	s.IncomeWithoutTax = s.IncomeWithoutTax/12 + r.baseline
	s.Tax = s.Tax / 12
	return s
}

func tax2018Rate() map[float64]float64 {
	return map[float64]float64{
		36000:  0.03,
		144000: 0.10,
		300000: 0.20,
		420000: 0.25,
		660000: 0.30,
		960000: 0.35,
		-1:     0.45,
	}
}

func quickDeduction(flag int, beforeTax []float64, rate map[float64]float64) ([]float64, map[float64]summary) {
	qdTax := make(map[float64]summary)
	var taxAmount, amount float64
	var afterTax []float64
	for _, base := range beforeTax {
		if base > 0 {
			taxAmount += (base - amount) * rate[base]
			v := summary{
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
	qdTax[-1] = summary{
		Rate: 0.45,
	}

	return afterTax, qdTax
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
