package main

import(
	"chain-pattern/pkg"
)

func main() {
	// код вызова функций
	// создаем устройство 
	device := pkg.Device{Name: "Device-1"}

	// создание данных
	data := &pkg.Data{}

	// создаем сервис обновления данных
	updateService := pkg.UpdateDataService{Name: "Update-1"}

	// создаем сервис сохранения данных
	saveService := pkg.DataService{}

	// выстраиваем цепочку что и зачем нужно вызывать
	// сначала устройство передает данные сервису обновления информации
	device.SetNext(&updateService)
	
	// сервис обновления передает данные сервису сохранения
	updateService.SetNext(&saveService)

	// обработка данных 
	device.Execute(data)
}