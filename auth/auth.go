package auth
import(
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	 "google.golang.org/api/compute/v1"	
)

func GetComputeService(jsonKey []byte,scope ...string)(*compute.Service,error){
	ctx := context.Background()
	
	conf, err := google.JWTConfigFromJSON(jsonKey,scope...)
	if err != nil {
		return nil, err
	}

	client := conf.Client(ctx);

	computeService,err:=compute.New(client)
	
		if err != nil {
			return nil, err
		}

	return computeService,nil
}

