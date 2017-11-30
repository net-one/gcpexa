package service

import (
	auth "github.com/gcpexa/auth"
	proto "github.com/gcpexa/services/instance/proto"
	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v1"
)

type instance struct {
	companyRepo companyRepository
}

func NewInstance(companyRepo companyRepository) *instance {
	return &instance{companyRepo: companyRepo}
}

func (inst *instance) CreateInstance(ctx context.Context, req *proto.CreateInstanceRequest, res *proto.CreateInstanceResponse) error {

	key, err := inst.companyRepo.GetKey(req.CompanyId)

	if err != nil {
		return err
	}

	computeClient, err := auth.GetComputeService(key, "https://www.googleapis.com/auth/compute")

	if err != nil {
		return err
	}
	 instanceInfo := compute.Instance{}
	instanceInfo.Name = req.InstanceName
	instanceInfo.MinCpuPlatform = req.MinCpuPlatform
	instanceInfo.MachineType = req.MachineType
	n := len(req.Metadata)

	if n > 0 {
		instanceInfo.Metadata.Items = make([]*compute.MetadataItems, n, 2*n)

		var i int32
		i = 0
		for key, value := range req.Metadata {
			instanceInfo.Metadata.Items[i].Key = key
			instanceInfo.Metadata.Items[i].Value = &value
			i++
		}
	}

	if(req.Tags!=nil){
	instanceInfo.Tags.Items = req.Tags.Items
	}

	m := len(req.Disks)

	if m > 0 {
		instanceInfo.Disks = make([]*compute.AttachedDisk, m, 2*m)

		for i := 0; i < m; i++ {
			instanceInfo.Disks[i]=&compute.AttachedDisk{}
			instanceInfo.Disks[i].Type = req.Disks[i].Type
			instanceInfo.Disks[i].Boot = req.Disks[i].Boot
			instanceInfo.Disks[i].Mode = req.Disks[i].Mode
			instanceInfo.Disks[i].AutoDelete = req.Disks[i].AutoDelete
			instanceInfo.Disks[i].DeviceName = req.Disks[i].DeviceName

			instanceInfo.Disks[i].InitializeParams=&compute.AttachedDiskInitializeParams{}
			instanceInfo.Disks[i].InitializeParams.SourceImage = req.Disks[i].InitializeParams.SourceImage
			instanceInfo.Disks[i].InitializeParams.DiskType = req.Disks[i].InitializeParams.DiskType
			instanceInfo.Disks[i].InitializeParams.DiskSizeGb = req.Disks[i].InitializeParams.DiskSizeGb
		}
	}

	instanceInfo.CanIpForward = req.CanIpForward

	p := len(req.NetworkInterfaces)

	if p > 0 {
		instanceInfo.NetworkInterfaces = make([]*compute.NetworkInterface, p, 2*p)

		for i := 0; i < p; i++ {
			instanceInfo.NetworkInterfaces[i]=&compute.NetworkInterface{}
			instanceInfo.NetworkInterfaces[i].Network = req.NetworkInterfaces[i].Network
			instanceInfo.NetworkInterfaces[i].Subnetwork = req.NetworkInterfaces[i].Subnetwork
			q := len(req.NetworkInterfaces[i].AccessConfigs)

			if q > 0 {
				instanceInfo.NetworkInterfaces[i].AccessConfigs=make([]*compute.AccessConfig,q)
				for j := 0; j < q; j++ {
					instanceInfo.NetworkInterfaces[i].AccessConfigs[j]=&compute.AccessConfig{}
					instanceInfo.NetworkInterfaces[i].AccessConfigs[j].Name = req.NetworkInterfaces[i].AccessConfigs[j].Name
					instanceInfo.NetworkInterfaces[i].AccessConfigs[j].Type = req.NetworkInterfaces[i].AccessConfigs[j].Type
				}
			}

			r := len(req.NetworkInterfaces[i].AliasIpRanges)

			if r > 0 {
				instanceInfo.NetworkInterfaces[i].AliasIpRanges=make([]*compute.AliasIpRange,r)
				
				for k := 0; k < r; k++ {
					instanceInfo.NetworkInterfaces[i].AliasIpRanges[k]=&compute.AliasIpRange{}
					instanceInfo.NetworkInterfaces[i].AliasIpRanges[k].IpCidrRange = req.NetworkInterfaces[i].AliasIpRanges[k].IpCidrRange
					instanceInfo.NetworkInterfaces[i].AliasIpRanges[k].SubnetworkRangeName = req.NetworkInterfaces[i].AliasIpRanges[k].SubnetworkRangeName
				}
			}

		}
	}

	instanceInfo.Description = req.Description
	instanceInfo.Labels = req.Labels
	instanceInfo.Scheduling=&compute.Scheduling{}
	instanceInfo.Scheduling.Preemptible = req.Scheduling.Preemptible
	instanceInfo.Scheduling.OnHostMaintenance = req.Scheduling.OnHostMaintenance
	instanceInfo.Scheduling.AutomaticRestart = &req.Scheduling.AutomaticRestart

	_,err=computeClient.Instances.Insert(req.Project, req.Zone, &instanceInfo).Do()

	
	if err != nil {
		return err
	}

	instanceResponse,err:=computeClient.Instances.Get(req.Project, req.Zone,instanceInfo.Name).Do()

	if err != nil {
		return err
	}
	res.InstanceName = instanceResponse.Name
	return nil
}

type companyRepository interface {
	GetKey(companyId int32) ([]byte, error)
}
