package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

type validationContextKey string

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	handler := newValidationHandler(newHelloWorldHandler())
	http.Handle("/helloworld", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

// validationHandler를 살펴보면 유요한 요청이 있을 때 이 요청에 대한 새 컨텍스트를 만든 다음, 요청의 Name 필드 값을 컨텍스트에 설정한다는 것을 알 수 있음
func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	// WithValue 메서드 호출 등을 통해 컨텍스트에 새로운 항목을 추가하면, 호출 메서드는 이전 컨텍스트의 복사본을 리턴함으로써 약간의 시간을 절약하는 대신 약간의 혼란을 초래
	// 컨텍스트에 대한 포인터를 WithValue의 매개 변수로 전달하면 이 포인터의 복사본이 전달되는데, 이 떄 사용한 포인터를 참조 해제해야 함
	// 포인터를 업데이트하기 위해, 포인터가 리턴되는 값(c)을 참조하도록 해야 하며 원래 가리키던 값은 참조 해제해야 함
	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

type helloWorldHandler struct {
}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloWorldResponse{Message: "Hello " + name}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

func fetchGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://google.com", nil)

	// 요청의 원본으로부터 시간 초과 컨텍스트를 만듦
	// 컨텍스트가 자동으로 취소되는 인바운드 요청과 달리 아웃바운드 요청에서는 취소 단계를 수동으로 수행해야함
	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)

	// 지연실행
	// 특정 문장 혹은 함수를 나중에 (defer를 호출하는 함수가 리턴하기 직전에) 실행하게 됨
	// 일반적으로 defer는 C#, Java 같은 언어에서의 finally 블럭처럼 마지막에 Clean-up 작업을 위해 사용
	defer cancelFunc()

	// http.Request 객체에 추가된 두 개의 새로운 컨텍스트 메서드 중 두 번째를 구현
	// WithContext 객체는 매개 변수로 입력받은 ctx 컨텍스트로 원본 요청의 컨텍스트를 변경한 얕은 복사본을 리턴
	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
