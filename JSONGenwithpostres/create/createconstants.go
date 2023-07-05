package create

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// Constants struct contains information needed for building constant.go in gateway
type Constants struct {
	FilepathToService   string
	ServiceName         string
	Methods             []Method
	IDLName             string
	GatewayName         string
	Load_Balancing_Type string
}

// Services struct contains necessary information for creating constants.go files
type Services struct {
	GATEWAY_URL         string
	ETCD_URL            string
	LOAD_BALANCING_TYPE string
	Service_Constants   []Constants
}

// Method struct contains information on a method in a service
type Method struct {
	MethodName string
}

// CreateConstant create constant.go files needed for connection to all services
func CreateConstant(services Services) {
	// Define the values for the constants

	// Define the template string

	templateString :=
		`package constants

	 import (
		"strings"
	 )
 
	func ToConstant(s string) string {
		return strings.ToUpper(strings.ReplaceAll(s, " ", "_"))
	}

	const (
		GATEWAY_URL              = "{{ .GATEWAY_URL }}"
	)

	const (
		ETCD_URL = "{{ .ETCD_URL }}" //connects to a single etcd instance in the cluster
		LOAD_BALANCING = "{{ .LOAD_BALANCING_TYPE }}"
	)`

	servicetemplate :=
		`
	const (
		FILEPATH_TO_{{ .ServiceName | ToConstant }}  = "{{ .FilepathToService }}"
		{{ .ServiceName | ToConstant }}_NAME         = "{{ .ServiceName }}" //name registered with svc registry as rpcendpoint
	)
	`

	// ToConstant function converts the input string to a constant format

	// Create a new template
	tmpl := template.Must(template.New("constants").Funcs(template.FuncMap{"ToConstant": ToConstant}).Parse(templateString))

	//move into gateway directory
	desiredDir := "gateway"
	err := os.Chdir(desiredDir)
	if err != nil {
		log.Fatalf("move to folder failed with %s\n", err)
	}

	// Create the output file
	err = os.MkdirAll("constants", os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating output folder: %s\n", err)
	}
	outputFile, err := os.Create("constants/constants.go")
	if err != nil {
		log.Fatalf("Error creating output file: %s\n", err)
	}
	defer outputFile.Close()

	// Execute the template with the constants and write the output to the file
	err = tmpl.Execute(outputFile, services)
	if err != nil {
		log.Fatalf("Error executing constants template: %s\n", err)
	}

	serviceTmpl := template.Must(template.New("services_constants").Funcs(template.FuncMap{"ToConstant": ToConstant}).Parse(servicetemplate))

	// Generate code for each method
	for _, constants := range services.Service_Constants {
		// Execute the method template with the current method and write the output to the file
		err = serviceTmpl.Execute(outputFile, constants)
		if err != nil {
			log.Fatalf("Error executing service constants template: %s\n", err)
		}
	}

	log.Println("Template generation completed successfully.")
}

func ToConstant(s string) string {
	return strings.ToUpper(strings.ReplaceAll(s, " ", "_"))
}
