package main

import (
	"os"
)

func checkAndCreateFolder(folderName string) {
	os.MkdirAll(folderName, os.ModePerm)
}
