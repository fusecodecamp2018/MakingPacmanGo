package Lesson_3

import (
	"net/http"
	"fmt"
	"log"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func main()  {
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}