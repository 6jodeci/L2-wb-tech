package data

// все объеты которые мы используем имеют JSON формат и для того чтобы работать с сервисом аналитики нужно использовать XML
// JsonDocument НЕ реализует SendXmlData
type JsonDocument struct {
	
}

func (doc JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
} 

// Adapter - дополнительная структура для работы с сервисом аналитики
type JsonDocumentAdapter struct {
	// подставляем в адаптер структуру json - документа
	jsonDocument *JsonDocument
}

// JsonDocumentAdapter реализует SendXmlData
func (adapter JsonDocumentAdapter) SendXmlData() {
	// ковертируем json в xml
	adapter.jsonDocument.ConvertToXml()
	println("Отправка Xml данных в сервис аналитики")
}