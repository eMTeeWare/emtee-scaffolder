package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	basePath := workingDirectory

	// Check if a file was dropped
	if len(os.Args) < 2 {
		fmt.Println("No data file specified. Please drag and drop a file onto this application.")
		return
	}

	// The first argument (os.Args[0]) is the application path, so file path starts at os.Args[1]
	dataFilePath := os.Args[1]

	// Open the data file
	file, err := os.Open(dataFilePath)
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

			createStructure(basePath, contentType, mediaType, mediaProvider, name, year)

		case "show":
			contentType := "Serien"
			if len(parts) != 7 {
				fmt.Println("Invalid line format:", line)
				continue
			}
			mediaType := parts[1]
			mediaProvider := parts[2]
			name := parts[3]
			year := parts[4]
			seasonNumber := parts[5]
			episodeCount := parts[6]

			createShowStructure(basePath, contentType, mediaType, mediaProvider, name, year, seasonNumber, episodeCount)

		default:
			fmt.Println("Invalid line format:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func createShowStructure(basePath string, contentType string, mediaType string, provider string, name string, year string, seasonNumber string, count string) {
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

	// Create show folder
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

	// Create Season folder
	err = os.Chdir(mediaFolderName)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
	seasonFolderName := "Season " + fmt.Sprintf("%02s", seasonNumber)
	err = os.MkdirAll(seasonFolderName, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	err = os.Chdir(seasonFolderName)
	if err != nil {
		fmt.Println("Error changing directory:", err)
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

	// iterate over episodes
	episodes, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := 1; i <= episodes; i++ {
		fileName := fmt.Sprintf("%s S%02sE%02d - %s.%s", mediaFolderName, seasonNumber, i, infix, fileType)
		filePath := fileName
		newFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		// Write placeholder content to the file
		_, err = newFile.WriteString(fmt.Sprintf("Content for %s", fileName))
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
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
