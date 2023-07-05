// Package gen contains code generating files needed to create customised Godzilla Gateway
package main

import (
	"create"
	"log"
	"os"
	"strings"
	"unicode"
)

func separateCamelCase(input string) string {
	var builder strings.Builder
	for i, char := range input {
		if i > 0 && unicode.IsUpper(char) {
			builder.WriteRune('_')
		}
		builder.WriteRune(unicode.ToLower(char))
	}
	return builder.String()
}

// Gateway struct contains necessary infromation of the gateway
type GatewayInfo struct {
	GatewayURL          string
	ETCD_URL            string
	GatewayName         string
	Load_Balancing_Type string
}

// ServiceInfo structs contains information needed for the gateway's connection to microservices
type ServiceInfo struct {
	IDLName string
}

func main() {

	info, serviceinfolist := GetIDL()

	gatewayexample := MakeServices(info, serviceinfolist)

	// install hz
	// Hzinstall()

	// hz gen
	Hzgen(info.GatewayName)

	//create the constant folder and files
	create.CreateConstant(gatewayexample)

	// //create gencli for all services
	for _, constant := range gatewayexample.Service_Constants {
		create.Creategencli(constant)
	}
	create.CreateMain()

	allhandlers := []create.HandlerInfo{}

	err := os.Chdir("../")
	if err != nil {
		log.Fatalf("move to directory folder failed with %s\n", err)
	}
	//Make Handler info for all every IDLs
	for _, idls := range serviceinfolist {
		allhandlers = append(allhandlers, MakeHandlerInfo(idls.IDLName, info.GatewayName))

	}

	//create handler for all methods
	for _, handler := range allhandlers {
		create.Createhandler(handler)
		err := os.Chdir("../")
		if err != nil {
			log.Fatalf("move to directory folder failed with %s\n", err)
		}
	}

	err = os.Chdir("gateway")
	if err != nil {
		log.Fatalf("move to directory folder failed with %s\n", err)
	}

	//change kitex v0.6.1 to v0.5.2 and netpoll v0.4.0 to v0.3.2 to fix some weird bug
	Update()

	//go mod tidy
	Tidy()

	//build exe
	Build(info.GatewayName)
}
