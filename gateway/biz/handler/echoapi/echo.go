// Code generated by hertz generator.

	package echoapi
	
	import (
		"context"
		"log"
	
		"winson/biz/model/echoapi"
	
		"github.com/cloudwego/hertz/pkg/app"
		"github.com/cloudwego/hertz/pkg/protocol/consts"
	)
	
	func Echo(ctx context.Context, c *app.RequestContext) {
		var req echoapi.Request
		err := c.BindAndValidate(&req)
		if err != nil {
			c.String(consts.StatusBadRequest, err.Error())
			return
		}

		reqBody, err := c.Body()
		if err != nil {
			log.Fatal(err)
		}

	
		cli := EchoGenericClient()
		resp, err := DoEcho(ctx, cli, string(reqBody)) // Pass the generic client and requestContext
		if err != nil {
			c.String(consts.StatusInternalServerError, "Failed to perform generic call")
			return
		}
	
		c.JSON(consts.StatusOK, resp)
	}
	