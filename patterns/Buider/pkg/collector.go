// Реализация интерфейса сборщика
package pkg

const(
	AsusCollectorType = "asus"
	HpCollectorType = "hp"
)

// сборщик методов создания комплектации компьютера
type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGraphicCard()
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType: 
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	}
}
