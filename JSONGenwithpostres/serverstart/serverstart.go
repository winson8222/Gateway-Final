package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	err := os.Chdir("../gateway")
	if err != nil {
		log.Fatal("cannot move into gateway folder")
	}

	for _, arg := range os.Args[1:] {
		fmt.Printf("Starting gateway server at: %s\n", arg)

		cmd := exec.Command("./gateway", arg)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatal(err)
		}

		// Start separate goroutines to read stdout and stderr
		go readOutput(stdout)
		go readOutput(stderr)

		// Start the command and do not wait for it to finish
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("All server processes started. Exiting...")
}

// readOutput reads data from the given reader and writes it to the console
func readOutput(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
