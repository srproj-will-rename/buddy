package buddy

import (
	"log"
	"os"
)

const (
	InputChannelSize  = 256
	OutputChannelSize = 256
)

var (
	loggerMw = log.New(os.Stdout, "buddy: ", log.Lmicroseconds)
)

/*
	Middlewares
*/
// func parseRequest(req *Request) {
// 	loggerMw.Print("I am going to parse a request!")
// }

func logger(req *Request) {
	loggerMw.Printf("New Event: %v | Payload: %v", req.session, req.message)
}

type MiddlewarePipeline struct {
	In          chan *Request
	Out         chan *Request
	middlewares []func(*Request)
}

func (mw *MiddlewarePipeline) Run() {
	func() {
		for {
			select {
			case req := <-mw.In:
				for _, f := range mw.middlewares {
					f(req)
				}
				mw.Out <- req
			}
		}
	}()
}

func NewMiddlewarePipeline() *MiddlewarePipeline {
	// List of middleware functions
	mw := []func(*Request){logger}

	return &MiddlewarePipeline{
		In:          make(chan *Request, InputChannelSize),
		Out:         make(chan *Request, OutputChannelSize),
		middlewares: mw,
	}
}
