// Что выведет программа? Объяснить вывод программы.
package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // запись в канал, после 10 значений горутина умирает
		}
		// close(ch) || нужно закрыть канал!
	}()
	for n := range ch { // чтение из открытого канала будет ждать данных из канала
		fmt.Println(n) // deadlock после 0 1 2 ... 9
	}
}
