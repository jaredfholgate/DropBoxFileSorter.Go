package main

import "testing"

func TestParsePhoto(t *testing.T) {
    parsedFolder := parseFolderName("2018-06-11 20.51.18.jpg")
    if parsedFolder != "Photos\\2018 06 June" {
       t.Errorf("Folder was incorrect, got: %s, want: %s.", parsedFolder, "Photos\\2018 06 June")
    }
}

func TestParseVideo(t *testing.T) {
    parsedFolder := parseFolderName("2018-06-11 20.51.18.mov")
    if parsedFolder != "Video\\2018 06 June" {
       t.Errorf("Folder was incorrect, got: %s, want: %s.", parsedFolder, "Photos\\2018 06 June")
    }
}