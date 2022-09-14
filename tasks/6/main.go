package main

import (
	"flag"
	"fmt"
	"os"
	"task6/util"
)

func main() {
	f := flag.String("f", "", "fields")     // выбрать поля (колонки)
	d := flag.String("d", ",", "delimiter") // использовать другой разделитель
	s := flag.Bool("s", false, "separated") // только строки с разделителем
	flag.Parse()

	// file - всегда последний аргумент команды
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("grep: ", err.Error())
		return
	}
	sets := util.InitFlags(*f, *d, *s)

	result, err := util.Cut(file, sets)

	if err != nil {
		fmt.Println("grep: ", err.Error())
	} else {
		fmt.Println(result)
	}
}
