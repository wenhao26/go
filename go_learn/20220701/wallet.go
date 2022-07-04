package main

import (
	"fmt"
	"net/http"
	"os"
	_ "net/http/pprof"
)

type Card struct {
	Cmd     string
	Amount  float64
	Balance float64
}

func InitAccount() *Card {
	return &Card{
		Cmd:     "",
		Amount:  0,
		Balance: 1000,
	}
}

func (card *Card) Details() {
	fmt.Println("...[OK] The balance is ", card.Balance)
}

func (card *Card) Income() {
	fmt.Println("[Tips]Enter revenue amount:")
	_, _ = fmt.Scanln(&card.Amount)
	card.Balance += card.Amount
	fmt.Println("...[OK] Income ", card.Amount, " the balance is", card.Balance)
}

func (card *Card) Pay() {
	fmt.Println("...[Tips]Enter expenditure amount:")
	_, _ = fmt.Scanln(&card.Amount)
	if card.Balance-card.Amount < 0 {
		fmt.Println("...[Err]Sorry, your credit is running low!")
		return
	}
	card.Balance -= card.Amount
	fmt.Println("...[OK] expenditure ", card.Amount, " the balance is", card.Balance)
}

func (card *Card) Exit() {
	fmt.Println("Confirm exit? y/n")

	input := ""
	for {
		_, _ = fmt.Scanln(&input)
		if input == "Y" {
			os.Exit(0)
		}
	}
}

func (card *Card) Menu() {
	for {
		fmt.Println("\n********* User Wallet *********")
		fmt.Println("Option: enter the corresponding digital operation...")
		fmt.Println("[1]Details [2]Income [3]Pay [4]Exit\n")
		fmt.Println("Input:")

		_, _ = fmt.Scanln(&card.Cmd)
		switch card.Cmd {
		case "1":
			card.Details()
			break
		case "2":
			card.Income()
			break
		case "3":
			card.Pay()
		case "4":
			card.Exit()
		default:
			fmt.Println("Invalid operation!")
		}
	}
}

func main() {
	//http.ListenAndServe("0.0.0.0:6060", nil)
	wallet := InitAccount()
	wallet.Menu()


}
