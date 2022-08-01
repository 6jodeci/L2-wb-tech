package pkg

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

// во время того как пользователь вносит деньги он не может запросить товар т.к находится в стадии оплаты
func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Выдача предмета в процессе")
}

// добавление товара невозможно т.к происходит процесс оплаты
func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("Выдача предмета в процессе")
}

// также нельзя выполнять дополнительное внесение денег во время первого внесения
func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Выдача предмета в процессе")
}

// можно только выдать товар когда деньги были внесены
func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Выдача товара")
	i.vendingMachine.ItemCount = i.vendingMachine.ItemCount - 1
	if i.vendingMachine.ItemCount == 0 {
		// если предмета нет в наличии, то ставим соотвутсвующий статус
		i.vendingMachine.SetState(i.vendingMachine.NoItem)
	} else {
		// если предмет есть в наличии, то ставим соотвутсвующий статус
		i.vendingMachine.SetState(i.vendingMachine.HasItem)
	}
	return nil
}
