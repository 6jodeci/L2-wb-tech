package main

import (
	"fmt"
	"sync"
	"time"
)

// or преобразует список каналов в один канал, запуская goroutine для каждого входящего канала,
// которая копирует значения в единственный исходящий канал
func or(cs ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup                  // вэйтгрупа
	out := make(chan interface{})          // объедененный канал
	output := func(c <-chan interface{}) { // копирование в общий канал
		// Запускаем output goroutine
		// для каждого входного канала в cs.
		// output копирует значения из c в out
		// до тех пор пока c не закрыт, затем вызывает wg.Done.
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(1) // увеличиваем счетчик на 1
	for _, c := range cs {
		go output(c)
	}

	// Запускаем goroutine чтобы закрыть out
	// когда все output goroutine заверешены.
	// Это должно начнаться после вызова wg.Add.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(1*time.Second),
		sig(3*time.Second),
	)

	fmt.Printf("Done after: %v\n", time.Since(start))

}
