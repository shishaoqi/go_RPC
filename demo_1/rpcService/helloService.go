package rpcService

type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello function reply:" + request
	return nil
}

func (p *HelloService) SelectUser(request string, reply *string) error {
	*reply = "select a user info"
	return nil
}
