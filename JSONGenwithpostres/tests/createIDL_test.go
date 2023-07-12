package test

import (
	"create"
	"testing"
)

func TestCreateIDL1(t *testing.T) {
	testConstants1 := create.Constants{
		ServiceName:       "HelloService",
		FilepathToService: "../idl/hello.thrift",
		Methods: []create.Method{
			{MethodName: "hello"},
		},
	}

	create.CreateIDL(testConstants1)

	t.Log("IDL file updated")
}

func TestCreateIDL2(t *testing.T) {
	testConstants1 := create.Constants{
		ServiceName:       "BizService",
		FilepathToService: "../idl/bizrequests.thrift",
		Methods: []create.Method{
			{MethodName: "BizMethod1"},
			{MethodName: "BizMethod2"},
			{MethodName: "BizMethod3"},
		},
	}

	create.CreateIDL(testConstants1)

	t.Log("IDL file updated")
}
