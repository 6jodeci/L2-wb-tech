package pkg

// Реализует интерфейс суши - бара
type SushiBar struct{}

// Реализует посещение суши - бара
func (v *People) VisitSushiBar(p *SushiBar) string {
	return p.BuySushi()
}

// Реализует прием посетителя в суши - баре
func (s *SushiBar) Accept(v Visitor) string {
	return v.VisitSushiBar(s)
}

// Реализует покупку суши
func (s *SushiBar) BuySushi() string {
	return "Buy sushi..."
}
