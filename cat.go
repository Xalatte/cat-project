package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ReaderIsNilErr = errors.New("reader is nil")
	WriterIsNilErr = errors.New("writer is nil")
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Используйте: ./cat 'название файла'")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("os.Open: %s", err.Error())
		return
	}
	defer file.Close()

	if err := printFiles(file, os.Stdout); err != nil {
		fmt.Printf("readText: %s\n", err.Error())
	}
}

func printFiles(r io.Reader, w io.Writer) error {
	if r == nil {
		return ReaderIsNilErr
	}

	if w == nil {
		return WriterIsNilErr
	}

	_, err := io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("io.Copy: %w", err)
	}

	return nil
}
