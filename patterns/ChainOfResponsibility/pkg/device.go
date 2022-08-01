// поведение для сервиса получения данных
package pkg

import "fmt"

type Device struct {
	Name string
	Next Service
}

func (device *Device) Execute(data *Data) {
	// сначала должны проверить состояние данных
	if data.GetSource { // если данные были полученны
		fmt.Printf("Данные из устройства [%s] уже были полченны\n", device.Name)
		// передаем следующему звену задачу на выполнение обработки поступивших данных
		device.Next.Execute(data)
		// выходим
		return
	}
	// если исходные данные не были приняты
	fmt.Printf("Получение данных из устройства [%s]\n", device.Name)
	data.GetSource = true
	// передаем следующему звену задачу на выполнение обработки поступивших данных
	device.Next.Execute(data)
}

func (device *Device) SetNext(svc Service) {
	// задаем следующее звено
	device.Next = svc
}
