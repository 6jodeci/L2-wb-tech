// Реализация поведения банка когда в банк приходит запрос по номеру карты для проверки баланса  

package pkg

import (
	"fmt"
	"time"
	"errors"
)

// объект банк который включает в себя название и выпущенные банковские карты
type Bank struct {
	Name string
	Cards []Card
}
// CheckBalance банк получает запрос о получении баланса карты
func (bank Bank) CheckBalance(cardNumber string) error {
	println("[Банк] Проверка баланса карты..: ", cardNumber)
	// ожидание ответа сервера
	time.Sleep(time.Second *5)
	// проходимся по всем картам проверяем номер карты
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			// пропускаем карты пока не найдем нужную таким образом в цикле проходя по всем существующим картам (не оптимально, просто для примера)
			continue
		}
		// проверка баланса нужной карты
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств!")
		}
		fmt.Println("[Банк] Остаток положительный: ", card.Balance)
	}
	// 
	return nil
}