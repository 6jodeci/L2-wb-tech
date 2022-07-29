// Реализация поведения когда карта является инициатором запроса в банк

package pkg

import "time"

// объект банкоская карта которая включает в себя имя, баланс и банк к которому она относится
type Card struct {
	Name string 
	Balance float64
	Bank *Bank 
}
// CheckBalance() проверяет баланс банковской карты покупателя
func (card Card) CheckBalance() error {
	println("[Карта] Запрос в банк для проверки баланса..")
	// ожидание ответа сервера
	time.Sleep(time.Second *5)
	// запрашиваем информацию о балансе
	return card.Bank.CheckBalance(card.Name)
}