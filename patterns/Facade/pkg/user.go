// Реализация поведения когда пользователь запрашивает информацию о балансе 
package pkg

// объект покупатель у которого есть имя и карта
type User struct {
	Name string
	Card *Card
}

// GetBalance возвращает баланс, у покупателя нет поля баланс но есть карта
func (user User) GetBalance() float64 {
	return user.Card.Balance
}