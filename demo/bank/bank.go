package bank

import (
	"fmt"
)

var balance int

func Deposit(amount int, name string) {
	balance = balance + amount
	fmt.Println(name, "=", Balance())
}

func Balance() int { return balance }
