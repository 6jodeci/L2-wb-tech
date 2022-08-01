package pkg

import "fmt"

type RoadStrategy struct{}

// поведение для построения маршрута, если мы едем на машине
func (r *RoadStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 30                 // средняя скорость
	trafficJam := 2                // оценка пробок
	total := endPoint - startPoint // длина маршрута
	// небольшой алгоритм для нахождения времени в пути
	totalTime := total * 40 * trafficJam // время для преодоление маршрута с учетом среднего времени и коэфицента пробок
	fmt.Printf("Road A: [%d] to B: [%d]|| Avg speed: [%d]|| TrafficJam: [%d]|| Total: [%d]|| TotalTime [%d] min\n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)

}
