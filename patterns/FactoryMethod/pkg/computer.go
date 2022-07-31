package pkg

import "fmt"

const (
	// что будем продавать
	ServerType           = "server"
	PersonalComputerType = "personal"
	NoteBookType         = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

// factory method который инициализирует структуры и возвращает интерфейс компьютера
// общая реализация после чего передаем типы компьютера
func New(typeName string) Computer {
	switch typeName {
	// если пользователь передает информацию о несуществующем типе
	default:
		fmt.Printf("Несуществующий тип объекта!: %s\n", typeName)
		return nil
		// в зависимости от каждого типа добавляем реализацию
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NoteBookType:
		return NewNoteBook()
	}
}
