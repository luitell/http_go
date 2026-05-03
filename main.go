package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}

	fmt.Printf("Reading data from file\n")
	fmt.Printf("======================\n")

	for i := range getLinesChannel(file) {
		fmt.Println("read:", i)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	lineChannel := make(chan string)
	go func() {
		defer f.Close()
		defer close(lineChannel)
		buffer := make([]byte, 8)
		currentLine := ""

		for {
			n, err := f.Read(buffer)
			if err != nil {
				if currentLine != "" {
					lineChannel <- currentLine
					currentLine = ""
				}
				if err == io.EOF {
					break
				}
			}

			currentChunk := string(buffer[:n])
			splitSections := strings.Split(currentChunk, "\n")

			for i := 0; i < len(splitSections)-1; i++ {
				lineChannel <- fmt.Sprintf("%s%s", currentLine, splitSections[i])
				currentLine = ""
			}
			currentLine += splitSections[len(splitSections)-1]

		}
	}()
	return lineChannel
}
