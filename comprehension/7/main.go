package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://stackoverflow.com/questions/12907653/what-is-the-difference-between-string-and-string-in-golang

// ...  создает переменный параметр. Он будет принимать ноль или более строковых аргументов и ссылаться на них как на фрагмент.
// массив интов, на выходе канал
func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() { // горутина
		for _, v := range vs { // проходим по массиву
			c <- v                                                        // значение v в канал c
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // рандомный слип
		}
		close(c) // закрываем канал

	}()
	return c
}

// значения a, b получаем из канала, на выходе канал
func merge(a, b <-chan int) <-chan int {
	c := make(chan int) // создание канала
	go func() {         // горутина
		for { // бесконечный цикл который никогда не остановится
			select {
			case v := <-a:
				c <- v // пишем в канал
			case v := <-b:
				c <- v // пишем в канал
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c { // проходимся по каналу и выводим содержимое, после прохода всех значений будут нули, т.к канал закрыт
		fmt.Println(v) // выведет 1 2 3 4 5 6 7 8 9 0 0 0 0 0 0 ... 0
	}
}
