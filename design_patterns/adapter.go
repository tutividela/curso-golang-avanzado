package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct {
}

func (CashPayment) Pay() {
	fmt.Printf("Payment using Cash")
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}
func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Pay using Bank Account %d\n",bankAccount)
}
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	accountNumber int
}
func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.accountNumber)
}

func main(){
	cash := CashPayment{}
	ProcessPayment(cash)

	bank := BankPaymentAdapter{
		BankPayment: &BankPayment{},
		accountNumber: 5,
	}
	ProcessPayment(&bank)
}