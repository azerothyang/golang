package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			fmt.Fprintf(writer, "an error" , err)
		}
		form := request.Form
		if numbers, err := form["numbers"]; err {
			fmt.Fprintf(writer, numbers[0])
		}
	})
	if err := http.ListenAndServe("0.0.0.0:8000", nil) ; err != nil {
		log.Fatal("fail to start server", err)
	}
}


