package main

import (
	"os"
	"testing"
)

func Test_GetParentFolder(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for dir != "" {
		t.Log(dir)

		dir = GetParentFolder(dir)
	}
}

func Test_FindFile(t *testing.T) {
	t.Log(FindFile("test.ini"))
	t.Log(FindFile("1.txt"))
	t.Log(FindFile("Windows"))
	t.Log(FindFile("123.log"))
}
