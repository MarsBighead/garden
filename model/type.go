package model

// Person General people to visit
type Person struct {
	Name    string
	Age     int
	Emails  []string
	Company string
	Role    string
}

// OnlineUser User struct for online party
type OnlineUser struct {
	User      []*Person
	LoginTime string
}

// Statistic struct for handler statistic
type Statistic struct {
	Numbers []float64
	Count   int
	Mean    float64
	Median  float64
	ErrMsg  string
	Get     bool
}
