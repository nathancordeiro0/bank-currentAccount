package main

import "fmt"

type CurrentAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func (c *CurrentAccount) Withdraw(value float64) string {
	canWithdraw := value <= c.balance

	if canWithdraw {
		c.balance -= value
		return "Saque realizado com sucesso."
	} else {
		return "Saldo insuficiente."
	}
}

func main() {
	p1 := CurrentAccount{}
	p1.holder = "Nathan"
	p1.balance = 500

	fmt.Println(p1.holder, p1.balance)

	fmt.Println(p1.Withdraw(-100))
	fmt.Println(p1.holder, p1.balance)
}
