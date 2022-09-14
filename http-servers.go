//https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		fmt.Println("it is a get request")
		fruit := req.FormValue("fruit")
		fmt.Fprintf(w, "%s Got!", fruit)

	case "POST":

		body, _ := ioutil.ReadAll(req.Body)
		if len(body) > 0 {
			fmt.Println("--------------it is a post request--------------")
			fmt.Println(string(body))
			w.Write([]byte("posted success!"))
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", nil)
}
