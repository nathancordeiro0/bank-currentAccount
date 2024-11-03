package accounts

import "bank/clients"

type CurrentAccount struct {
	Holder                      clients.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (c *CurrentAccount) Withdraw(value float64) string {
	canWithdraw := value > 0 && value <= c.balance

	if canWithdraw {
		c.balance -= value
		return "Successful withdrawal"
	} else {
		return "Insufficient balance."
	}
}

func (c *CurrentAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Successful deposit. New balance: $", c.balance
	} else {
		return "The deposit amount is less than zero. Balance: R$", c.balance
	}
}

func (c *CurrentAccount) Transfer(value float64, target *CurrentAccount) bool {
	if value < c.balance && value > 0 {
		c.balance -= value
		target.Deposit(value)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
