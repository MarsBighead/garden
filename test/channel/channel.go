package channel

import (
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, " is ready")
	c <- 1
}

func nonBuffer() {
	c = make(chan int, 2)
	for i := 0; i < 10; i++ {
		go ready("Fruit Juice", 2)
		go ready("Tea", 2)
		go ready("Coffee", 1)
		if i == 6 {
			go ready("Add Tea", 2)
			<-c
		}
		fmt.Println("Customer ID:", i, "I'm waiting, but not too long!")
		<-c
		<-c
		<-c
	}

}
