package main

import(
	"golang.org/x/net/context"
	_"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"fmt"
	_"google.golang.org/api/googleapi/transport"
	"io/ioutil"
	_"google.golang.org/api/option"
)
func main() {
	ctx := context.Background()
	jsonKey, err := ioutil.ReadFile("kuber-4e4ff93dd8f5.json")
	if err != nil {
		//return nil, err

		fmt.Println(err)
	}
	conf, err := google.JWTConfigFromJSON(
		jsonKey,
		"https://www.googleapis.com/auth/compute",
	)
	if err != nil {
		//return nil, err
		fmt.Println(err)
	}


	client := conf.Client(ctx);

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

	fmt.Printf("%d",instance.Id)


}