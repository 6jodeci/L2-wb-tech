/*Реализовать простейший telnet-клиент.
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
Требования:
    1. Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
	После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
    2. Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
    3. При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
	Если сокет закрывается со стороны сервера, программа должна также завершаться.
	При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// go run telnet.go --timeout=10s localhost 8080

const commandError = `Ошибка при вызове команды!:
	example: go-telnet --timeout=10s host port`

type Telnet struct {
	Timeout time.Duration
	Host    string
	Port    string
}

// createClient создает новый telnet client
func (t *Telnet) createClient(conn net.Conn) {
	fmt.Println("Welcome to Telnet!")
	console := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)
	for {
		fmt.Print("Сообщение на сервер > ")
		text, err := console.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading string: %v", err)
			return
		}
		text = strings.TrimSpace(text)

		fmt.Fprintf(conn, text+"\n")
		if text == "exit" {
			fmt.Fprintf(os.Stdout, "%s\n", "connection closed")
			return
		}

		message, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading from conn: %v\n", err)
			os.Exit(1)
		}
		message = strings.TrimSpace(message)
		fmt.Printf("Сообщение от сервера < %s\n", message)
	}
}

// parseArgs обрабатывает полученные на вход флаги
func (t *Telnet) parseArgs() bool {
	if len(os.Args) == 3 {
		t.Host = os.Args[1]
		t.Port = os.Args[2]
		return true
	}
	if len(os.Args) == 4 {
		arg := os.Args[1]
		substr := "--timeout="
		if strings.Contains(arg, substr) {
			timeDuration := strings.TrimPrefix(arg, substr)
			timeout, err := time.ParseDuration(timeDuration)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
			t.Timeout = timeout
		} else {
			return false
		}
		t.Host = os.Args[2]
		t.Port = os.Args[3]
		return true
	}
	return false
}

func main() {
	t := &Telnet{}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT)

	ok := t.parseArgs()
	if !ok {
		fmt.Fprintf(os.Stderr, "%s\n", commandError)
		os.Exit(1)
	}
	d := net.Dialer{Timeout: t.Timeout}
	conn, err := d.Dial("tcp", t.Host+":"+t.Port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	go t.createClient(conn)

	select {
	case <-quit:
		fmt.Fprintf(os.Stdout, "%s\n", "programm finished")
		os.Exit(0)
	}
}
