package main

import (
	"flag"
	"log"
	"os"
	"time"
)

var (
	sourcePath string
	targetPath string
)

func init() {
	flag.StringVar(&sourcePath, "sourcePath", "", "The source folder path.")
	flag.StringVar(&targetPath, "targetPath", "", "The target folder path.")
	flag.Parse()
}

func main() {
	if sourcePath == "" || targetPath == "" {
		panic("Source path and target path are required!")
	}
	moveFiles()
}

func moveFiles() {
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
	fiveMinutesAgo := time.Now().Add(time.Duration(-5) * time.Minute)
	if fiveMinutesAgo.Sub(fileDate) > 0 {
		return true
	}
	return false
}
