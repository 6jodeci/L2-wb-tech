package pkg

import "fmt"

type WalkStrategy struct{}

// поведение для построения маршрута, если мы идем пешком
func (r *WalkStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 4                  // средняя скорость
	total := endPoint - startPoint // длина маршрута
	// небольшой алгоритм для нахождения времени в пути
	totalTime := total * 60 // время для преодоление маршрута с учетом среднего времени
	fmt.Printf("Walk A: [%d] to B: [%d]|| Avg speed: [%d]|| Total: [%d]|| TotalTime [%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}
