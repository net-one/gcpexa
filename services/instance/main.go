package main

import(
	"log"
	micro "github.com/micro/go-micro"
	proto "github.com/gcpexa/services/instance/proto"
	"golang.org/x/net/context"	
)

type Instance struct {}

func(i *Instance) CreateInstance(ctx context.Context,req *proto.CreateInstanceRequest,res *proto.CreateInstanceResponse) error{
	res.InstanceName="instance-1"
	return nil
}


func main(){

	service:=micro.NewService(
		micro.Name("instance"),
	)

	service.Init()

	proto.RegisterInstanceHandler(service.Server(),new (Instance))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}