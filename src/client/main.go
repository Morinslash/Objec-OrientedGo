package main

import (
	"Objec-OrientedGo/src/payment"
)

func main() {
	var option payment.PaymentOption

	option = payment.CreateCreditCard(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	option.ProcessPayment(500)
	option = payment.CreateCashAccount()
	option.ProcessPayment(500)
}
