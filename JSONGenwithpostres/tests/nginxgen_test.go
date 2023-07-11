package test

import (
	"create"
	"testing"
)

func TestNignxGen(t *testing.T) {
	testConstants1 := create.Constants{
		ServiceName: "HelloService",
		Methods: []create.Method{
			{MethodName: "hello"},
		},
	}

	testConstants2 := create.Constants{
		ServiceName: "HelloService",
		Methods: []create.Method{
			{MethodName: "echo"},
		},
	}

	testServices := create.Services{
		GATEWAY_URL: "0.0.0.0:80",
		Service_Constants: []create.Constants{
			testConstants1,
			testConstants2,
		},
	}
	create.NginxConfig(testServices)

	t.Log("Testing creation of nignx.conf")
}
