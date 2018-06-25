package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"
)

func parseFolderName(fileName string) string {
	dateString := string([]rune(fileName)[0:10])
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		log.Println("Error occurred parsing: " + dateString)
		return ""
	}

	return rootFolder(filepath.Ext(fileName)) + "\\" + fmt.Sprintf("%04d", date.Year()) + " " + fmt.Sprintf("%02d", date.Month()) + " " + date.Month().String()
}

func rootFolder(extension string) string {
	rootPath := "Photos"
	switch extension {
	case ".mov", ".mp4":
		rootPath = "Video"
	}
	return rootPath
}
