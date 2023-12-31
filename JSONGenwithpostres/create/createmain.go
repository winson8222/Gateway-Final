package create

import (
	"io/ioutil"
	"log"
)

// Modify main.go file of the gateway, such that it takes in an argument
func CreateMain() {
	// Define the values for the constants

	// Define the template string

	templateString :=
		`// Code generated by hertz generator.

		package main
		
		import (
			"os"
		
			"github.com/cloudwego/hertz/pkg/app/server"
		)
		
		func main() {
			url := "0.0.0.0:" + os.Args[1]
			h := server.Default(server.WithHostPorts(url))
		
			register(h)
			h.Spin()
		}
		`

	// ToConstant function converts the input string to a constant format

	//move into gateway directory

	// Create the output file
	err := ioutil.WriteFile("main.go", []byte(templateString), 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("main.go file created successfully.")
}
