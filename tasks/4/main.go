/*
Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func RemoveDuplicates(slice []string) []string {
	// создание пустого слайса
	unique := make([]string, 0)
	// мапа с типом bool для проверки на множества
	keys := make(map[string]bool)

	for _, v := range slice {
		if !keys[v] {
			keys[v] = true
			unique = append(unique, v)
		}
	}
	return unique
}

func FindAnagrams(slice []string) map[string][]string {
	for i := range slice {
		// переводим все символы к нижнему регистру
		slice[i] = strings.ToLower(slice[i])
	}
	anReapeted := RemoveDuplicates(slice)
	// создаем "временную мапу" где ключ хранит отсортированные слова
	tempMap := make(map[string][]string, 0)

	// делаем сортировку c использованием пакета Sort
	for _, v := range anReapeted {
		sortedRunes := []rune(v)
		sort.Slice(sortedRunes, func(i, j int) bool {
			return sortedRunes[i] < sortedRunes[j]
		})
		// кастим руны в строку
		sortedString := string(sortedRunes)
		// добавляем строку во "временную" мапу
		tempMap[sortedString] = append(tempMap[sortedString], v)
	}
	// создаем мапу для хранения результата
	resultMap := make(map[string][]string, 0)
	// проверяем "временную" мапу перед переносом в основную
	for _, v := range tempMap {
		// если длина элемента больше 1, то переносим в основную мапу
		if len(v) > 1 {
			// начинаем с нулевого элемента (первого добавленного)
			resultMap[v[0]] = v
			sort.Strings(v)
		}
	}
	return resultMap
}

func main() {
	var values = []string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик", "кот", "ток"}
	fmt.Println("Входные данные:", values)
	fmt.Println("Анаграмы:", FindAnagrams(values))
}
