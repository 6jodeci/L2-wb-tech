// Завод по производсту компьютеров
package pkg

// структура завода
type Factory struct {
	Collector Collector
}

// NewFactory инициализирует новый завод 
// завод должен уметь менять бренды в зависимости от ситуации поэтому передаем ему интерфейс
func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

// SetCollector меняет комплектацию компьютеров
func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

// CreateComputer последовательно выполняет шаги сборки // ГЛАВНЫЙ МЕТОД!
func (factory *Factory) CreateComputer() Computer {
	// можно добавлять создавать изменять шаги и выполнять в нужной последовательности
	factory.Collector.SetCore()
	factory.Collector.SetMemory()
	factory.Collector.SetBrand()
	factory.Collector.SetGraphicCard()
	factory.Collector.SetMonitor()
	// возвращает укомплетованную сборку на существующем заводе
	return factory.Collector.GetComputer()
}
