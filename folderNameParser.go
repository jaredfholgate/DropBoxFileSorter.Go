package main

import (
	"fmt"
	"path/filepath"
	"time"
)

func parseFolderName(fileName string) string {
	dateSlice := []rune(fileName)
	dateString := string(dateSlice[0:10])
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		fmt.Println("Error occurred parsing " + dateString)
		return ""
	}

	extension := filepath.Ext(fileName)

	var rootPath string

	switch extension {
	case ".mov", ".mp4":
		rootPath = "Video"
	default:
		rootPath = "Photos"
	}

	folderName := rootPath + "\\" + fmt.Sprintf("%04d", date.Year()) + " " + fmt.Sprintf("%02d", date.Month()) + " " + date.Month().String()
	return folderName
}
