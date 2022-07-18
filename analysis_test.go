package main

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	t.Log("wd:", wd)

	testdata := filepath.Join(wd, "testdata")
	analysistest.Run(t, testdata, Analyzer, "p")
}
