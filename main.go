package main

import(
	"google.golang.org/api/compute/v1"
	"fmt"
	"io/ioutil"
	"./auth"
)
func main() {
	jsonKey, err := ioutil.ReadFile("kuber-4e4ff93dd8f5.json");
	if err != nil {
		//return nil, err

		fmt.Println(err)
	}
	client, err := auth.GetClient(jsonKey,"https://www.googleapis.com/auth/compute");
	if err != nil {
		//return nil, err

		fmt.Println(err)
	}




	computeService,err:=compute.New(client)

	if err != nil {
		//return nil, err
		fmt.Println(err)
	}

	instanceGet:=computeService.Instances.Get("kuber-180507","us-central1-c","instance-1")

	instance,err:=instanceGet.Do()

	if err != nil {
		//return nil, err
		fmt.Println(err)
	}

	fmt.Println(instance.Kind)


}