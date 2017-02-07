package main

import (
    "fmt"
    "strings"
    "strconv"
    "log"
    "os"
    "html/template"
    "net/http"

    "garden/statistics/api"
    cst "garden/controllers/statistics"

)


func main() {
    http.HandleFunc("/", cst.HomePage)
    http.HandleFunc("/t", testPage)
    http.HandleFunc("/api/pb", apiPb)
    if err :=http.ListenAndServe(":8080", nil); err!=nil {
        log.Fatal("Fail to start server localhost:8080", err)
    }
    fmt.Printf("Server is running on http://localhost:8080")
    select{};

}
func testPage(writer http.ResponseWriter, req *http.Request) {
    templ, err := template.ParseFiles("template/test.htm")  
    checkError(err) 
    err = req.ParseForm() // Must be called before writing response
    checkError(err) 
    if req.Method == "GET" {  
        err = templ.Execute(writer,"") 
    } else if req.Method == "POST" { 
        fmt.Printf("POST method verified is %s\n",req.Method)
        if numbers, message, ok := cst.ProcessRequest(req); ok {
            stats := cst.GetStats(numbers)
            fmt.Printf("Calculate result is %v.\n",stats)
            err = templ.Execute(writer, stats)
        } else if message != "" {
            fmt.Printf("Error msg is: %s.\n",message)
            stats := cst.Statistics{
                 Get: false,
                 ErrMsg: message,
            }
            err = templ.Execute(writer, stats)
        }
    }
}

// apiPb  Provide api with protobuf
func apiPb(writer http.ResponseWriter, req *http.Request) {
     api.GetPb(writer) 
}

//func processRequest(request *http.Request) ([]float64, string, bool) {
//    var numbers []float64
//    if slice, found := request.Form["numbers"]; found && len(slice) > 0 {
//        text := strings.Replace(slice[0], ",", " ", -1)
//        for _, field := range strings.Fields(text) {
//            if x, err := strconv.ParseFloat(field, 64); err != nil {
//                return numbers, "'" + field + "' is invalid", false
//            } else {
//                numbers = append(numbers, x)
//            }
//        }
//    }
//    fmt.Printf("Numbers not in controller %v.\n", numbers)
//    if len(numbers) == 0 {
//        return numbers, "No data input!", false // no data first time form is shown
//    }
//    return numbers, "", true
//}

// checkError -Simplify error return checking
func checkError(err error) {
     if err != nil {
         fmt.Println("Fatal error ", err.Error())
      os.Exit(1)
   }
}
