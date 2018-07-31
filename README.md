# grpc-go-db-handler


### Before you begin

* **Go version** (gRPC requires Go 1.6 or higher)
	*	`$ go version` 
* **Install Protocol Buffers v3**
	* Install the protoc compiler that is used to generate gRPC service code. The simplest way to do this is to download pre-compiled binaries for your platform(protoc-<version>-<platform>.zip) from here:[https://github.com/google/protobuf/releases](https://github.com/google/protobuf/releases).
	* Unzip this file.
	* Update the environment variable `PATH` to include the path to the protoc binary file.

### Install

`go get github.com/yue-best-practices/grpc-go-db-handler`

### Example

```go
package main

import (
	dbhandler "github.com/yue-best-practices/grpc-go-db-handler"
	"log"
)

func main() {
	handler,err:=dbhandler.NewDbHandler()

	if err!=nil{
		log.Fatalf("err:%v",err)
	}

	result,err :=handler.Get("tbl_user", 1) // <tableName> <id Value>


	if err!=nil{
		log.Fatalf("err:%v",err)
	}


	log.Printf("result:%v",result)
}
```

### Run

`$ DB_SERVICE_SOURCE=<db-service-source> DB_SERVICE_ADDRESS=<db-service-address> go run main.go`

