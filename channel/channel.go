package main

import (
    "fmt"
    "time"
)

var c chan int

func ready(w string, sec int){
    time.Sleep(time.Duration(sec)*time.Second)
    fmt.Println(w," is ready")
     c<-1
}


func main() {
    c = make(chan int)
    go ready("Fruit Juice", 2)
    go ready("Tea", 2)
    go ready("Coffee", 1)
    fmt.Println("I'm waiting, but not too long!")
    <-c
    <-c
    <-c
}
