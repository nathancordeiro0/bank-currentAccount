package accounts

import "bank/clients"

type SavingsAccount struct {
	Holder                                 clients.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingsAccount) Withdraw(value float64) string {
	canWithdraw := value > 0 && value <= c.balance

	if canWithdraw {
		c.balance -= value
		return "Successful withdrawal"
	} else {
		return "Insufficient balance."
	}
}

func (c *SavingsAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Successful deposit. New balance: $", c.balance
	} else {
		return "The deposit amount is less than zero. Balance: R$", c.balance
	}
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
