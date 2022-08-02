package pkg

// Реализует коллекцию мест для посещения
type City struct {
	places []Place
}

// Добавить новое место в коллекцию
func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}