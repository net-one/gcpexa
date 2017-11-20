package main

import (
	"fmt"

	micro "github.com/micro/go-micro"
	proto "github.com/gcpexa/services/instance/proto"
	"golang.org/x/net/context"
)


func main() {

	service := micro.NewService(micro.Name("instance.client"))

	instance := proto.NewInstanceClient("instance", service.Client())

	res, err := instance.CreateInstance(context.TODO(), &proto.CreateInstanceRequest{InstanceName: "",CompanyId:1})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.InstanceName)
}