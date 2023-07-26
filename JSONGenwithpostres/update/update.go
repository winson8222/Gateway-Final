package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	err := os.Chdir("../")
	if err != nil {
		log.Fatalf("move to directory folder failed with %s\n", err)
	}
	// Args1: URL Args2: build type name Arg2: LB
	cmd := exec.Command("./gen", os.Args[1], "update", os.Args[2])

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		log.Println("Standard Output:", stdout.String())
		log.Fatalf("create new server files failed with %s\n", err)
	}

	log.Println("Standard Output:", stdout.String())
	fmt.Print("New gateway files created\n")

	urls := []string{"8888", "8889", "8890"} //old ports
	for i, url := range urls {
		err := os.Chdir("shutdown")
		if err != nil {
			log.Fatalf("move to directory folder failed with %s\n", err)
		}
		// stopcmd := exec.Command("./shutdown.exe", url)
		// index := fmt.Sprint(i + 1)

		var stopcmd *exec.Cmd
		if runtime.GOOS == "windows" {
			stopcmd = exec.Command("./shutdown.exe", url)
		} else {
			stopcmd = exec.Command("./shutdown", url)
		}

		index := fmt.Sprint(i + 1)

		stopcmd.Stdout = os.Stdout
		stopcmd.Stderr = os.Stderr

		err = stopcmd.Run()
		if err != nil {
			log.Fatalf("server %s shutdown failed with %s\n", index, err)
		}
		fmt.Printf("server %s stopped\n", index)

		err = os.Chdir("../serverstart")
		if err != nil {
			log.Fatalf("move to directory folder failed with %s\n", err)
		}

		// startcmd := exec.Command("./serverstart.exe", url)
		startcmd := exec.Command("./serverstart", url)

		startcmd.Stdout = os.Stdout
		startcmd.Stderr = os.Stderr

		err = startcmd.Run()
		if err != nil {
			log.Fatalf("server %s start failed with %s\n", index, err)
		}
		fmt.Printf("server %s restarted\n", index)
		err = os.Chdir("..")
		if err != nil {
			log.Fatalf("move to directory folder failed with %s\n", err)
		}
	}

	DeleteExe()

	err = os.Chdir("update")
	if err != nil {
		log.Fatalf("move to directory folder failed with %s\n", err)
	}

}

// Delete the temp gateway file
func DeleteExe() {

	_, err := os.Stat("gateway/gateway~")

	if os.IsNotExist(err) {
		fmt.Print("File does not exist.\n")
	} else {
		err := os.Remove("gateway/gateway~")

		if err != nil {
			// If there was an error, print it out
			log.Fatal(err)
		}
		fmt.Print("temp gateway deleted")
	}

	_, err = os.Stat("gateway/gateway.exe~")

	if os.IsNotExist(err) {
		fmt.Print("File does not exist.\n")
	} else {
		err := os.Remove("gateway/gateway.exe~")

		if err != nil {
			// If there was an error, print it out
			log.Fatal(err)
		}
		fmt.Print("temp gateway deleted")
	}

}
