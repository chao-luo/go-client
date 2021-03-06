package googleClient

import (
	"go-http-client"
)


type GoogleClient struct {
	*client.CClient
}

func(g *GoogleClient) newClient(factory *googleClientFactory)(*GoogleClient){
	parentClient := factory.clientFactory.NewClient()
	return &GoogleClient{parentClient}
}


func(g *GoogleClient) GetGoogle()(*client.CResponse, error) {
	request := g.Get()
	request.Header("User-Agent", "my-client")
	request.AddParam("key", "value")
	return request.Execute()
}