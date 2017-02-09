package main

import (
    "fmt"

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
    http.HandleFunc("/api/xdu", api.Xdu)
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

// apiPb  Only provide get binary protobuf data
func apiPb(w http.ResponseWriter, req *http.Request) {
    pbData:=api.GetPb()
    w.Write(pbData)
}

// checkError -Simplify error return checking
func checkError(err error) {
     if err != nil {
         fmt.Println("Fatal error ", err.Error())
      os.Exit(1)
   }
}
