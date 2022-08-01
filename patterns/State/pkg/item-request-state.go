package pkg

import "fmt"

type ItemRequstedState struct {
	vendingMachine *VendingMachine
}

// когда пользователь запрашивает товар мы должны вернуть состояние о том что товар был запрошен
func (i *ItemRequstedState) RequestItem() error {
	return fmt.Errorf("Предмет уже был запрошен")
}

// добавление товара невозможно т.к товар уже был запрошен
func (i *ItemRequstedState) AddItem(count int) error {
	return fmt.Errorf("Предмет в процессе обработки")
}

// после запроса пользователь должен внести деньги
func (i *ItemRequstedState) InsertMoney(money int) error {
	if money < i.vendingMachine.ItemPrice {
		fmt.Errorf("Внесенных денег недостаточно. Пожалуйста внесите ещё %d", i.vendingMachine.ItemPrice)
	}
	fmt.Println("Внесенных денег достаточно")
	// машина получает деньги
	i.vendingMachine.SetState(i.vendingMachine.HasMoney)
	return nil
}

// не можем сделать выдачу товара потому-что не внесли деньги
func (i *ItemRequstedState) DispenseItem() error {
	return fmt.Errorf("Прежде чем получить товар внесите деньги")
}
