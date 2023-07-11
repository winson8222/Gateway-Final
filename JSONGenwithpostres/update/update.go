package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	err := os.Chdir("../")
	if err != nil {
		log.Fatalf("move to directory folder failed with %s\n", err)
	}

	cmd := exec.Command("./gen.exe", os.Args[1], os.Args[2], os.Args[3])

	err = cmd.Run()
	if err != nil {
		log.Fatalf("create new server files failed with %s\n", err)
	}

	fmt.Print("New gateway files created\n")

	urls := []string{"8888", "8889", "8890"} //old ports
	for i, url := range urls {
		stopcmd := exec.Command("powershell", "-ExecutionPolicy", "Unrestricted", "-File", "shutdown.ps1", url)
		index := fmt.Sprint(i + 1)

		stopcmd.Stdout = os.Stdout
		stopcmd.Stderr = os.Stderr

		err := stopcmd.Run()
		if err != nil {
			log.Fatalf("server %s shutdown failed with %s\n", index, err)
		}
		fmt.Printf("server %s stopped\n", index)

		startcmd := exec.Command("powershell", "-ExecutionPolicy", "Unrestricted", "-File", "serverstart.ps1", urls[i])

		startcmd.Stdout = os.Stdout
		startcmd.Stderr = os.Stderr

		err = startcmd.Run()
		if err != nil {
			log.Fatalf("server %s start failed with %s\n", index, err)
		}
		fmt.Printf("server %s restarted\n", index)
	}

	DeleteExe()

	err = os.Chdir("update")
	if err != nil {
		log.Fatalf("move to directory folder failed with %s\n", err)
	}

}

func DeleteExe() {

	_, err := os.Stat("gateway/gateway.exe~")

	if os.IsNotExist(err) {
		fmt.Print("File does not exist.")
	} else {
		err := os.Remove("gateway/gateway.exe~")

		if err != nil {
			// If there was an error, print it out
			log.Fatal(err)
		}
		fmt.Print("temp gateway deleted")
	}

}
