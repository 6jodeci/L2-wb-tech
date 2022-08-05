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

// https://golangdocs.com/golang-download-files
func main() {

	var fullURLFile string
	fmt.Println("Введите ссылку для скачивания")
	fmt.Print("> ")
	// Вставляем ссылку на скачку файла || для примера я взял iso-образ своего linux-дистрибутива
	fmt.Scan(&fullURLFile) // https://iso.pop-os.org/22.04/amd64/nvidia/11/pop-os_22.04_amd64_nvidia_11.iso

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
