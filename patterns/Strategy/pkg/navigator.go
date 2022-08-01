package pkg

type Navigator struct {
	Strategy
}

// реализация, когда можно налету задать поведение
func (nav *Navigator) SetStrategy(str Strategy) {
	// присваиваем новую стратегию
	nav.Strategy = str
}
