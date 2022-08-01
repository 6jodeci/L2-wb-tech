package main

import (
	"fmt"
	"log"
	"state-pattern/pkg"
)

// клиентский вызов
func main() {
	// инициализация машины добавив 1 товар с ценой 10
	vendingMachine := pkg.NewVendingMachine(1, 10)
	// пользователь запрашивает состояние - можно ли купить товар
	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	// пользователь внес деньги
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// выдача товара
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	// добавление товара
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	// запрос товара
	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// оплата товара
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// выдача товара
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
