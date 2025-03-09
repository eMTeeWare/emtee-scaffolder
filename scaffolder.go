package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get the absolute path of the folder where the application is running
	workingDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error retrieving executable path:", err)
		return
	}

	fmt.Println("Application is running in folder:", workingDirectory)

	basepath := workingDirectory

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
		parts := strings.Split(line, ";")

		contentType := parts[0]

		switch contentType {
		case "movie":
			contentType = "Filme"
			if len(parts) != 5 {
				fmt.Println("Invalid line format:", line)
				continue
			}
			mediaType := parts[1]
			mediaProvider := parts[2]
			name := parts[3]
			year := parts[4]

			createStructure(basepath, contentType, mediaType, mediaProvider, name, year)

		case "show":
			fmt.Println("Shows not yet supported", line)
			continue
		default:
			fmt.Println("Invalid line format:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func createStructure(basePath string, contentType string, mediaType string, provider string, name string, year string) {

	err := os.Chdir(basePath)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Create contentType folder
	err = os.MkdirAll(contentType, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	// Create media folder
	err = os.Chdir(contentType)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
	mediaFolderName := name + " (" + year + ")"
	err = os.MkdirAll(mediaFolderName, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	infix := ""
	fileType := ""
	if mediaType == "disc" && provider == "bluray" {
		infix = "BR Dummy"
		fileType = "disc"
	} else if mediaType == "disc" && provider == "dvd" {
		infix = "DVD Dummy"
		fileType = "disc"
	} else if mediaType == "disc" && provider == "uhd" {
		infix = "UHD Dummy"
		fileType = "disc"
	} else if mediaType == "stream" && provider == "apple" {
		infix = "Apple TV Dummy"
		fileType = "strm"
	} else if mediaType == "stream" && provider == "amazon" {
		infix = "Prime Video Dummy"
		fileType = "strm"
	} else {
		fmt.Println("Unsupported mediaType: " + mediaType + " or provider: " + provider)
		return
	}

	// Create the file
	fileName := mediaFolderName + " - " + infix + "." + fileType
	filePath := mediaFolderName + "/" + fileName
	newFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
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
