/*
Взаимодействие с ОС
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

    cd - смена директории (в качестве аргумента могут быть то-то и то)
    pwd - показать путь до текущего каталога
    echo - вывод аргумента в STDOUT
    kill - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
    ps - выводит общую информацию по запущенным процессам в формате такой-то формат
	Так же требуется поддерживать функционал fork/exec-команд
	Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
	*Шелл — это обычная консольная программа, которая будучи запущенной,
	в интерактивном сеансе выводит некое приглашение в STDOUT и ожидает ввода пользователя через STDIN.
	Дождавшись ввода, обрабатывает команду согласно своей логике и при необходимости выводит результат на экран.
	Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var ErrorNoPath = errors.New("path not found")

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to my cmd!")
	for {
		// принтлн для показа ожидания ввода
		fmt.Print("> ")
		// считывание команды с клавиатуры
		input, err := reader.ReadString('\n') // читаем до момента переноста строки
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// удаление символа переноса новой строки
	input = strings.TrimSuffix(input, "\n") // ** TODO: \r i.d.k what is preffix
	// разделение чтобы добавить флаги для команд
	args := strings.Split(input, " ")
	// проверка на наличие вложенных команд
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrorNoPath
		}
		// меняем директорию и возвращаем ошибку
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}
	// подготовка команды к выполнению
	cmd := exec.Command(args[0], args[1:]...)
	// настройка устройства ввода и вывода
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	// выполнение команды и возвращение ошибки
	return cmd.Run()
}
