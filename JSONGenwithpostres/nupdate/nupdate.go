package nupdate

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

func NReload() {

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("nginx", "-s", "reload")
	}

	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		cmd = exec.Command("sudo", "nginx", "-s", "reload")
	}

	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to reload NGINX server: %s", err)
	}

	fmt.Println("NGINX server reloaded!")

	// Allow some time for the NGINX server to start up
	time.Sleep(2 * time.Second)
}
