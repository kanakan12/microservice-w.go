package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Go에서 struct는 Custom Data Type을 표현하는데 사용
// Go의 struct는 필드들의 집합체이며 필드들의 컨테이너
// struct는 필드 데이타만을 가지며 메소드를 갖지 않음
type helloWorldResponse struct {
	// 구조체 필드의 태그를 사용하면 출력이 어떻게 표시되는지 제어 가능
	// 해당과 같은 방법으로 소문자 출력 가능
	// 출력 필드를 "message"로 바꿈
	Message string `json:"message"`

	// 해당 필드를 출력하지 않음
	Author string `json:"-"`

	// 값이 비어있으면 해당 필들를 출력하지 않음
	Date string `json:",omitempty"`

	// 출력을 문자열로 변환하고 이름을 "id"로 변환
	ID int `json:"id,string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	//response := helloWorldResponse{Message: "HelloWorld"}
	response := helloWorldResponse{Message: "HelloWorld", Author: "author_Test", Date: "date_Test", ID: 11111}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}
