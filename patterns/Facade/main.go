// Паттерн «Фасад».
// Пример когда приходит пользователь и пытается купить товар в магазине
package main

import (
	"facade-pattern/pkg"
)

// создание нужных данных
var (
	bank = pkg.Bank{
		Name:  "Bank",
		Cards: []pkg.Card{},
	}
	card1 = pkg.Card{
		Name:    "BANK-CARD1",
		Balance: 300,
		Bank:    &bank,
	}
	card2 = pkg.Card{
		Name:    "BANK-CARD2",
		Balance: 5,
		Bank:    &bank,
	}
	user1 = pkg.User{
		Name: "Покупатель-1",
		Card: &card1,
	}
	user2 = pkg.User{
		Name: "Покупатель-2",
		Card: &card2,
	}
	prod = pkg.Product{
		Name:  "Coca-Cola",
		Price: 150,
	}
	shop = pkg.Shop{
		Name: "Drink Shop",
		Products: []pkg.Product{
			prod,
		},
	}
)

func main() {
	println("[Банк] Выпуск карт..")
	bank.Cards = append(bank.Cards, card1, card2)
	// первый покупатель
	println("[", user1.Name, "]")
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
	// второй покупатель
	println("[", user2.Name, "]")
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
}