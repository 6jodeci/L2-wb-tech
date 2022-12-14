## Состояние (State)

Паттерн State относится к поведенческим паттернам уровня объекта.

----------------------------------------------------------------------
Паттерн State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния и является объектно-ориентированной реализацией конечного автомата. Поведение объекта изменяется настолько, что создается впечатление, будто изменился класс объекта.

----------------------------------------------------------------------
Паттерн должен применяться:

- Когда поведение объекта зависит от его состояния.
- Поведение объекта должно изменяться во время выполнения программы.
- Состояний достаточно много и использовать для этого условные операторы, разбросанные по коду, достаточно затруднительно.

----------------------------------------------------------------------

Плюсы:
+ Избавляет от множества больших условных операторов состояний.
+ Концетрирует в одном месте код связанный с опеределенными состояниями. Упрощает код контекста.

----------------------------------------------------------------------
Минусы:
+ Может неоправданно усложнить код, если состояний мало и они редко меняются

----------------------------------------------------------------------
[!] В описании паттерна применяются общие понятия, такие как Класс, Объект, Абстрактный класс. Применимо к языку Go, это Пользовательский Тип, Значение этого Типа и Интерфейс. Также в языке Go вместо наследования используется композиция.

----------------------------------------------------------------------
Diagram:

![image](https://user-images.githubusercontent.com/65400970/181828484-8b65feef-f42b-4f7b-a0eb-86955660719b.png)

----------------------------------------------------------------------
Diagram2: 

![image](https://user-images.githubusercontent.com/65400970/181830329-cefc02ff-9c8c-411e-9b8a-0ff992dc90a6.png)