// Создать программу печатающую точное время с использованием NTP -библиотеки.
// Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp.
// Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(time)
}
