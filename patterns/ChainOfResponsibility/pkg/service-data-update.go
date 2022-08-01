// поведение для сервиса обработки данных
package pkg

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (update *UpdateDataService) Execute(data *Data) {
	// сначала должны проверить состояние данных
	if data.UpdateSource { // если данные были обработаны
		fmt.Printf("Данные из устройства [%s] уже были обработаны \n", update.Name)
		// передаем следующему звену задачу на обрабрику поступивших данных
		update.Next.Execute(data)
		// выходим
		return
	}
	// если исходные данные не были приняты
	fmt.Printf("Обновление данных из устройства [%s]\n", update.Name)
	data.UpdateSource = true
	// передаем следующему звену задачу на выполнение обработки поступивших данных
	update.Next.Execute(data)
}

func (update *UpdateDataService) SetNext(svc Service) {
	// задаем следующее звено
	update.Next = svc
}
