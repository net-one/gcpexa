package service

import(
	proto "github.com/gcpexa/services/instance/proto"
	"golang.org/x/net/context"	
)

type instance struct {
	companyRepo companyRepository
}

func NewInstance(companyRepo companyRepository)*instance{
return &instance{companyRepo:companyRepo}
}

func(i *instance) CreateInstance(ctx context.Context,req *proto.CreateInstanceRequest,res *proto.CreateInstanceResponse) error{

	_,err:=i.companyRepo.GetKey(req.CompanyId);

	if(err==nil){
		return nil
	}

	res.InstanceName="instance-1"
	return nil
}

type companyRepository interface {
	GetKey(companyId int32)([]byte, error)
}