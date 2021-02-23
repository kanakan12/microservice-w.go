package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type Response struct {
	Message string
}

// Benchmark => 성능을 측정하는 기능
// 메서드 이름은 항상 Benchmark로 시작하여야 함
// Benchmark 다음에 함수 이름이 오며, 함수 이름의 첫 글자는 항상 대문자로 작성
// 항상 *testing.B타입의 매개변수를 받음
func BenchmarkHelloHandlerVariable(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := Response{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(response)
		fmt.Fprint(writer, string(data))
	}
}

func BenchmarkHelloHandlerEncoder(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := Response{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(response)
	}
}

func BenchmarkHelloHandlerEncoderReference(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := Response{Message: "Hello World"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(&response)
	}
}

// $ go test -v -run="none" -bench=. -benchtime="5s" -benchmem
// 바이트 배열로 마샬링하는 것보다 Encoder를 사용하는 것이 거의 50% 정도 더 빠름
