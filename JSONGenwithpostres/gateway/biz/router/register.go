// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	echoapi "gateway/biz/router/echoapi"
	http "gateway/biz/router/http"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	echoapi.Register(r)

	http.Register(r)
}
