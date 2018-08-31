package main

func main() {
	//income := 12000.00
}

type report struct {
	Income  float64
	Tax2011 *tax
	Tax2018 *tax
}
type tax struct {
	IncomeWithTax float64
	Income        float64
	Tax           float64
}

func (r *report) tax2011(income float64, class int) {
	if class == 0 {
		r.Tax2011 = &tax{
			Income: income,
		}
	} else if class == 1 {
		r.Tax2011 = &tax{
			Income: income,
		}
	}
}
func (r *report) havIncomeWithTax2011() {
	if r.Income*12 > 9600 {

	}

}
func (r *report) havIncome2011() {

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
func tax2011QuickDeduction() map[int]float64 {
	baseline := []int{
		18000,
		54000,
		108000,
		420000,
		660000,
		960000,
		-1,
	}
	rate := tax2011RateByYear()
	tqd2011 := make(map[int]float64)
	var tax, base float64
	for _, v := range baseline {
		if v > 0 {
			tax += (float64(v) - base) * rate[v]
			tqd2011[v] = tax
			base = float64(v)
		}

	}
	tqd2011[-1] = 0.45
	return tqd2011
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

func tax2011RateByYear() map[int]float64 {
	return map[int]float64{
		18000:  0.03,
		54000:  0.10,
		108000: 0.20,
		420000: 0.25,
		660000: 0.30,
		960000: 0.35,
		-1:     0.45,
	}
}
