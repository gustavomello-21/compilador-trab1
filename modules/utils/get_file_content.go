package utils

import (
	"bufio"
	"fmt"
	"os"
)

var (
	fileLines []string
)

func GetFileContent() []string {
	file, err := os.Open("./code.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	if err = file.Close(); err != nil {
		fmt.Printf("Falha ao fechar o arquivo")
	}

	return fileLines
}
