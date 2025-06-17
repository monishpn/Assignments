package main

import (
	"fmt"
	"os"
)

type PaymentMethod interface {
	Pay(amount float64)
}

type CreditCard struct {
	cardNumber string
}

func (c CreditCard) Pay(amt float64) {
	fmt.Printf("[CREDIT CARD] Paid ₹%.2f using card ending with %s\n", amt, c.cardNumber[len(c.cardNumber)-4:len(c.cardNumber)])

}

type PayPal struct {
	email string
}

func (p PayPal) Pay(amount float64) {
	fmt.Printf("[PAYPAL] Paid ₹%.2f using PayPal account: %s\n", amount, p.email)
}

type UPI struct {
	UPIID string
}

func (u UPI) Pay(amount float64) {
	fmt.Printf("[UPI] Paid ₹%.2f using UPI: %s\n", amount, u.UPIID)
}

// OTPss
func genOTP(s interface{}) {
	switch s.(type) {
	case CreditCard:
		fmt.Println("[CREDIT CARD] OTP sent to registered number")
	case UPI:
		fmt.Println("[UPI] OTP sent to registered device")
	default:
		return
	}
}

func main() {
	var x int = 1
	var amt float64
	var pay PaymentMethod

	for x < 4 && x > 0 {
		fmt.Println("\n\nEnter the Amount to be sent")
		fmt.Scanf("%f", &amt)
		x = 0

		fmt.Println("Enter the Mode of Transaction\n1. Credit Card\n2. Paypal\n3. UPI")
		fmt.Scanf("%d", &x)

		switch x {
		case 1:
			fmt.Println("Enter the Card number")
			var card string
			fmt.Scanf("%s", &card)
			c := CreditCard{card}
			pay = c
		case 2:
			fmt.Println("Enter the Mail ID")
			var id string
			fmt.Scanf("%s", &id)
			p := PayPal{id}
			pay = p
		case 3:
			fmt.Println("Enter the UPI ID")
			var id string
			fmt.Scanf("%s", &id)
			u := UPI{id}
			pay = u
		default:
			x = 4
			fmt.Println("Thank You")
			os.Exit(0)
		}
		fmt.Println("\n")
		genOTP(pay)
		pay.Pay(amt)
	}
}
