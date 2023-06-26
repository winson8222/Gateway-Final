package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Install hz for the local system
func Hzinstall() {
	// execute install hz
	cmd1 := exec.Command("go", "install", "github.com/cloudwego/hertz/cmd/hz@latest")
	err := cmd1.Run()
	if err != nil {
		log.Fatalf("install hertz failed with %s\n", err)
	}

	// Set the environment variable for the second command
	os.Setenv("GO111MODULE", "on")

	// execute install thriftgo
	cmd2 := exec.Command("go", "install", "github.com/cloudwego/thriftgo@latest")
	err = cmd2.Run()
	if err != nil {
		log.Fatalf("install thriftgo failed with %s\n", err)
	}
}

// Generate the hz based gateway under gateway folder
func Hzgen(name string) {
	//check idl files
	IDLs, err := GetIDLs()
	if err != nil {
		log.Fatalf("get IDL files failed with", err)
	}

	//create new folder for hz
	err = os.MkdirAll("gateway", os.ModePerm)
	if err != nil {
		files, err := ioutil.ReadDir("idl")
		if err != nil {
			log.Fatal(err)
		}
		IDLs := []string{}
		for _, file := range files {
			IDLs = append(IDLs, file.Name())
		}
	}

	//move to directory
	desiredDir := "gateway"
	err = os.Chdir(desiredDir)
	if err != nil {
		log.Fatalf("move to folder failed with %s\n", err)
	}

	number := 0

	for _, file := range IDLs {

		if number <= 0 {
			cmd1 := exec.Command("hz", "new", "-module", name, "-idl", "../idl/"+file)
			err = cmd1.Run()
			if err != nil {
				log.Fatalf("hz gen failed with %s\n", err)
			}
			Tidy()
			number += 1
		} else {
			cmd1 := exec.Command("hz", "update", "-idl", "../idl/"+file)
			err = cmd1.Run()
			if err != nil {
				log.Fatalf("hz gen failed with %s\n", err)
			}
			Tidy()
		}

	}
	err = os.Chdir("..")
	if err != nil {
		log.Fatalf("move to folder failed with %s\n", err)
	}

}
