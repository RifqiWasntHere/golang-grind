package simple

type SayHello interface {
	Hello(name string) string
}

type SayHelloService struct {
	SayHello
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

// provider

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func NewSayHelloService(sayHello SayHello) *SayHelloService {
	return &SayHelloService{
		SayHello: sayHello,
	}
}
