package censys

//Define a URL as a constant

const BaseURL = "https://search.censys.io/api"

//Define a Client struct, used for maintaining your API tokens across requests
type NewClient struct{
	id string
	secret string
}

func New(id string,secret string) *NewClient {
	return &NewClient{id: id,secret: secret}
} 
