// Реализация поведения магазина когда приходит покупатель и хочет купить товар
package pkg

import (
	"time"
	"errors"
)

// объект магазин который включает в себя название и ассортимент товаров для покупки
type Shop struct {
	Name string 
	Products []Product
}
// !!!ОСНОВНОЙ ФАСАД!!!
// Sell(), реализует поведение когда покупатель покупает товар
func (shop Shop) Sell(user User, product string) error {
	// логика поведения когда пользователь приходити покупает какой-то товар
	println("[Магазин] Запрос к пользователю, для получения остатка по карте..")
	// ожидание ответа сервера
	time.Sleep(time.Second *5)
	err := user.Card.CheckBalance()
	// если баланс отрицательный или какая-то другая ошибка
	if err != nil {
		return err
	}
	println("[Магазин] Проверка - может ли пользователь", user.Name, "оплатить покупку..")
	// ожидание ответа сервера
	time.Sleep(time.Second *5)
	// проходим по всем товарам магазина
	for _, prod := range shop.Products {
		// проверка названий продута
		if prod.Name != product {
			// пропускаем продукт
			continue
		}
		// если недостаточно средств возвращаем ошибку о недостатке баланса
		if prod.Price > user.GetBalance() {
			println("[Магазин] Недостаточно средств для покупки товара", prod.Name)
			return errors.New("")
		}
		println("[Магазин] Покупка прошла успешно!", prod.Name)
	}
	return nil
}