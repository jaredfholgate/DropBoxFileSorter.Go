package main

import (
	"testing"
	"time"
)

func TestShouldCopyFile(t *testing.T) {
	shouldCopy := overFiveMinutes(time.Now().Add(time.Duration(-6) * time.Minute))
	if !shouldCopy {
		t.Errorf("Should not copy is incorrect, got: %s, want: %s.", "true", "false")
	}
}

func TestShouldNotCopyFile(t *testing.T) {

	shouldCopy := overFiveMinutes(time.Now())
	if shouldCopy {
		t.Errorf("Should copy is incorrect, got: %s, want: %s.", "true", "false")
	}
}
