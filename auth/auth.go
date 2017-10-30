package auth
import(
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"net/http"
)

func GetClient(jsonKey []byte,scope ...string)(*http.Client,error){
	ctx := context.Background()
	
	conf, err := google.JWTConfigFromJSON(jsonKey,scope...)
	if err != nil {
		return nil, err
	}

	client := conf.Client(ctx);

	return client,nil
}

