package main

import "fmt"
import "os"

func main() {
	sourcePath := os.Args[1]
	targetPath := os.Args[2]

	for _, f := range getFiles(sourcePath) {
		fileName := f.Name()
		targetFolderName := parseFolderName(fileName)
		fullSourceFileName := sourcePath + "\\" + fileName
		fullTargetFolderName := targetPath + "\\" + targetFolderName
		fullTargetFileName := fullTargetFolderName + "\\" + fileName
		fmt.Println(fullSourceFileName)
		fmt.Println(fullTargetFolderName)
		checkAndCreateFolder(fullTargetFolderName)
		fmt.Println(fullTargetFileName)

		copyFile(fullSourceFileName, fullTargetFileName);
	}
}
