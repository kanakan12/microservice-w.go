package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}

	// NewEncode => Marshal의 결과를 바이트 배열에 저장하는 대신 HTTP 응답에 바로 쓸 수 있음
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// ResponseWriter
// 세 가지 매서드를 정의하는 인터페이스

// 1. WriteHeader 메서드를 통해서 전송될 헤더들의 맵(map)을 리턴
// Header()

// 2. 연결(connection)에 데이터를 씀
// WriteHeader 메서드가 아직 호출되지 않았다면 Write 메서드는 WriteHeader(http.StatusOK)를 호출
// Write([]byte) (int, error)

// 3. 상태 코드(status code)를 포함한 HTTP 응답 헤더를 전송
// WrieHeader(int)
