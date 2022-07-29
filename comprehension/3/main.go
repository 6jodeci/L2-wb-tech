// Что выведет программа? Объяснить вывод программы.
// Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil // внутри os.PathError лежит интерфейс
	return err                  // возвращает nil, *os.PathError
}

func main() {
	err := Foo()                             // объявляем переменную err и передаем в нее значение функции Foo() == nil, *os.PathError
	fmt.Println(err)                         // result == nil
	fmt.Println(err == nil)                  // nil без типа != nil с типом || false
	fmt.Println(err == (*os.PathError)(nil)) // true
}
