package pkg

// Реализует интерфейс Place
type Pizzeria struct{}

// Реализует посещение пиццерии
func (v *People) VisitPizzeria(p *Pizzeria) string {
	return p.BuyPizza()
}

// Прием посетителя
func (p *Pizzeria) Accept(v Visitor) string {
	return v.VisitPizzeria(p)
}

// Покупка пиццы.
func (p *Pizzeria) BuyPizza() string {
	return "Buy pizza..."
}
