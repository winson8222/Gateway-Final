package hz_gen

import (
	"fmt"
	"idl_gen"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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
	IDLs, err := idl_gen.GetIDLs()
	if err != nil {
		log.Fatalf("get IDL files failed with %s\n", err)
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

	ClearGateway()

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
				log.Fatalf("hz update failed with %s\n", err)
			}
			Tidy()
		}

	}
	err = os.Chdir("..")
	if err != nil {
		log.Fatalf("move to folder failed with %s\n", err)
	}

}

func ClearGateway() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to retrieve current working directory:", err)
	}

	err = filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Print("hello")
			return err
		}

		if info.IsDir() {
			// Delete directories
			if path != wd {
				err := os.RemoveAll(path)
				if err != nil {
					log.Println("Failed to delete directory:", path, "-", err)
				} else {
					log.Println("Deleted directory:", path)
				}
			}
		} else {
			// Delete files
			filename := info.Name()
			if !strings.HasPrefix(filename, "gateway") {
				err := os.Remove(path)
				if err != nil {
					log.Println("Failed to delete file:", path, "-", err)
				} else {
					log.Println("Deleted file:", path)
				}
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deletion completed.")
}
