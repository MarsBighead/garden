package main

import (
	"fmt"

	"log"
	"net/http"
	"os"

	"garden/statistics/api"
)

func main() {
	http.HandleFunc("/api/pb", apiPb)
	http.HandleFunc("/api/xdu", api.Xdu)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Fail to start server localhost:8080", err)
	}
	fmt.Printf("Server is running on http://localhost:8080")
	select {}

}

// apiPb  Only provide get binary protobuf data
func apiPb(w http.ResponseWriter, req *http.Request) {
	pbData := api.GetPb()
	w.Write(pbData)
}

// checkError -Simplify error return checking
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
