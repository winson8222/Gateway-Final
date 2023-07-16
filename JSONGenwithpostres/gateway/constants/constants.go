package constants

	 import (
		"strings"
	 )
 
	func ToConstant(s string) string {
		return strings.ToUpper(strings.ReplaceAll(s, " ", "_"))
	}

	const (
		GATEWAY_URL              = "0.0.0.0:80"
	)

	const (
		ETCD_URL = "0.0.0.0:20000" //connects to a single etcd instance in the cluster
		LOAD_BALANCING = "ROUND_ROBIN"
	)
	const (
		FILEPATH_TO_HELLOSERVICE  = "../idl/hello.thrift"
		HELLOSERVICE_NAME         = "HelloService" //name registered with svc registry as rpcendpoint
	)
	
	const (
		FILEPATH_TO_BIZSERVICE  = "../idl/bizrequests.thrift"
		BIZSERVICE_NAME         = "BizService" //name registered with svc registry as rpcendpoint
	)
	