/**
 *  db handler 管理模块
 */
package grpc_go_db_handler

import "google.golang.org/grpc"

func initPool(){
	for i:=0;i<poolSize;i++{
		go func() {
			conn, err := grpc.Dial(ins.address, grpc.WithInsecure())
			if err!=nil{
				pool<-nil
			}else{
				pool<-conn
			}
		}()

	}
}

func getClient() *grpc.ClientConn{
	return <-pool
}

func releaseClient(client *grpc.ClientConn) {
	pool<-client
}