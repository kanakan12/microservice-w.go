package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	// 전송된 JSON은 Body 필드에서 접근할 수 있음
	// Body는 io.ReadCloser 인터페이스를 스트림으로 구현
	// []Byte나 문자열을 리턴하지 않음
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// 클라이언트에서 ioutil.ReadAll을 호출했다면, 클라이언트는 자동으로 닫히지 않기 때문에 Close()를 호출해야함
	// ServeHTTP 핸들러에서 사용될 때는 서버가 자동으로 요청 스트림을 닫기 때문에 호출하지 않아도 됨

	var request helloWorldRequest

	// 바이트 슬라이스나 문자열을 논리적 자료 구조로 변경하는 것
	// 입력 받은 객체의 키가 구조체 필드의 이름이나 태크와 일치하는 필드를 찾는데, 이 때 대소문자를 구분하지는 않지만 정확히 일치하는 것이 좋음
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// http.Request 객체 중 일부 설명

// Method는 HTTP 요청 방식을 지정 (GET, POST, PUT 등)
// Mehod string

// Header는 서버가 수신한 요청 헤더 필드를 가지고 있음
// Header는 타입은 map[string] []string에 대한 연결
// Header Header

// Body는 요청의 본문
// Body io.ReadCloser

// 입력 => curl localhost:8080/helloworld -d '{"name":"KAKA"}'
