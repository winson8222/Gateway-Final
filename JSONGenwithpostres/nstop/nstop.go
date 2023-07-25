package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {

	cmd := exec.Command("sudo", "nginx", "-s", "quit")

	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to stop NGINX server: %s", err)
	}

	fmt.Println("NGINX server stopped!")

	// Allow some time for the NGINX server to start up
	time.Sleep(2 * time.Second)
}

// //for Windows
// func main() {

// 	err := os.Chdir("../nginx")
// 	if err != nil {
// 		log.Fatal("cannot move into gateway folder")
// 	}

// 	cmd := exec.Command("./nginx", "-s", "quit")

// 	err = cmd.Start()
// 	if err != nil {
// 		log.Fatalf("Failed to stop NGINX server: %s", err)
// 	}

// 	fmt.Println("NGINX server stopped!")

// 	// Allow some time for the NGINX server to start up
// 	time.Sleep(2 * time.Second)
// }
