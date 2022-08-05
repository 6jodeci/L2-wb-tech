// Реализовать утилиту wget с возможностью скачивать сайты целиком.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	fileName    string
	fullURLFile string
)

func main() {

	var fullURLFile string
	fmt.Println("Welcome to Wget!")
	fmt.Println("Введите ссылку для скачивания:")
	fmt.Print("> ")
	// Вставляем ссылку на скачку файла || страницы == для примера я взял iso-образ своего linux-дистрибутива
	fmt.Scan(&fullURLFile) // https://iso.pop-os.org/22.04/amd64/nvidia/11/pop-os_22.04_amd64_nvidia_11.iso
	// https://golangify.com/func для примера скачивания сайта использовал эту страницу

	// Генерирует название скаченного файла исходя из ссылки
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]

	// Создает пустой файл
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Заполнение файла контентом
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	// Выводит информацию о скаченном файле
	fmt.Printf("Downloaded a file %s with size %d", fileName, size)

}
