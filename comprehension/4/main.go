// Что выведет программа? Объяснить вывод программы.
package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // запись в канал
		}
		// close(ch) || нужно закрыть канал!
	}()
	for n := range ch { // чтение из открытого канала
		fmt.Println(n)
	}
}
