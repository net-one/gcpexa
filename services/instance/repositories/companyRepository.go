package repositories

import(
	"io/ioutil"
)

type CompanyRepository struct{

}

func NewCompanyRepository()*CompanyRepository{
	return &CompanyRepository{}
}

func (c *CompanyRepository) GetKey(companyId int32)([]byte, error){

	key, err := ioutil.ReadFile("../../../k.json");
	if err != nil {
		return nil, err
	}

	return key,nil
}