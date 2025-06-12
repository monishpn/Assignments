package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func exchange(a float32, cur, tgt string) float32 {
	var usd float32 = 85.45
	var eur float32 = 98.11
	var jpy float32 = 0.59

	cur = strings.ToUpper(cur)
	tgt = strings.ToUpper(tgt)

	if cur == "USD" {
		a *= usd
	} else if cur == "JPY" {
		a *= jpy
	} else if cur == "EUR" {
		a *= eur
	} else if cur == "INR" {
		a *= 1
	} else {
		fmt.Printf("Invalid currency value, %s not in the databse", cur)
		os.Exit(0)
	}

	if tgt == "INR" {
		return a
	} else if tgt == "USD" {
		return a / usd
	} else if tgt == "JPY" {
		return a / jpy
	} else if tgt == "EUR" {
		return a / eur
	}

	fmt.Printf("Invalid currency value, %s not in the databse", tgt)
	os.Exit(0)
	return 0

}

func main() {
	var a float32
	var cur, tgt string

	now := time.Now().Hour()

	if now >= 6 && now < 12 {
		fmt.Println("Good Morning!!!!")
	} else if now >= 12 && now < 16 {
		fmt.Println("Good Afternoon!!!")
	} else {
		fmt.Println("Good Evening!!!")
	}

	fmt.Print("Enter the amount : ")
	_, errA := fmt.Scanln(&a)
	if errA != nil {
		fmt.Print("Invalid input")
		os.Exit(0)
	}

	fmt.Print("Enter the currency : ")
	_, errA = fmt.Scanln(&cur)
	if errA != nil {
		fmt.Print("Invalid input")
		os.Exit(0)
	}

	fmt.Print("Enter the Target currency : ")
	_, errA = fmt.Scanln(&tgt)
	if errA != nil {
		fmt.Print("Invalid input")
		os.Exit(0)
	}

	fmt.Printf("%.2f %s in %s is %.2f %s", a, cur, tgt, exchange(a, cur, tgt), tgt)

}
