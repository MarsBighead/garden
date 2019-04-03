package page

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

//PageList list all handle in the page
func (s *Service) PageList(w http.ResponseWriter, r *http.Request) {
	log.Print("Running http handle model.HomeList!")
	tpl, err := template.ParseFiles(s.Environment["TEMPLATE"] + "/list.htm")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}

// PaddingTemplate  build web page with template from the set value
func (s *Service) PaddingTemplate(w http.ResponseWriter, r *http.Request) {
	zoro := Person{
		Name:    "zoro",
		Age:     27,
		Emails:  []string{"dg@gmail.com", "dk@hotmail.com"},
		Company: "Omron",
		Role:    "SE"}

	zoe := Person{
		Name:   "zoe",
		Age:    26,
		Emails: []string{"test@gmail.com", "d@hotmail.com"}}

	onlineUser := OnlineUser{User: []*Person{&zoro, &zoe}}

	log.Print("Running http handle model.HomeTemplate!")

	t, err := template.ParseFiles(s.Environment["TEMPLATE"] + "/template/tpl.htm")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, onlineUser)
	if err != nil {
		log.Fatal(err)
	}
}

// ProtocalHTTP method for test http protocal output in dashboard
func ProtocalHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse parameters, default none
	fmt.Println(r.Form)
	fmt.Println("User-Agent:", r.Header.Get("User-Agent"))
	fmt.Println("HTTP scheme:", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	var content = "Hello!\n" + "<a href=\"\">matrix API</a>"

	w.Write([]byte(content)) //这个写入到w的是输出到客户端的
}

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistic</title>
<body><h3>Statistic</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/statistic" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
	pageBottom = `</body></html>`
	anError    = `<p class="error">%s</p>`
)

// AdvancedStatistic Advanced statistic with template http page
func (s *Service) AdvancedStatistic(writer http.ResponseWriter, req *http.Request) {
	fmt.Printf("POST method is %s\n", req.Method)
	tpl, err := template.ParseFiles(s.Environment["TEMPLATE"] + "/statistics.htm")
	checkError(err)
	err = req.ParseForm() // Must be called before writing response
	checkError(err)
	if req.Method == "GET" {
		err = tpl.Execute(writer, "")
	} else if req.Method == "POST" {
		if numbers, message, ok := ProcessRequest(req); ok {
			stats := getStatistic(numbers)
			fmt.Printf("Calculate result is %v.\n", stats)
			err = tpl.Execute(writer, stats)
		} else if message != "" {
			fmt.Printf("Error msg is: %s.\n", message)
			stats := Statistic{
				Get:    false,
				ErrMsg: message,
			}
			err = tpl.Execute(writer, stats)
		}
	}
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

// HomeStatistic Statistic number
func HomeStatistic(writer http.ResponseWriter, req *http.Request) {
	fmt.Printf("HomePage handler from controller\n")
	fmt.Printf("HTTP Method: %s.\n", req.Method)
	err := req.ParseForm() // Must be called before writing response
	fmt.Printf("Initialize homePage form: %v.\n", req.ParseForm())
	fmt.Fprint(writer, pageTop, form)
	if err != nil {
		fmt.Fprintf(writer, anError, err)
	} else {
		if numbers, message, ok := ProcessRequest(req); ok {
			stats := getStatistic(numbers)
			fmt.Fprint(writer, formatStats(stats))
		} else if message != "" {
			fmt.Fprintf(writer, anError, message)
		}
	}
	fmt.Fprint(writer, pageBottom)
}

// ProcessRequest  Management request from client
func ProcessRequest(request *http.Request) ([]float64, string, bool) {
	var numbers []float64
	if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
		text := strings.Replace(slice[0], ",", " ", -1)
		for _, field := range strings.Fields(text) {
			n, err := strconv.ParseFloat(field, 64)
			if err != nil {
				return numbers, "'" + field + "' is invalid", false
			}
			numbers = append(numbers, n)
		}
	}
	fmt.Printf("Get numbers %v in controller.\n", numbers)
	if len(numbers) == 0 {
		return numbers, "No data input!", false // no data first time form is shown
	}
	return numbers, "", true
}

func formatStats(stats Statistic) string {
	return fmt.Sprintf(`<table border="1">
        <tr><th colspan="2">Results</th></tr>
        <tr><td>Numbers</td><td>%v</td></tr>
        <tr><td>Count</td><td>%d</td></tr>
        <tr><td>Mean</td><td>%f</td></tr>
        <tr><td>Median</td><td>%f</td></tr>
        </table>`, stats.Numbers, len(stats.Numbers), stats.Mean, stats.Median)
}

// getStatistic Get statistic result
func getStatistic(numbers []float64) (stats Statistic) {
	stats.Numbers = numbers
	stats.Count = len(numbers)
	sort.Float64s(stats.Numbers)
	stats.Mean = sum(numbers) / float64(len(numbers))
	stats.Median = median(numbers)
	stats.Get = true
	return stats
}

func sum(numbers []float64) (total float64) {
	if len(numbers) == 0 {
		return 0.0
	}
	for _, x := range numbers {
		total += x
	}
	return total
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

// checkError -Simplify error return checking
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
