package service

// HelloService defines the interface for the hello service
type HelloService interface {
	SayHello(name string) string
}

// helloService is the implementation of HelloService
type helloService struct{}

// NewHelloService creates a new HelloService instance
func NewHelloService() HelloService {
	return &helloService{}
}

// SayHello returns a greeting message
func (s *helloService) SayHello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
