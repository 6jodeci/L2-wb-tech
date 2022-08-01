package main

import (
	"strategy-pattern/pkg"
)

var (
	start      = 10              // точка начала
	end        = 100             // точка окончания
	strategies = []pkg.Strategy{ // выбор стратегии
		&pkg.PublicTransportStrategy{},
		&pkg.RoadStrategy{},
		&pkg.WalkStrategy{},
	}
)

func main() {
	// создание контекста навигатора
	nav := pkg.Navigator{}
	// перебор массива со стратегиями для динамического изменения поведения в навигаторе
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		// построение маршрута в зависимости от стратегии
		nav.Route(start, end)
	}
}
