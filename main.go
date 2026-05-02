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
	defer file.Close()

	fmt.Printf("Reading data from file\n")
	fmt.Printf("======================")
	buffer := make([]byte, 8)
	currentLine := ""

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if currentLine != "" {
				fmt.Printf("read: %s\n", currentLine)
				currentLine = ""
			}
			if err != io.EOF {
				break
			}
			return
		}

		currentChunk := string(buffer[:n])
		splitSections := strings.Split(currentChunk, "\n")

		for i := 0; i < len(splitSections)-1; i++ {
			fmt.Printf("read: %s%s\n", currentLine, splitSections[i])
			currentLine = ""
		}
		currentLine += splitSections[len(splitSections)-1]

	}
}
