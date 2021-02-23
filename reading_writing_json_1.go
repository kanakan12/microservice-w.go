package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	// Go에서는 소문자로 된 프로퍼티는 내보낼 수 없음
	// Mashal 함수는 이러한 프로퍼터를 무시하여 출력에 포함시키지 않음
	Message string
}

func main() {
	// := => 변수 선언 및 대입
	port := 8080

	// 요청된 Request Path에 어떤 Request 핸들러를 사용할 지를 저정하는 라우팅 역할
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	// ListenAndServe => 지정된 포트에 웹 서버를 열고 클라이언트 Request를 받아들여 Go 루틴에 작업을 할당
	// Fatal => 에러 로그를 출력하고 프로그램을 완전히 종료
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello World"}
	// 논리적 구조를 로우 바이트로 변경하는 것을 '마샬링' 또는 '인코딩'이라 함
	// interface 타입의 매개 변수를 하나 입력 받음
	// interface는 Go의 모든 타입을 나타내기 때문에 모든 타입을 매개 변수로 사용할 수 있음
	// tuple을 리턴
	data, err := json.Marshal(response)
	if err != nil {
		// 현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}
