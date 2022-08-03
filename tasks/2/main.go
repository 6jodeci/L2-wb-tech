// Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны
// Реализовать поддержку escape-последовательностей
// В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidStringDigit = errors.New("Строка не должна начинаться с числа!")
var ErrInvalidTwoDigit = errors.New("Два числа не должны идти друг за другом!")
var ErrInvalidSlash = errors.New("Слеш не должен быть последним символом в строке!")

func ChangeString(s string) (string, error) {
	// Перевод строки в слайс рун
	runes := []rune(s)
	// Определяет количество символов
	runeLen := len(runes)
	// Длина строки больше нуля и начинается с цифры, выводим ошибку
	if runeLen > 0 && unicode.IsDigit(runes[0]) {
		return "", ErrInvalidStringDigit
	}
	// Билдер для эфективной конкатенации и создания строк
	buildString := &strings.Builder{}
	var curSymbol rune
	// Проход по каждому символу
	for i := 0; i < runeLen; i++ {
		actualRune := runes[i]
		// Символ это число
		if unicode.IsDigit(actualRune) {
			// Следующий символ это число -- ошибка
			if i+1 < runeLen && unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidTwoDigit
			}
			// Перевод руну в инт после чего в строку
			count, err := strconv.Atoi(string(actualRune))
			if err != nil {
				return "Ошибка выполнения: ", err
			}
			// Увеличение количества символов в строке
			NewRunes(buildString, curSymbol, count)
			continue
		}
		// Получен слеш
		if actualRune == '\\' {
			i++
			// Слеш - последний символ строки
			if i == runeLen {
				return "", ErrInvalidSlash
			}
			// Добавление в руну элемента, идущего за слешем
			actualRune = runes[i]
		}
		// После буквы стоит число то сохраняем в руну
		if i+1 < runeLen && unicode.IsDigit(runes[i+1]) {
			curSymbol = actualRune
			continue
		}
		// Пишет в строку букву, после которой нет цифры или символ, который идет за слешем
		buildString.WriteRune(actualRune)

	}
	return buildString.String(), nil

}

func NewRunes(runes *strings.Builder, r rune, n int) {
	// Записывает в руну n раз
	for i := 0; i < n; i++ {
		runes.WriteRune(r)
	}
}

func main() {
	var str string
	fmt.Print("Введите строку для изменения: ")
	// Ввод строки для изменения с клавиатуры
	fmt.Scan(&str)
	// Вызов переменной на изменение строки
	str, err := ChangeString(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
