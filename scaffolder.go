package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the data file
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) < 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		folderName := parts[0]
		fileName := parts[1]

		// Create the folder
		err := os.MkdirAll(folderName, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			continue
		}

		// Create the file
		filePath := folderName + "/" + fileName
		newFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer func(newFile *os.File) {
			err := newFile.Close()
			if err != nil {
				fmt.Println("Error creating file:", err)
			}
		}(newFile)

		// Write placeholder content to the file
		_, err = newFile.WriteString(fmt.Sprintf("Content for %s", fileName))
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
