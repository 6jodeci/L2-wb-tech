package pkg

// структура Asus collector должны присутствуют все те же поля что и в обычном компьбтере
// реализация интерфейса сборщика для бренда Asus
type AsusCollector struct {
	Core int
	Brand string
	Memory int
	Monitor int
	GraphicCard int
}

// Реализация интерфейсов 
func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 8
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *AsusCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *AsusCollector) GetComputer() Computer{
	return Computer{
		Core: collector.Core,
		Brand: collector.Brand,
		Memory: collector.Memory,
		Monitor: collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}