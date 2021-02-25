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

type helloWorldRequest struct {
	Name string `json:"name"`
}

type Sometwhing struct{}

func main() {
	port := 8080

	cathandler := http.FileServer(http.Dir("./images"))
	// StripPrefix => HTTP 요청 경로에서 특정 문자열 제거
	http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))
	//http.Handle("/cat/", cathandler)

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

// http://127.0.0.1:8080/cat/cat.jpg

// http.Handle("/images/", newFooHanlder())
// http.Handle("/images/persian/", newBarHandler())
// http.Handle("/images", newBuzzHandler())

// /images                  => Buzz
// /images 				    => Foo
// /images/cat 			    => Foo
// /images/cat.jpg          => Foo
// /images/persian/cat.jpg  => Bar
