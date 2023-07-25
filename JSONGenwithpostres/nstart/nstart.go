package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	nginxPath := "/usr/local/nginx" // Replace with the actual path to your Nginx folder

	err := os.Chdir(nginxPath)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
	cmd := exec.Command("sudo", "nginx")

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error running Nginx:", err)
		return
	}

	fmt.Println("Nginx is running...")
	// Additional code here

	// Allow some time for the NGINX server to start up
	time.Sleep(2 * time.Second)
}
