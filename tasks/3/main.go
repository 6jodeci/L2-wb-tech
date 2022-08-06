/*
    3. Утилита sort
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

// Cтруктура хранения флагов
type Flags struct {
	k       int
	n, r, u bool
}

// Реализация структуры для взаимодействия с ней
func InitFlags(k *int, n, r, u *bool) Flags {
	return Flags{
		k: *k,
		n: *n,
		r: *r,
		u: *u,
	}
}

func (f Flags) MainSort(c []byte) {

	// кастит массив байт в строку
	str := strings.Split(string(c), "\n")
	// если флаг = вывод уникальных значений
	if f.u {
		// создание структуры для отбора множеств
		keys := map[string]struct{}{}
		for _, r := range str {
			keys[r] = struct{}{}
		}
		// пересоздает строку
		str = make([]string, 0, len(keys))
		for k := range keys {
			str = append(str, k)
		}
	}

	data := make([][]string, len(str))
	for i := range str {
		data[i] = strings.Split(str[i], " ")
	}

	// если сортируемая колонка указана
	if f.k > -1 {
		switch true {
		case f.n && f.r:
			SortByNumDesc(&f, data)
		case f.n:
			SortByNumAsc(&f, data)
		case f.r:
			SortByLengthDesc(&f, data)
		default:
			SortByLengthAsc(&f, data)
		}
		fmt.Println(data)
	}
	// если сортируемая колонка не указана
	if f.k == -1 {
		switch true {
		case f.n && f.r:
			SortFullByNumDesc(&f, str)
		case f.n:
			SortFullByNumAsc(&f, str)
		case f.r:
			SortFullByLengthDesc(&f, str)
		default:
			SortFullByLengthAsc(&f, str)
		}
		fmt.Println(str)
	}

}

// Сортировка по возрастанию, если указана колонка по длине строки
func SortByLengthAsc(f *Flags, data [][]string) {
	sort.SliceStable(data, func(i, j int) bool {
		return len(data[j][f.k]) < len(data[i][f.k])
	})
}

// Сортировка по уменьшению, если указана колонка по длине строки
func SortByLengthDesc(f *Flags, data [][]string) {
	sort.SliceStable(data, func(i, j int) bool {
		return len(data[j][f.k]) > len(data[i][f.k])
	})
}

// Сортировка по возрастанию, если указана колонка по длине строки и числовому значению
func SortByNumAsc(f *Flags, data [][]string) {
	sort.SliceStable(data, func(i, j int) bool {
		return data[j][f.k] > data[i][f.k]
	})
}

// Сортировка по уменьшению, если указана колонка по числовому значению
func SortByNumDesc(f *Flags, data [][]string) {
	sort.SliceStable(data, func(i, j int) bool {
		return data[j][f.k] < data[i][f.k]
	})
}

// Сортировка по возрастанию, по длине строки без указания колонки
func SortFullByLengthAsc(f *Flags, data []string) {
	sort.SliceStable(data, func(i, j int) bool {
		return len(data[j]) < len(data[i])
	})
}

// Сортировка по уменьшению,  по длине строки без указания колонки
func SortFullByLengthDesc(f *Flags, data []string) {
	sort.SliceStable(data, func(i, j int) bool {
		return len(data[j]) > len(data[i])
	})
}

// Сортировка по возрастанию, по числовому значению
func SortFullByNumAsc(f *Flags, data []string) {
	sort.SliceStable(data, func(i, j int) bool {
		return data[j] > data[i]
	})
}

// Сортировка по уменьшению, по числовому значению
func SortFullByNumDesc(f *Flags, data []string) {
	sort.SliceStable(data, func(i, j int) bool {
		return data[j] < data[i]
	})
}

func main() {
	// Установка флагов
	k := flag.Int("k", -1, "Указание колонки для сортировки")
	n := flag.Bool("n", false, "Сортировать по числовому значению")
	r := flag.Bool("r", false, "Сортировать в обратном порядке")
	u := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	flags := InitFlags(k, n, r, u)

	// Путь к файлу всегда будет в конце вызываемой команды || go run flags parametr filename
	filePath := os.Args[len(os.Args)-1]
	// Читаем файл
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Сортируем файл с использованием нужного флага
	flags.MainSort(file)
}
