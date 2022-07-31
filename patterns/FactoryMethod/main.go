package main

import (
	"factory-method/pkg"
)

var types = []string{pkg.PersonalComputerType, pkg.NoteBookType, pkg.ServerType, "non-existent type"}

func main() {
	// клиентский вызов
	for _, typeName := range types {
		// создаем компьютер в зависимости от типа который передал пользователь
		computer := pkg.New(typeName)
		// проверка инициализации
		if computer == nil {
			// пропускам и выводится сообщение
			continue
		}
		// если объект найден выводим информацию о том какой объект был проинициализирован
		computer.PrintDetails()
	}
}
