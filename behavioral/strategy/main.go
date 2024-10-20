package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount int) string
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount int) string {
	return fmt.Sprintf("Paid %d using Credit Card", amount)
}

type Paypal struct{}

func (p *Paypal) Pay(amount int) string {
	return fmt.Sprintf("Paid %d using Paypal", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount int) string {
	return p.strategy.Pay(amount)
}

func main() {
	paymentContext := PaymentContext{}

	paymentContext.SetStrategy(&CreditCard{})
	fmt.Println(paymentContext.Pay(100))

	paymentContext.SetStrategy(&Paypal{})
	fmt.Println(paymentContext.Pay(200))
}
