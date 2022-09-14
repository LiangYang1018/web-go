package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	//var uid string

	var fruit string
	if req.Method == "GET" {
		//uid = req.FormValue("uid")
		fmt.Println("it is a get request")
		fruit = req.FormValue("fruit")
	} else if req.Method == "POST" {
		//uid = req.PostFormValue("uid")
		fmt.Println("it is a post request")
		fruit = req.PostFormValue("fruit")
	}
	fmt.Fprintf(w, "%s receieved!", fruit)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
