package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/building-microservices-with-go/chapter1/rpc/contract"
)

const port = 1234

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	
	// rpc 패키지에 있는 Register 함수는 지정된 인터페이스의 일부 메서드를 기본 서버에 공개하고 클라이언트가 서비스에 연결된 클라이언트가 해당 메서드를 호출할 수 있게 함
	rpc.Register(helloWorld)

	// 주어진 프로토콜을 사용해 소켓을 생성하고 IP 주소와 포트에 바인딩
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer l.Close()

	// 무한 루프
	// 모든 연결에 대해 수신 대기하고 있는 ListenAndServe와는 달리 RPC 서버는 각 연결을 개별적으로 처리
	// 처음 연결을 처리하고 나서 Accept를 호출해 후속 연결을 처리하거나 애플리케이션을 종료
	// Accept는 블로킹 메서드이므로 현재 서비스에 연결하려는 클라이언트가 없는 경우 Accept 메서드는 연결이 이루어 질 때까지 대기
	for {
		conn, _ := l.Accept()
		// ServeConn 메서드는 지정된 연결에서 DefalutServer 메서드를 실행하고 클라이언트가 완료될 때까지 대기
		go rpc.ServeConn(conn)
	}
}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

// Listen() 함수는 Listener 인터페이스를 구현하는 인스턴스를 리턴

// type Listener interface
// Accept 함수는 리스너의 다음 연결을 기다리고 있다가 그것을 리턴
// Accept() (Conn, error)

// Close 함수는 리스너를 닫음
// Close() error

// Addr은 리스너의 네트워크 주소를 리턴
// Addr() Addr


// 통신 프로토콜 측면에서 ServeConn은 Gob wire 타입을 사용
// Gob 형식은 Go 프로세스 사이의 통신을 용이하게 하기 위해 특별히 고안
// Protocol Buffer와 같은 것보다 사용하기 쉬우면서 좀 더 효율적일 수 있는 것을 만들려는 아이디어를 바탕으로 설계
// Gob를 사용할 때, 원본 데이터와 대상 데이터의 값과 타입이 정확히 일치할 필요가 없음
// ex) 구조체를 보낼 때 필드가 원본 구조체에는 있지만 대상 구조체에는 없으면 디코더는 이 필드를 무시하고 에러 없이 처리를 계속 진행
// ex) 특정 필드가 원본 데이터에는 없지만 대상에 있어도 디코더는 이 필드를 무시하고 나머지 메시지를 성공적으로 처리