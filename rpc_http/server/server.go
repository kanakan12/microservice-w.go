package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/building-microservices-with-go/chapter1/rpc_http/contract"
)

const port = 1234

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	// 해당 메서드는 HTTP를 통해 RPC를 사용하기 위해 반드시 필요
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	log.Printf("Server starting on port %v\n", port)

	http.Serve(l, nil)
}

// 전송 프로토콜 HTTP를 사용해야 하는 경우, rpc 패키지의 HandleHTTP 메서드를 호출해 이를 쉽게 처리할 수 있음

// HandleHTTP 메서드는 애플리케이션에서 두 개의 엔드 포인트를 설정
// DefaultRPCPath = "/_goRPC"
// DefaultDebugPath = "/debug/rpc"
// DefaultDebugPath에 브라우저의 경로를 지정했다고 해서 웹 브라우저에서 API와 쉽게 통신할 수 있는 것은 아님
// 메시지는 여전히 gob로 인코딩돼 있기 때문에 자바스크립트로 gob 인코더와 디코더를 작성해야 하는데 실제로 가능한지는 확실치 않음