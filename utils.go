package main

import (
	"bufio"
	"log"
	"os"
)

func readFileLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	var fileInput []string
	for scanner.Scan() {
		line := scanner.Text()
		fileInput = append(fileInput, line)
	}
	return fileInput
}
