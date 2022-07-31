package pkg

import "fmt"

type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

// базовый конструктор сервиса
func NewPersonalComputer() Computer {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

// реализация сервера
func (pc PersonalComputer) GetType() string {
	return pc.Type
}

// реализация сервера
func (pc PersonalComputer) PrintDetails() {
	fmt.Printf("Type:||%s|| == Core:[%d] Mem:[%d], Monitor[%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
