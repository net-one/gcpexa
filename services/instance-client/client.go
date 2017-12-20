package main

import (
	"fmt"
	proto "github.com/net-one/gcpexa/services/instance/proto"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"	
	"golang.org/x/net/context"
)

func main() {
	cmd.Init()
	
	//createInstance()
	//getInstance()
	deleteInstance()
}

func getInstance() {


	service := micro.NewService(micro.Name("instance.client"))

	instance := proto.NewInstanceClient("instance", service.Client())

	req := &proto.GetInstanceRequest{}
	req.CompanyId = 1
	req.InstanceName = "new-instance"
	req.Zone = "us-central1-c"
	req.Project = "kuber-180507"

	res, err := instance.GetInstance(context.TODO(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Status)
	fmt.Println(res.Id)
	fmt.Println(res.Name)
	fmt.Println(res.Preemptible)

}


func deleteInstance() {
	service := micro.NewService(micro.Name("instance.client"))

	instance := proto.NewInstanceClient("instance", service.Client())

	req := &proto.DeleteInstanceRequest{}
	req.CompanyId = 1
	req.InstanceName = "new-instance"
	req.Zone = "us-central1-c"
	req.Project = "kuber-180507"

	res, err := instance.DeleteInstance(context.TODO(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Status)
	fmt.Println(res.HttpErrorStatusCode)
	fmt.Println(res.HttpErrorMessage)
	fmt.Println(res.HttpStatusCode)

	if res.Errors != nil {
		for _, v := range res.Errors {
			fmt.Println(v.Code)
			fmt.Println(v.Location)
			fmt.Println(v.Message)
		}
	}

}

func createInstance() {
	service := micro.NewService(micro.Name("instance.client"))

	instance := proto.NewInstanceClient("instance", service.Client())

	req := &proto.CreateInstanceRequest{}
	req.CompanyId = 1
	req.InstanceName = "new-instance"
	req.Zone = "us-central1-c"
	req.Project = "kuber-180507"
	req.MinCpuPlatform = "Automatic"
	req.MachineType = "projects/kuber-180507/zones/us-central1-c/machineTypes/n1-standard-1"

	disk1 := &proto.CreateInstanceRequest_DISKS{}
	disk1.Type = "PERSISTENT"
	disk1.Boot = true
	disk1.Mode = "READ_WRITE"
	disk1.AutoDelete = true
	disk1.DeviceName = "new-instance"
	disk1.InitializeParams = &proto.CreateInstanceRequest_DISKS_INITIALIZEPARAMS{}
	disk1.InitializeParams.SourceImage = "projects/debian-cloud/global/images/debian-9-stretch-v20171025"
	disk1.InitializeParams.DiskType = "projects/kuber-180507/zones/us-central1-c/diskTypes/pd-standard"
	disk1.InitializeParams.DiskSizeGb = 10

	disks := make([]*proto.CreateInstanceRequest_DISKS, 1)
	disks[0] = disk1
	req.Disks = disks

	ni1 := &proto.CreateInstanceRequest_NETWORKINTERFACES{}
	ni1.Network = "projects/kuber-180507/global/networks/default"
	ni1.Subnetwork = "projects/kuber-180507/regions/us-central1/subnetworks/default"
	accessConfig1 := &proto.CreateInstanceRequest_NETWORKINTERFACES_ACCESSCONFIGS{}
	accessConfig1.Name = "External NAT"
	accessConfig1.Type = "ONE_TO_ONE_NAT"

	// aliasIpRange:=&proto.CreateInstanceRequest_NETWORKINTERFACES_ALIASIPRANGES{}
	// aliasIpRange.IpCidrRange=""
	// aliasIpRange.SubnetworkRangeName=""

	ni1.AccessConfigs = make([]*proto.CreateInstanceRequest_NETWORKINTERFACES_ACCESSCONFIGS, 1)
	ni1.AccessConfigs[0] = accessConfig1
	//ni1.AliasIpRanges=make([]*proto.CreateInstanceRequest_NETWORKINTERFACES_ALIASIPRANGES,1)
	//ni1.AliasIpRanges[0]=aliasIpRange
	req.NetworkInterfaces = make([]*proto.CreateInstanceRequest_NETWORKINTERFACES, 1)

	req.NetworkInterfaces[0] = ni1

	req.Scheduling = &proto.CreateInstanceRequest_SCHEDULING{}
	req.Scheduling.Preemptible = false
	req.Scheduling.OnHostMaintenance = "MIGRATE"
	req.Scheduling.AutomaticRestart = true

	res, err := instance.CreateInstance(context.TODO(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.InstanceName)
}
