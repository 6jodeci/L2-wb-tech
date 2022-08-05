package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*Функция megre объединияет два канала в один и в main'e печатаются значения из этого
общего канала. Однако, поскольку каналы a и b в какой то момент закрываются, а селект
внутри merge не проверяет закрытость каналов, megre начинает считывать данные из
закрытых каналов, а закрытые каналы возвращают дефолтное значение типа канала. В нашем
случае каналы типа int и дефолтное значение - 0.*/
func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() { // горутина
		for _, v := range vs { 
			c <- v                                                        // значение v в канал c
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // рандомный слип
		}
		close(c) 

	}()
	return c
}

// значения a, b получаем из канала, на выходе канал
func merge(a, b <-chan int) <-chan int {
	c := make(chan int) 
	go func() {         
		for { 
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
		fmt.Println(v) // 1, 2, 3, 4, 5, 6, 7, 8, 0, 0, 0, 0...
	}
}
