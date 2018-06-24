package main

import "fmt"
import "os"

func main() {
	sourcePath := os.Args[1]
	targetPath := os.Args[2]

	for _, f := range getFiles(sourcePath) {
		fileName := f.Name()
		targetFolderName := parseFolderName(fileName)
		if targetFolderName == "" {
			fmt.Println("Unable to parse file name into folder name: " + fileName)
		} else {
			fullSourceFileName := sourcePath + "\\" + fileName
			fullTargetFolderName := targetPath + "\\" + targetFolderName
			fullTargetFileName := fullTargetFolderName + "\\" + fileName
			fmt.Println("Source File: " + fullSourceFileName)
			fmt.Println("Target Folder: " + fullTargetFolderName)
			fmt.Println("Target File: " + fullTargetFileName)
			
			fmt.Println("Creating folder is if does not extist: " + fullTargetFolderName)
			checkAndCreateFolder(fullTargetFolderName)
			
			fmt.Println("Copying file: " + fullTargetFileName)
			if copyFile(fullSourceFileName, fullTargetFileName) {
				fmt.Println("Deleting original file: " + fullTargetFileName)
				os.Remove(fullSourceFileName)
			} 
		}
	}
}
