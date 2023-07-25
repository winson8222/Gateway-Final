package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type Process struct {
	Id   string
	Name string
}

// for windows
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Port numbers are required.")
	}

	var wg sync.WaitGroup
	processesToKill := make(map[string]Process) // Mapping of port:process
	var hasError bool = false
	// Collect the processes
	for _, portStr := range os.Args[1:] {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatalf("Invalid port number: %s", portStr)
		}

		wg.Add(1)
		go func(p int) {
			defer wg.Done()

			// Get the processes
			output, err := exec.Command("powershell", "-Command", fmt.Sprintf("netstat -ano | findstr :%d", p)).CombinedOutput()

			if err != nil {
				hasError = true
				log.Printf("Error getting process on port %d: %v. Output: %s", p, err, string(output))
				return
			}

			if len(output) > 0 {
				lines := strings.Split(strings.TrimSpace(string(output)), "\n")
				for _, line := range lines {
					parts := strings.Fields(line)
					processId := parts[len(parts)-1] // The process id is the last item on the line

					// Get the process name
					nameOutput, err := exec.Command("powershell", "-Command", fmt.Sprintf("(Get-Process -Id %s).Name", processId)).CombinedOutput()
					if err != nil {
						log.Printf("Error getting process name for PID %s on port %d: %v. Output: %s", processId, p, err, string(nameOutput))
						return
					}
					processName := strings.TrimSpace(string(nameOutput))

					key := fmt.Sprintf("%d:%s", p, processId) // Key is port:process
					processesToKill[key] = Process{Id: processId, Name: processName}
				}
			} else {
				hasError = true
				log.Printf("No process found running on port %d", p)
			}
		}(port)
	}

	wg.Wait()

	// Now kill the processes
	for key, process := range processesToKill {
		err := exec.Command("powershell", "-Command", fmt.Sprintf("Stop-Process -Id %s -Force", process.Id)).Run()
		if err != nil {
			hasError = true
			log.Printf("Error terminating process PID %s (name: %s) on port %s: %v", process.Id, process.Name, strings.Split(key, ":")[0], err)
		} else {
			log.Printf("Terminated process PID %s (name: %s) on port %s", process.Id, process.Name, strings.Split(key, ":")[0])
		}
	}

	if hasError {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

// macos version
// func main() {
// 	if len(os.Args) < 2 {
// 		log.Fatal("Port numbers are required.")
// 	}

// 	var wg sync.WaitGroup
// 	processesToKill := make(map[string]Process) // Mapping of port:process

// 	// Collect the processes
// 	for _, portStr := range os.Args[1:] {
// 		port, err := strconv.Atoi(portStr)
// 		if err != nil {
// 			log.Fatalf("Invalid port number: %s", portStr)
// 		}

// 		wg.Add(1)
// 		go func(p int) {
// 			defer wg.Done()

// 			// Get the processes
// 			output, err := exec.Command("powershell", "-Command", fmt.Sprintf("netstat -ano | findstr :%d", p)).CombinedOutput()

// 			if err != nil {
// 				log.Printf("Error getting process on port %d: %v. Output: %s", p, err, string(output))
// 				return
// 			}

// 			if len(output) > 0 {
// 				lines := strings.Split(strings.TrimSpace(string(output)), "\n")
// 				for _, line := range lines {
// 					parts := strings.Fields(line)
// 					processId := parts[len(parts)-1] // The process id is the last item on the line

// 					// Get the process name
// 					nameOutput, err := exec.Command("powershell", "-Command", fmt.Sprintf("(Get-Process -Id %s).Name", processId)).CombinedOutput()
// 					if err != nil {
// 						log.Printf("Error getting process name for PID %s on port %d: %v. Output: %s", processId, p, err, string(nameOutput))
// 						return
// 					}
// 					processName := strings.TrimSpace(string(nameOutput))

// 					key := fmt.Sprintf("%d:%s", p, processId) // Key is port:process
// 					processesToKill[key] = Process{Id: processId, Name: processName}
// 				}
// 			} else {
// 				log.Printf("No process found running on port %d", p)
// 			}
// 		}(port)
// 	}

// 	wg.Wait()

// 	// Now kill the processes
// 	for key, process := range processesToKill {
// 		err := exec.Command("powershell", "-Command", fmt.Sprintf("Stop-Process -Id %s -Force", process.Id)).Run()
// 		if err != nil {
// 			log.Printf("Error terminating process PID %s (name: %s) on port %s: %v", process.Id, process.Name, strings.Split(key, ":")[0], err)
// 		} else {
// 			log.Printf("Terminated process PID %s (name: %s) on port %s", process.Id, process.Name, strings.Split(key, ":")[0])
// 		}
// 	}
// }
