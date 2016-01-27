package common

const(
	POST = "POST"
	PUT = "PUT"
	GET = "GET"
	DELETE = "DELETE"
	PATCH = "PATCH"

)

type RestClient struct {

	basecommand string
	method string
	options map[string]*string
}

var DefaultClient = new(RestClient)


func NewRestClient(basecommand string, method string, args []string, options []CliOption) (*RestClient, error){

	var err error
	self := new(RestClient)
	self.basecommand = basecommand
	self.method = method

	self.options, err = GetUserValues(basecommand, args, options)

	return self, err
}

func (*RestClient) invoke() (string, error){
	var err error
	var response string



	return response, err
}