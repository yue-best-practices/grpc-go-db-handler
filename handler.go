package grpc_go_db_handler

import (
	"google.golang.org/grpc"
	"log"
	pb "github.com/yue-best-practices/grpc-go-db-handler/pb"
	"golang.org/x/net/context"
	"time"
	"encoding/json"
	"sync"
)

type dbHandler struct {
	address string
	dataSource string
}


var (
	pool=make(chan *grpc.ClientConn,poolSize) //grpc 连接池
	ins *dbHandler
	once sync.Once
)

func NewDbHandler() (*dbHandler,error){
	var e error
	once.Do(func() {
		e=checkEnv()
		if e!=nil{
			return
		}
		ins=&dbHandler{address:address,dataSource:source}
		e=initPool()
	})
	return ins,e
}



func (db *dbHandler) Get(table string,paras interface{},arg ...string) (interface{},error){

	//conn, err := grpc.Dial(db.address, grpc.WithInsecure())
	//
	//if err!=nil{
	//	log.Printf("Get Conn Error:%v",err)
	//	return nil,err
	//}
	//defer conn.Close()
	conn:=getClient()

	defer releaseClient(conn)

	c := pb.NewDbServiceClient(conn)


	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	var dataSource string

	if arg!=nil{
		dataSource=arg[0]
	}else{
		dataSource=db.dataSource
	}

	t,_:=json.Marshal(table)
	p,_:=json.Marshal(paras)
	d,_:=json.Marshal(dataSource)

	r,err:=c.Get(ctx,&pb.GetRequest{Table:string(t),Paras:string(p),DataSource:string(d)})

	if err != nil {
		log.Printf("Get Result Error:%v",err)
		return nil,err
	}

	var result interface{}
	err = json.Unmarshal([]byte(r.Result),&result)

	if err!=nil{
		log.Printf("Get Json Error:%v",err)
		return nil,err
	}

	return result,nil
}




//func (db *dbHandler) Get(table string,paras interface{},arg ...string) (interface{},error){
//
//	conn, err := grpc.Dial(db.address, grpc.WithInsecure())
//
//	if err!=nil{
//		log.Printf("Get Conn Error:%v",err)
//		return nil,err
//	}
//	defer conn.Close()
//
//	c := pb.NewDbServiceClient(conn)
//
//
//	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
//	defer cancel()
//
//	var dataSource string
//
//	if arg!=nil{
//		dataSource=arg[0]
//	}else{
//		dataSource=db.dataSource
//	}
//
//	t,_:=json.Marshal(table)
//	p,_:=json.Marshal(paras)
//	d,_:=json.Marshal(dataSource)
//
//	r,err:=c.Get(ctx,&pb.GetRequest{Table:string(t),Paras:string(p),DataSource:string(d)})
//
//	if err != nil {
//		log.Printf("Get Result Error:%v",err)
//		return nil,err
//	}
//
//	var result interface{}
//	err = json.Unmarshal([]byte(r.Result),&result)
//
//	if err!=nil{
//		log.Printf("Get Json Error:%v",err)
//		return nil,err
//	}
//
//	return result,nil
//}


func (db *dbHandler) GetOne(table string,where string,paras interface{},arg ...string) (interface{},error){
	conn, err := grpc.Dial(db.address, grpc.WithInsecure())

	if err!=nil{
		log.Printf("GetOne Conn Error:%v",err)
		return nil,err
	}
	defer conn.Close()

	c := pb.NewDbServiceClient(conn)


	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	var dataSource string

	if arg!=nil{
		dataSource=arg[0]
	}else{
		dataSource=db.dataSource
	}

	t,_:=json.Marshal(table)
	w,_:=json.Marshal(where)
	p,_:=json.Marshal(paras)
	d,_:=json.Marshal(dataSource)

	r,err:=c.GetOne(ctx,&pb.GetOneRequest{Table:string(t),Where:string(w),Paras:string(p),DataSource:string(d)})

	if err!=nil{
		log.Printf("GetOne Result Error:%v",err)
		return nil,err
	}

	var result interface{}
	err = json.Unmarshal([]byte(r.Result),&result)

	if err!=nil{
		log.Printf("GetOne Json Error:%v",err)
		return nil,err
	}

	return result,nil

}


func (db *dbHandler) List(table string,where interface{},arg ...string) (interface{},error){
	conn, err := grpc.Dial(db.address, grpc.WithInsecure())

	if err!=nil{
		log.Printf("List Conn Error:%v",err)
		return nil,err
	}
	defer conn.Close()

	c := pb.NewDbServiceClient(conn)


	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	var dataSource string

	if arg!=nil{
		dataSource=arg[0]
	}else{
		dataSource=db.dataSource
	}

	t,_:=json.Marshal(table)
	w,_:=json.Marshal(where)
	d,_:=json.Marshal(dataSource)

	r,err:=c.List(ctx,&pb.ListRequest{Table:string(t),Where:string(w),DataSource:string(d)})

	if err!=nil{
		log.Printf("List Result Error:%v",err)
		return nil,err
	}

	var result interface{}
	err = json.Unmarshal([]byte(r.Result),&result)

	if err!=nil{
		log.Printf("List Json Error:%v",err)
		return nil,err
	}

	return result,nil
}


func (db * dbHandler) Save(table string,paras interface{},arg ...string) (interface{},error) {
	conn, err := grpc.Dial(db.address, grpc.WithInsecure())

	if err!=nil{
		log.Printf("Save Conn Error:%v",err)
		return nil,err
	}
	defer conn.Close()

	c := pb.NewDbServiceClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	where:="id"

	dataSource:=db.dataSource

	if arg !=nil{
		if arg[0]!=""{
			where=arg[0]
		}

		if arg[1]!=""{
			dataSource=arg[1]
		}
	}

	t,_:=json.Marshal(table)
	p,_:=json.Marshal(paras)
	w,_:=json.Marshal(where)
	d,_:=json.Marshal(dataSource)

	r,err:=c.Save(ctx,&pb.SaveRequest{Table:string(t),Paras:string(p),Where:string(w),DataSource:string(d)})

	if err!=nil{
		log.Printf("Save Result Error:%v",err)
		return nil,err
	}

	var result interface{}
	err=json.Unmarshal([]byte(r.Result),&result)
	if err!=nil{
		log.Printf("Save Json Error:%v",err)
		return nil,err
	}

	return result,nil
}


func (db *dbHandler) Update(table string,where interface{},paras interface{},arg ...string) (interface{},error){
	conn, err := grpc.Dial(db.address, grpc.WithInsecure())

	if err!=nil{
		log.Printf("Save Conn Error:%v",err)
		return nil,err
	}
	defer conn.Close()

	c := pb.NewDbServiceClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	dataSource:=db.dataSource

	if arg!=nil{
		dataSource=arg[0]
	}

	t,_:=json.Marshal(table)
	w,_:=json.Marshal(where)
	p,_:=json.Marshal(paras)
	d,_:=json.Marshal(dataSource)

	r,err:=c.Update(ctx,&pb.UpdateRequest{Table:string(t),Where:string(w),Paras:string(p),DataSource:string(d)})

	if err!=nil{
		log.Printf("Update Result Error:%v",err)
		return nil,err
	}

	return r.Result,nil
}


func (db *dbHandler) Del(table string,id interface{},arg ...string) (interface{},error){
	conn, err := grpc.Dial(db.address, grpc.WithInsecure())

	if err!=nil{
		log.Printf("Del Conn Error:%v",err)
		return nil,err
	}
	defer conn.Close()

	c := pb.NewDbServiceClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	dataSource:=db.dataSource

	if arg!=nil{
		dataSource=arg[0]
	}

	t,_:=json.Marshal(table)
	i,_:=json.Marshal(id)
	d,_:=json.Marshal(dataSource)

	r,err:=c.Del(ctx,&pb.DelRequest{Table:string(t),Id:string(i),DataSource:string(d)})

	if err!=nil{
		log.Printf("Del Result Error:%v",err)
		return nil,err
	}

	return r.Result,err
}


func (db *dbHandler) MultiGet(table string,id interface{},arg ...string) (interface{},error){
	conn, err := grpc.Dial(db.address, grpc.WithInsecure())

	if err!=nil{
		log.Printf("MultiGet Conn Error:%v",err)
		return nil,err
	}
	defer conn.Close()

	c := pb.NewDbServiceClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	field:="id"
	dataSource:=db.dataSource

	if arg!=nil{

		if arg[0]!=""{
			field=arg[0]
		}

		if arg[1]!=""{
			dataSource=arg[1]
		}

	}

	t,_:=json.Marshal(table)
	i,_:=json.Marshal(id)


	log.Printf("===id:%s",string(i))

	f,_:=json.Marshal(field)
	d,_:=json.Marshal(dataSource)

	r,err:=c.MultiGet(ctx,&pb.MultiGetRequest{Table:string(t),Id:string(i),Field:string(f),DataSource:string(d)})

	if err!=nil{
		log.Printf("MultiGet Result Error:%v",err)
		return nil,err
	}

	var result interface{}

	err=json.Unmarshal([]byte(r.Result),&result)

	if err!=nil{
		log.Printf("MultiGet Json Error:%v",err)
		return nil,err
	}

	return result,nil
}


func (db *dbHandler) ToOne(table string,where string,paras string,result interface{},arg ...string) (interface{},error){
	conn, err := grpc.Dial(db.address, grpc.WithInsecure())

	if err!=nil{
		log.Printf("MultiGet Conn Error:%v",err)
		return nil,err
	}
	defer conn.Close()

	c := pb.NewDbServiceClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	dataSource:=db.dataSource

	if arg!=nil{
		dataSource=arg[0]
	}

	t,_:=json.Marshal(table)
	w,_:=json.Marshal(where)
	p,_:=json.Marshal(paras)
	_r,_:=json.Marshal(result)
	d,_:=json.Marshal(dataSource)

	r,err:=c.ToOne(ctx,&pb.ToOneRequest{Table:string(t),Where:string(w),Paras:string(p),Result:string(_r),DataSource:string(d)})

	if err!=nil{
		log.Printf("ToOne Result Error:%v",err)
		return nil,err
	}

	var res interface{}

	err=json.Unmarshal([]byte(r.Result),&res)

	if err!=nil{
		log.Printf("ToOne Json Error:%v",err)
		return nil,err
	}

	return res,nil
}



func (db *dbHandler) ToMany(table string,where string,paras string,result interface{}, arg ...string) (interface{},error){
	conn, err := grpc.Dial(db.address, grpc.WithInsecure())

	if err!=nil{
		log.Printf("ToMany Conn Error:%v",err)
		return nil,err
	}
	defer conn.Close()

	c := pb.NewDbServiceClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	dataSource:=db.dataSource

	if arg!=nil{
		dataSource=arg[0]
	}

	t,_:=json.Marshal(table)
	w,_:=json.Marshal(where)
	p,_:=json.Marshal(paras)
	_r,_:=json.Marshal(result)
	d,_:=json.Marshal(dataSource)

	r,err:=c.ToMany(ctx,&pb.ToManyRequest{Table:string(t),Where:string(w),Paras:string(p),Result:string(_r),DataSource:string(d)})

	if err!=nil{
		log.Printf("ToMany Result Error:%v",err)
		return nil,err
	}

	var res interface{}

	err=json.Unmarshal([]byte(r.Result),&res)

	if err!=nil{
		log.Printf("ToMany Json Error:%v",err)
		return nil,err
	}

	return res,nil
}


func (db *dbHandler) Count(table string,where interface{},arg ...string) (interface{},error){
	conn, err := grpc.Dial(db.address, grpc.WithInsecure())

	if err!=nil{
		log.Printf("Count Conn Error:%v",err)
		return nil,err
	}
	defer conn.Close()

	c := pb.NewDbServiceClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	dataSource:=db.dataSource

	if arg!=nil{
		dataSource=arg[0]
	}

	t,_:=json.Marshal(table)
	w,_:=json.Marshal(where)
	d,_:=json.Marshal(dataSource)

	r,err:=c.Count(ctx,&pb.CountRequest{Table:string(t),Where:string(w),DataSource:string(d)})

	if err!=nil{
		log.Printf("Count Result Error:%v",err)
		return nil,err
	}

	var result interface{}

	err=json.Unmarshal([]byte(r.Result),&result)

	if err!=nil{
		log.Printf("Count Json Error:%v",err)
		return nil,err
	}

	return result,nil

}

func (db *dbHandler) Sum(table string,field string, where interface{},arg ...string) (interface{},error) {
	conn, err := grpc.Dial(db.address, grpc.WithInsecure())

	if err!=nil{
		log.Printf("Sum Conn Error:%v",err)
		return nil,err
	}
	defer conn.Close()

	c := pb.NewDbServiceClient(conn)

	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	dataSource:=db.dataSource

	if arg!=nil{
		dataSource=arg[0]
	}

	t,_:=json.Marshal(table)
	w,_:=json.Marshal(where)
	f,_:=json.Marshal(field)
	d,_:=json.Marshal(dataSource)

	r,err:=c.Sum(ctx,&pb.SumRequest{Table:string(t),Where:string(w),Field:string(f),DataSource:string(d)})

	if err!=nil{
		log.Printf("Sum Result Error:%v",err)
		return nil,err
	}

	var result interface{}

	err=json.Unmarshal([]byte(r.Result),&result)

	if err!=nil{
		log.Printf("Sum Json Error:%v",err)
		return nil,err
	}

	return result,nil
}




