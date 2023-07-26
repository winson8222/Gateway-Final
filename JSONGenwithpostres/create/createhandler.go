package create

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"unicode"
)

// Createhandler create modify existing hz generated handler code
func Createhandler(service HandlerInfo) {

	funcs := template.FuncMap{
		"title": func(s string) string {
			for i, v := range s {
				return string(unicode.ToUpper(v)) + s[i+1:]
			}
			return ""
		},
	}

	headertamplate := `// Code generated by hertz generator.

	package {{ .IDLName }}
	
	import (
		"context"
		"sync"
		"{{ .GatewayName }}/biz/model/{{ .IDLName }}"
		"github.com/cloudwego/kitex/client/genericclient"
		"github.com/cloudwego/hertz/pkg/app"
		"github.com/cloudwego/hertz/pkg/protocol/consts"
	)
	`

	methodtemplate := `
	func {{ title .MethodName }}(ctx context.Context, c *app.RequestContext) {
		var req {{ .IDLName }}.{{ .RequestStruct }}
		err := c.BindAndValidate(&req)
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}

		cli := myGenClientPool.Get().(genericclient.Client)
		defer myGenClientPool.Put(cli)
		st, err := c.Body()
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
	
	
		resp, err := Do{{ title .MethodName }}(ctx, cli, string(st)) // Pass the generic client and requestContext
		if err != nil {
			c.String(consts.StatusInternalServerError, "Failed to perform generic call")
			return
		}
	
		c.JSON(consts.StatusOK, resp)
	}

	`

	genclitemplate := `	
	var myGenClientPool = sync.Pool{
		New: func() interface{} {
			// Create and return a new object of the type you want to pool.
			return {{ title .ServiceName }}GenericClient()
		},
	}`

	//enter gateway folder
	desiredDir := "gateway"
	err := os.Chdir(desiredDir)
	if err != nil {
		log.Fatalf("move to folder failed with %s\n", err)
	}
	// Create the output file
	outputFile, err := os.Create("biz/handler/" + service.ServiceInfo.IDLName + "/" + service.ServiceInfo.HandlerFile + ".go")
	if err != nil {
		log.Fatalf("Error creating output file: %s\n", err)
	}
	defer outputFile.Close()

	// Create a new template for the generic client
	headerTmpl := template.Must(template.New("header").Parse(headertamplate))

	err = headerTmpl.Execute(outputFile, service.ServiceInfo)
	if err != nil {
		log.Fatalf("Error executing generic client template: %s\n", err)
	}

	// Create a new template for the method
	methodTmpl := template.Must(template.New("method").Funcs(funcs).Parse(methodtemplate))

	// Generate code for each method
	for _, method := range service.Handlers {
		// Execute the method template with the current method and write the output to the file
		err = methodTmpl.Execute(outputFile, method)
		if err != nil {
			log.Fatalf("Error executing method template: %s\n", err)
		}
	}

	gencliTmpl := template.Must(template.New("gencli").Funcs(funcs).Parse(genclitemplate))

	gencliTmpl.Execute(outputFile, service.Handlers[0])

	fmt.Print("Handler code for " + service.Handlers[len(service.Handlers)-1].ServiceName + " Generated successfully.\n")

}

// Handler struct contains information needed for a handling an api call for a specific method
type Handler struct {
	MethodName    string
	ServiceName   string
	IDLName       string
	RequestStruct string
}

// HandlerServiceInfor struct contains service related info needed for creating handlers
type HandlerServiceInfo struct {
	IDLName     string
	GatewayName string
	HandlerFile string
}

// HandlerInfor contains HandlerServiceInfo and Handler information for all methods
type HandlerInfo struct {
	ServiceInfo HandlerServiceInfo
	Handlers    []Handler
}
