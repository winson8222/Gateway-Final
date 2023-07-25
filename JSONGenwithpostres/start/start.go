package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	//start nginx and servers
	err := os.Chdir("../")
	if err != nil {
		log.Fatalf("move to directory folder failed with %s\n", err)
	}

	Start()
}

func Start() {
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Unrestricted", "-File", "nstart.ps1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Executing nstart.ps1...")
	err := cmd.Run()
	if err != nil {
		log.Fatal("Failed to execute nstart.ps1:", err)
	}

	// Execute serverstart.ps1
	cmd = exec.Command("powershell", "-ExecutionPolicy", "Unrestricted", "-File", "serverstart.ps1", "8888",
		"8889", "8890")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Executing serverstart.ps1...")
	err = cmd.Run()
	if err != nil {
		log.Fatal("Failed to execute serverstart.ps1:", err)
	}

	log.Println("Servers started successfully.")
}
