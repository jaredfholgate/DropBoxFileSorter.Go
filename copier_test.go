package main

import (
	"time"
	"testing"
)

func TestShouldCopyFile(t *testing.T) {
	oldTime, err := time.Parse("2006-01-02","2018-01-01")
	if err != nil {
		t.Error("Failed to parse date.")
	}
	shouldCopy := shouldCopyFile(oldTime)
    if !shouldCopy {
       t.Errorf("Should not copy is incorrect, got: %s, want: %s.", "true", "false")
    }
}

func TestShouldNotCopyFile(t *testing.T) {

    shouldCopy := shouldCopyFile(time.Now())
    if shouldCopy {
       t.Errorf("Should copy is incorrect, got: %s, want: %s.", "true", "false")
    }
}