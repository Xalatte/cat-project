package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Используйте: ./cat-project 'название файла'")
	}
	Args()
}
func Args() {
	for i := 1; i < len(os.Args); i++ {
		fileData, err := os.ReadFile(os.Args[i])
		if err != nil {
			fmt.Println("Ошибка: ", os.Args[i], err)
		}
		fmt.Println(string(fileData))
	}
}
