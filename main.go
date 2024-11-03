package main

import (
	account "bank/accounts"
	client "bank/clients"
	"fmt"
)

func PayBill(account verifyAccount, value float64) {
	account.Withdraw(value)
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {

	p1 := client.Holder{Name: "Person 1", CPF: "123.123.123-12", Occupation: "Developer"} // Account 1's holder.
	a1 := account.CurrentAccount{Holder: p1, AgencyNumber: 123, AccountNumber: 123456}    // Creating Person 1's Current Account.
	a1.Deposit(1000)                                                                      // Deposit method.
	fmt.Println("Person 1's balance: $", a1.GetBalance())                                 // GetBalance method.
	PayBill(&a1, 60)                                                                      // PayBill method.
	a1.Withdraw(250)                                                                      // Withdraw method.
	fmt.Println("Person 1's new balance: $", a1.GetBalance())                             // New balance

	fmt.Println("")

	p2 := client.Holder{Name: "Person 2", CPF: "456.456.456-45", Occupation: "Developer"}              // Account 2's holder.
	a2 := account.SavingsAccount{Holder: p2, AgencyNumber: 456, AccountNumber: 456789, Operation: 200} // Creating Person 2's Savings Account.
	a2.Deposit(500)                                                                                    // Deposit method.
	fmt.Println("Person 2's balance: $", a2.GetBalance())                                              // GetBalance method.
	PayBill(&a2, 30)                                                                                   // PayBill method.
	a2.Withdraw(250)                                                                                   // Withdraw method.
	fmt.Println("Person 2's new balance: $", a2.GetBalance())                                          // New balance

	fmt.Println("")

	p3 := client.Holder{Name: "Person 3", CPF: "135.135.135-13", Occupation: "Developer"} // Account 3's holder.
	a3 := account.CurrentAccount{Holder: p3, AgencyNumber: 135, AccountNumber: 135791}    // Creating Person 3's Current Account.
	a3.Deposit(2500)                                                                      // Deposit method.
	a3.Transfer(900, &a1)                                                                 // Transfer method.
	fmt.Println("Person 3's new balance: $", a3.GetBalance())                             // Person 3's new balance
	fmt.Println("Person 1's new balance: $", a1.GetBalance())                             // Person 's new balance
}
