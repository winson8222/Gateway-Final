package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	err := os.Chdir("../")
	if err != nil {
		log.Fatal("Failed to move to main directory")
	}

	// Execute serverstop.ps1
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Unrestricted", "-File", "shutdown.ps1", "8888", "8889", "8890")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Executing shutdown.ps1...")
	err = cmd.Run()
	if err != nil {
		log.Fatal("Failed to execute shutdown.ps1:", err)
	}

	cmd = exec.Command("powershell", "-ExecutionPolicy", "Unrestricted", "-File", "nstop.ps1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Println("Executing nstop.ps1...")
	err = cmd.Run()
	if err != nil {
		log.Fatal("Failed to execute nstop.ps1:", err)
	}

	log.Println("Servers shutdown successfully.")

}
