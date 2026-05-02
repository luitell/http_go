package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("error reading file", err)
	}
	defer file.Close()
	buffer := make([]byte, 8)

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				break
			}
			return
		}
		fmt.Printf("read: %s\n", buffer[:n])
	}
}
