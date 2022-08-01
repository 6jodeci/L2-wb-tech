package pkg

type Service interface {
	Execute(*Data)
	// задача следующего исполнителя
	SetNext(Service)
}

// структура которая будет переходить от одного сервиса к другому
type Data struct {
	// признак того, выполнился ли прием данных
	GetSource bool
	// ставит отметку на тот сервис, который обработал данные
	UpdateSource bool
}
