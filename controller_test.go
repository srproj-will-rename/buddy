package buddy

import (
	"testing"
	"fmt"
)

type TestController struct{}

func (c *TestController) Print(req *Request) {
	fmt.Printf("Hello, world! Here is my request!!\n")
	req.End()
}

func TestControllerFactory(t *testing.T) {
	NewController(TestController{})
	// handle test
}

func TestMethodInvocation(t *testing.T) {
	tc := AppController{
		Controller: TestController{},
		persist: false,
	}
	Add("test", TestController{}, "Print")
	route := Lookup("test")
	req := NewRequest(nil, nil)

	tc.Invoke(route, req)
}