package pkg

import "fmt"

type VendingMachine struct {
	// состояние
	HasItem       State
	ItemRequested State
	HasMoney      State
	NoItem        State
	CurrentState  State
	// количество и цена
	ItemCount int
	ItemPrice int
}

// инициализация новой машины
func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	// первоначальная инициализация состояний
	v := &VendingMachine{
		ItemCount: itemCount,
		ItemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequstedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}

	// первоночальное состояние машины
	v.SetState(hasItemState)
	v.HasItem = hasItemState
	v.ItemRequested = itemRequestedState
	v.HasMoney = hasMoneyState
	v.NoItem = noItemState
	return v
}

// реализация поведения
func (v *VendingMachine) RequestItem() error {
	return v.CurrentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.CurrentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.CurrentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.CurrentState.DispenseItem()
}

// задает состояние
func (v *VendingMachine) SetState(s State) {
	v.CurrentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("Добавление %d предметов\n", count)
	v.ItemCount = v.ItemCount + count
}
