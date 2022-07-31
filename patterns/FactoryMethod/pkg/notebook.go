package pkg

import "fmt"

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

// базовый конструктор сервиса
func NewNoteBook() Computer {
	return Notebook{
		Type:    NoteBookType,
		Core:    4,
		Memory:  8,
		Monitor: true,
	}
}

// реализация сервера
func (pc Notebook) GetType() string {
	return pc.Type
}

// реализация сервера
func (pc Notebook) PrintDetails() {
	fmt.Printf("Type:||%s|| == Core:[%d] Mem:[%d], Monitor[%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
