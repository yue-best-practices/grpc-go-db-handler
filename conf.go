package grpc_go_db_handler

import (
	"os"
	"fmt"
)

var (
	address=""
	source=""
)


func checkEnv() error{

	ad:=os.Getenv("DB_SERVICE_ADDRESS")
	if ad==""{
		return fmt.Errorf(" env DB_SERVICE_ADDRESS not set")
	}
	address=ad

	sr:=os.Getenv("DB_SERVICE_SOURCE")

	if sr!=""{
		source=sr
	}

	return nil
}