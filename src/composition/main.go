package composition

import "fmt"

type Account struct {
}

func (a *Account) AvailableFunds() float32 {
	fmt.Println("Listing available founds")
	return 0
}
func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment")
	return true
}

type CreditAccount struct {
	Account
}

func main() {
	creditAccount := &CreditAccount{}
	creditAccount.AvailableFunds()
	creditAccount.ProcessPayment(500)
}
