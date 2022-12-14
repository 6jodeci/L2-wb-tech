## Строитель (Builder)

Паттерн Builder относится к порождающим паттернам уровня объекта.

Паттерн Builder определяет процесс поэтапного построения сложного продукта. После того как будет построена последняя его часть, продукт можно использовать.

----------------------------------------------------------------------
Плюсы:
+ Паттерн позволяет создавать пошагово какой-то продукт который завсисит от маленький составляющих частей
+ Позволяет использовать один и тот же код для создания различных объктов
+ Изолирует сложный код в сборке объекта и его основной бизнес логики

----------------------------------------------------------------------
Минусы:
- Усложняет программу из-за введения дополнительных объектов, структур интерфейсов и тд
- Клиент будет привязан к конкретному объекту строителя т.к в интерфейсе может не быть метод и его придется добавить. 

----------------------------------------------------------------------

Релизация паттерна "Строитель" на примере проекта "Фабрика производста компьютеров" приводился пример двух фабрик Asus и Hp.
Основная функция приходится на создание пошаговой конфигурации и в эту конфигурацию передается соответсвующий сборщик 

----------------------------------------------------------------------
Важно понимать, что сложный объект это не обязательно объект оперирующий несколькими другими объектами в смысле ООП. Например, нам нужно получить документ состоящий из заголовка, введения, содержания и заключения. Наш документ, это сложный объект. Что бы был какой-то единый порядок составления документа, мы будем использовать паттерн Builder.

----------------------------------------------------------------------
[!] В описании паттерна применяются общие понятия, такие как Класс, Объект, Абстрактный класс. Применимо к языку Go, это Пользовательский Тип, Значение этого Типа и Интерфейс. Также в языке Go заместо общепринятого наследования используется агрегирование и встраивание.

----------------------------------------------------------------------
Diagram:

![image](https://user-images.githubusercontent.com/65400970/181805044-27b44466-65d7-4160-935e-e2c1ee809c36.png)
