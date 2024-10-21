package main

import (
	"fmt"
	"log"
)

// Facade: WalletFacade simplifies complex subsystems interactions
type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Setting up account and security code")
	return &WalletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
	}
}

func (w *WalletFacade) addMoney(amount int) error {
	err := w.account.checkAccount()
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode()
	if err != nil {
		return err
	}
	w.wallet.add(amount)
	return nil
}

func (w *WalletFacade) deductMoney(amount int) error {
	err := w.account.checkAccount()
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode()
	if err != nil {
		return err
	}
	return w.wallet.deduct(amount)
}

// Subsystem: Account
type Account struct {
	id string
}

func newAccount(id string) *Account {
	return &Account{id: id}
}

func (a *Account) checkAccount() error {
	fmt.Println("Account verified")
	return nil
}

// Subsystem: SecurityCode
type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) checkCode() error {
	fmt.Println("Security code verified")
	return nil
}

// Subsystem: Wallet
type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) add(amount int) {
	w.balance += amount
	fmt.Printf("Added %d to wallet. Current balance: %d\n", amount, w.balance)
}

func (w *Wallet) deduct(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("not enough balance")
	}
	w.balance -= amount
	fmt.Printf("Deducted %d from wallet. Current balance: %d\n", amount, w.balance)
	return nil
}

// Client Code
func main() {
	facade := newWalletFacade("abc", 1234)

	err := facade.addMoney(100)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	err = facade.deductMoney(50)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
