package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	// http 패키지에서 HandleFunc 매서드 호출
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)

	// HTTP 서버 시작
	// ListenAndServe => 서버를 바인드할 TCP 네트워크 주소와 요청을 라우팅하는 데 사용할 핸들러 두 가지 매개 변수 사용
	// IP 주소에 서버를 바인딩
	// nil => 포이터, 인터페이스, 맵, 슬라이스, 채널, 함수의 zero value
	// zero value => 명시적인 초기값을 할당하지 않고, 변수를 만들었을 때 해당 변수가 갖게 되는 값
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world\n")
}
