# grpc-go-db-handler


### Before you begin

* **Go version** (gRPC requires Go 1.6 or higher)
	*	`$ go version`

### Install

`go get github.com/yue-best-practices/grpc-go-db-handler`

### Environment

| 名称 | 释义 | 示例 | 是否必传 | 默认值
| --- | --- | --- | --- | ---
| `DB_SERVICE_ADDRESS` | DbService地址(ip:port) | `127.0.0.1:8000` |是 |
| `DB_SERVICE_SOURCE` | DbService项目源(Server方配置) | `testSource` | 是 |
| `POOL_SIZE` | `grpc`连接池数量 | 5 | 否 | 5
| `EXPIRE_TIME` | 每次请求超时时间(单位:`ms`) | 1000 | 否 | 1000
 
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

