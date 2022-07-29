package pkg

// структура Hp collector должны присутствуют все те же поля что и в обычном компьютере
// реализация интерфейса сборщика для бренда Hp
type HpCollector struct {
	Core int
	Brand string
	Memory int
	Monitor int
	GraphicCard int
}

// Реализация интерфейсов 
func (collector *HpCollector) SetCore() {
	collector.Core = 4
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 2
}

func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *HpCollector) GetComputer() Computer{
	return Computer{
		Core: collector.Core,
		Brand: collector.Brand,
		Memory: collector.Memory,
		Monitor: collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}