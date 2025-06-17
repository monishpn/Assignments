package main

import (
	"fmt"
	"os"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (acc BankAccount) Display() {
	fmt.Printf("BankAccount Owner: %s\nBank Balance: %0.2f\n\n", acc.Owner, acc.Balance)
}

func (acc *BankAccount) Deposit(amt float64) {
	acc.Balance += amt
}

func (acc *BankAccount) Withdraw(amt float64) {
	if acc.Balance < amt {
		fmt.Println("Error: Cannot withdraw, There is no enough money")
		return
	}
	acc.Balance -= amt

}

func main() {
	x := 2

	acc := &BankAccount{"Monish", 1000}

	for x != 4 {

		fmt.Println("1. Deposit Money")
		fmt.Println("2. Display Balance")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Enter Your Choice : ")
		_, err := fmt.Scanf("%d", &x)
		if err != nil {
			fmt.Println("Invalid Choice")
			x = 4
			break
		}
		fmt.Println()

		switch x {
		case 1:
			fmt.Println("Enter The Amount to Deposit : ")
			var dep float64
			fmt.Scanf("%f", &dep)
			acc.Deposit(dep)
		case 2:
			acc.Display()
		case 3:
			fmt.Println("Enter The Amount to Withdraw : ")
			var wit float64
			fmt.Scanf("%f", &wit)
			acc.Withdraw(wit)
		default:
			x = 4
			fmt.Println("Thank You.....!!!!")
			os.Exit(1)
		}

	}
}
