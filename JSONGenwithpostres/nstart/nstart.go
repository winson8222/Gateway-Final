package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	opath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Chdir("../nginx/conf/")
	if err != nil {
		log.Fatal("move to folder failed")
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	npath, exist := os.LookupEnv("NGINX_PATH")

	if !exist {
		log.Fatal("Nginx not installed Or folder not set to NGINX_PATH")
		return
	}

	err = os.Chdir(npath)
	if err != nil {
		log.Fatal("move to nginx folder failed")
	}

	cmd := exec.Command("nginx", "-c", path+"/nginx.conf")

	err = cmd.Start()
	if err != nil {
		// Print the captured logs
		fmt.Println("Error running Nginx:", err)
		return
	}

	err = os.Chdir(opath)
	if err != nil {
		log.Fatal("return to original location failed")
	}

	fmt.Println("Nginx is running...")
	// Additional code here

	// Allow some time for the NGINX server to start up
	time.Sleep(2 * time.Second)
}
