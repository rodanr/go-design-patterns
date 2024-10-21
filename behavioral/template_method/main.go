package main

import "fmt"

// IOtp defines the template for OTP operations
type IOtp interface {
	genOTP(int) string
	saveOTP(string)
	getMessage(string) string
	send(string) error
}

type Otp struct {
	method IOtp
}

// genAndSendOTP is the template method that executes the steps
func (o *Otp) genAndSendOTP(length int) error {
	otp := o.method.genOTP(length)
	o.method.saveOTP(otp)
	message := o.method.getMessage(otp)
	return o.method.send(message)
}

// Sms is the concrete implementation for sending OTP via SMS
type Sms struct{}

func (s *Sms) genOTP(length int) string {
	return "4321" // simplified random OTP generation
}

func (s *Sms) saveOTP(otp string) {
	fmt.Println("SMS: Saving OTP to cache:", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) send(message string) error {
	fmt.Println("SMS: Sending message:", message)
	return nil
}

// Email is the concrete implementation for sending OTP via Email
type Email struct{}

func (e *Email) genOTP(length int) string {
	return "5678" // simplified random OTP generation
}

func (e *Email) saveOTP(otp string) {
	fmt.Println("EMAIL: Saving OTP to cache:", otp)
}

func (e *Email) getMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (e *Email) send(message string) error {
	fmt.Println("EMAIL: Sending message:", message)
	return nil
}

func main() {
	// Using SMS to send OTP
	sms := &Sms{}
	otp := Otp{method: sms}
	otp.genAndSendOTP(4)

	fmt.Println()

	// Using Email to send OTP
	email := &Email{}
	otp = Otp{method: email}
	otp.genAndSendOTP(4)
}
