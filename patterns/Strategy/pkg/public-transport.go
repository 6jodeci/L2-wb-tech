package pkg

import "fmt"

type PublicTransportStrategy struct{}

// поведение для построения маршрута, если мы едем на общественном транспорте
func (r *PublicTransportStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 40                 // средняя скорость
	total := endPoint - startPoint // длина маршрута
	// небольшой алгоритм для нахождения времени в пути
	totalTime := total * 40 // время для преодоление маршрута с учетом среднего времени
	fmt.Printf("Public transport A: [%d] to B: [%d]|| Avg speed: [%d]|| Total: [%d]|| TotalTime [%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}
