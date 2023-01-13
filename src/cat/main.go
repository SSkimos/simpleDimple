package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// argsWithoutProg := os.Args[1]

	file, err := openFile("abobaa.txt")

	if err != nil {
		log.Fatal(err)
	}

	// сделал хз что
	defer file.Close()

	// объявил ридер и сделал массив
	reader := bufio.NewReader(file)
	buf := make([]byte, 16)

	// цикл
	for {
		// беру строку
		n, err := reader.Read(buf)

		// проверка ошибок
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		// вывод
		fmt.Print(string(buf[0:n]))
	}
	// вывод новой строки
	fmt.Println()
}

func openFile(pathToFile string) (*os.File, error) {
	file, err := os.Open(pathToFile)
	return file, err
}
