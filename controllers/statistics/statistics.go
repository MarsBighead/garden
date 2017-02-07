package statistics 

import (
    "fmt"
    "strings"
    "strconv"
    "os"
    "sort"
    "net/http"
)

type Statistics struct {
    Numbers [] float64
    Count      int
    Mean       float64
    Median     float64
    ErrMsg     string 
    Get        bool
}
const (
    pageTop    = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
    form       = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
    pageBottom = `</body></html>`
    anError    = `<p class="error">%s</p>`
)


func HomePage(writer http.ResponseWriter, req *http.Request) {
    fmt.Printf("HomePage handler from controller\n")
    fmt.Printf("HTTP Method: %s.\n",req.Method)
    err := req.ParseForm() // Must be called before writing response
    fmt.Printf("Initialize homePage form: %v.\n",req.ParseForm())
    fmt.Fprint(writer, pageTop, form)
    if err != nil {
        fmt.Fprintf(writer, anError, err)
    } else {
        if numbers, message, ok := ProcessRequest(req); ok {
            stats := GetStats(numbers)
            fmt.Fprint(writer, formatStats(stats))
        } else if message != "" {
            fmt.Fprintf(writer, anError, message)
        }
    }
    fmt.Fprint(writer, pageBottom)
}

func ProcessRequest(request *http.Request) ([]float64, string, bool) {
    var numbers []float64
    if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
        text := strings.Replace(slice[0], ",", " ", -1)
        for _, field := range strings.Fields(text) {
            if x, err := strconv.ParseFloat(field, 64); err != nil {
                return numbers, "'" + field + "' is invalid", false
            } else {
                numbers = append(numbers, x)
            }
        }
    }
    fmt.Printf("Get numbers %v in controller.\n", numbers)
    if len(numbers) == 0 {
        return numbers, "No data input!", false // no data first time form is shown
    }
    return numbers, "", true
}

func formatStats(stats Statistics) string {
    return fmt.Sprintf(`<table border="1">
        <tr><th colspan="2">Results</th></tr>
        <tr><td>Numbers</td><td>%v</td></tr>
        <tr><td>Count</td><td>%d</td></tr>
        <tr><td>Mean</td><td>%f</td></tr>
        <tr><td>Median</td><td>%f</td></tr>
        </table>`, stats.Numbers, len(stats.Numbers), stats.Mean, stats.Median)
}

func GetStats(numbers [] float64) (stats Statistics) {
    stats.Numbers = numbers
    stats.Count = len(numbers)
    sort.Float64s(stats.Numbers)
    stats.Mean = sum(numbers)/ float64(len(numbers))
    stats.Median = median(numbers)
    stats.Get = true
    return stats  
}

func sum(numbers [] float64)(total float64){
    if len(numbers) == 0 {
        return  0.0
    }
    for _,x := range numbers {
        total +=x
    }
    return total
}

func median(numbers [] float64) float64{
    middle := len(numbers)/2
    result := numbers[middle]
    if len(numbers)%2==0 {
        result = (result + numbers[middle-1])/2
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
