package grpc_go_db_handler

import (
	"os"
	"fmt"
	"strconv"
)

var (
	address=""
	source=""
	poolSize=5
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

	pSize:=os.Getenv("POOL_SIZE")
	if pSize!=""{

		ps,err:=strconv.Atoi(pSize)

		if err!=nil{
			return fmt.Errorf("env POOL_SIZE is invalid")
		}

		poolSize=ps

	}

	return nil
}