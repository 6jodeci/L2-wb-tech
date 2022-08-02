package pkg

// интерфейс места, которое должен посетить посетитель
type Place interface {
	Accept(v Visitor) string
}
