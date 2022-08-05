// Concrete Receiver
package pkg

import "fmt"

type Tv struct {
	IsRunning bool
}

func (t *Tv) On() {
	t.IsRunning = true
	fmt.Println("Включаю телевизор")
}

func (t *Tv) Off() {
	t.IsRunning = false
	fmt.Println("Выключаю телевизор")
}