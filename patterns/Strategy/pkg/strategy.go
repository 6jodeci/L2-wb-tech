package pkg

type Strategy interface {
	// принимает две точки - начало и конец маршрута
	Route(startPoint, endPoint int)
}
