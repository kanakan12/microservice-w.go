package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/building-microservices-with-go/chapter1/rpc/contract"
)

const port = 1234

func CreateClient() *rpc.Client {
	// 클라이언트 생성
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse

	// Call 메서드를 사용해 서버의 이름 붙여진 함수를 호출
	// Call 메서드는 서버에서 HelloworldResponse에 대한 접근에 에러가 없다고 생각할 수 있을 만한 응답을 작성해서 Call 메서드로 회신 할 때까지 대기하는 블로킹 함수
	// 요청을 처리하다가 에러가 발생하면 이 에러가 Call 메서드의 리턴값이 되므로 이를 적당히 처리할 수 있음
	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
