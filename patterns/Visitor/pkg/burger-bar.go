package pkg

// Реализует интерфейс place
type BurgerBar struct{}

// Реализует посещение бургерной
func (v *People) VisitBurgerBar(p *BurgerBar) string {
	return p.BuyBurger()
}

// Прием посетителя.
func (b *BurgerBar) Accept(v Visitor) string {
	return v.VisitBurgerBar(b)
}

// Продажа бургера
func (b *BurgerBar) BuyBurger() string {
	return "Buy burger..."
}
