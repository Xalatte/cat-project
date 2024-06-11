package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Используйте: ./cat-project 'название файла'")
	}
	Slice := os.Args[1:]
	Args(Slice)
}

func Args(Slice []string) {
	for i := 0; i < len(Slice); i++ {
		file, err := os.Open(Slice[i])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
