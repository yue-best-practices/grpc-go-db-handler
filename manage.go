/**
 *  db handler 管理模块
 */
package grpc_go_db_handler

import "google.golang.org/grpc"

func initPool() error{
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
	return <-pool
}

func releaseClient(client *grpc.ClientConn) {
	pool<-client
}