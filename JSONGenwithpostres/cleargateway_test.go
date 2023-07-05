package main

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestMyFunction(t *testing.T) {
	ClearGateway()

	files, err := ioutil.ReadDir("gateway")
	if err != nil {
		log.Fatal("Error occur")
	}

	for _, file := range files {
		filename := file.Name()
		if !strings.HasPrefix(filename, "gateway") {
			t.Error("left over files in gateway")
		}
	}
	t.Log("Gateway files are cleared")
}
