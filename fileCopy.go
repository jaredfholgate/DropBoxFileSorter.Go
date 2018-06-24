package main

import (
    "os"
    "log"
    "io"
)

// Copy a file
func copyFile(sourceFile string, destinationFile string) {
    // Open original file
    originalFile, err := os.Open(sourceFile)
    if err != nil {
        log.Fatal(err)
    }
    defer originalFile.Close()

    // Create new file
    newFile, err := os.Create(destinationFile)
    if err != nil {
        log.Fatal(err)
    }
    defer newFile.Close()

    // Copy the bytes to destination from source
    bytesWritten, err := io.Copy(newFile, originalFile)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Copied %d bytes.", bytesWritten)
    
    // Commit the file contents
    // Flushes memory to disk
    err = newFile.Sync()
    if err != nil {
        log.Fatal(err)
    }
}