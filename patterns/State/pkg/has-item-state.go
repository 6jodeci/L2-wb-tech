package pkg

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

// в этой стадии пользователь запрашивает товар
func (i *HasItemState) RequestItem() error {
	if i.vendingMachine.ItemCount == 0 {
		// если товара нет в наличии меняем статус
		i.vendingMachine.SetState(i.vendingMachine.NoItem)
		fmt.Errorf("Предмета нет в наличии")
	}
	fmt.Printf("Предмет запрошен\n")
	// устанавливаем статус "товар в наличии"
	i.vendingMachine.SetState(i.vendingMachine.ItemRequested)
	return nil
}

// добавление товара
func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d предметов было добавленно\n", count)
	// увеличиваем товар в машине
	i.vendingMachine.IncrementItemCount(count)
	return nil
}

// также нельзя выполнять дополнительное внесение денег до выбора товара
func (i *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Сначала выберите товар")
}

// товар нельзя выдать т.к его не выбрали не оплатили
func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("Сначала выберите товар")
}
