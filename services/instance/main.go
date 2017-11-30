package main

import (
	proto "github.com/gcpexa/services/instance/proto"
	repo "github.com/gcpexa/services/instance/repositories"
	instanceService "github.com/gcpexa/services/instance/service"
	micro "github.com/micro/go-micro"
	"log"
)

func main() {

	service := micro.NewService(
		micro.Name("instance"),
	)

	service.Init()

	companyRepo := repo.NewCompanyRepository()
	instance := instanceService.NewInstance(companyRepo)
	proto.RegisterInstanceHandler(service.Server(), instance)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
