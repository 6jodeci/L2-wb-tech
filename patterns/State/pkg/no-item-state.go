package pkg

import "fmt"

// состояние инкапсулирует VendingMachine
type NoItemState struct {
	vendingMachine *VendingMachine
}

// реализует интерфейс State, когда нет товара в наличии
func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("Предмета нет в наличии")
}

// добавление товара 
func (i *NoItemState) AddItem(count int) error {
	// увеличиваем количество товара
	i.vendingMachine.IncrementItemCount(count)
	// меняем состояние товара на "Есть в наличии"
	i.vendingMachine.SetState(i.vendingMachine.HasItem)
	return nil
}

// когда пользователь начинает покупать возвращаем ошибку об отсутсвии товара
func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Предмета нет в наличии")
}

// когда пользователь пытается получить товар тоже возвращаем ошибку об отстутсвии товара
func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("Предмета нет в наличии")
}