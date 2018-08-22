/**
 *  db handler 管理模块
 */
package grpc_go_db_handler

import "google.golang.org/grpc"

func initPool() error{
	pool=make(chan *grpc.ClientConn,poolSize)
	for i:=0;i<poolSize;i++{
		conn, err := grpc.Dial(ins.address, grpc.WithInsecure())
		if err!=nil{
			return err
		}else{
			pool<-conn
		}
	}
	return nil
}

func getClient() *grpc.ClientConn{
	if len(pool)>0{
		return <-pool
	}else{
		conn, err := grpc.Dial(ins.address, grpc.WithInsecure())
		if err!=nil{
			return <-pool
		}else{
			return conn
		}
	}
}

func releaseClient(client *grpc.ClientConn) {
	if len(pool)<poolSize{
		pool<-client
	}else{
		client.Close()
	}
}