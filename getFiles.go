package main

import (
	"io/ioutil"
	"log"
	"os"
)

func getFiles(folderPath string) []os.FileInfo {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	return files
}
