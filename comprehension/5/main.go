// Что выведет программа? Объяснить вывод программы.
package main

import "fmt"

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()    // err == nil, *customError
	if err != nil { // nil без типа != nil с типом
		fmt.Println("error")
		return
	}
	fmt.Println("ok")
}
