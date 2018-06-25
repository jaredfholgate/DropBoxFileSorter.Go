package main

import (
	"log"
	"os"
	"time"
)

func main() {
	sourcePath := os.Args[1]
	targetPath := os.Args[2]

	for _, f := range getFiles(sourcePath) {
		fileName := f.Name()

		if !shouldCopyFile(f.ModTime()) {
			log.Println("Skipping file as it is less than an hour old: " + fileName)
		} else {
			targetFolderName := parseFolderName(fileName)
			if targetFolderName == "" {
				log.Println("Unable to parse file name into folder name: " + fileName)
			} else {
				fullSourceFileName := sourcePath + "\\" + fileName
				fullTargetFolderName := targetPath + "\\" + targetFolderName
				fullTargetFileName := fullTargetFolderName + "\\" + fileName
				log.Println("Source File: " + fullSourceFileName)
				log.Println("Target Folder: " + fullTargetFolderName)
				log.Println("Target File: " + fullTargetFileName)

				log.Println("Creating folder is if does not extist: " + fullTargetFolderName)
				checkAndCreateFolder(fullTargetFolderName)

				log.Println("Copying file: " + fullTargetFileName)
				if copyFile(fullSourceFileName, fullTargetFileName) {
					log.Println("Deleting original file: " + fullTargetFileName)
					os.Remove(fullSourceFileName)
				} else {
					log.Println("Unalbe to copy file: " + fullTargetFileName)
				}
			}
		}
	}
}

func shouldCopyFile(fileDate time.Time) bool {
	hourAgo := time.Now().Add(time.Duration(-1) * time.Minute)
	if hourAgo.Sub(fileDate) > 0 {
		return true
	}
	return false
}
