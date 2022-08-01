// поведение для сервиса сохранения данных
package pkg

import "fmt"

type DataService struct {
	Next Service
}

func (save *DataService) Execute(data *Data) {
	// сначала должны проверить состояние данных
	if !data.UpdateSource { // если данные были полученны
		fmt.Printf("Данные не были обновлены \n")
		// выходим
		return
	}
	// сохраняем данные
	fmt.Printf("Данные сохранены \n")
}

func (save *DataService) SetNext(svc Service) {
	// задаем следующее звено
	save.Next = svc
}
