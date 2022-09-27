//https://www.golangprograms.com/example-to-handle-get-and-post-request-in-golang.html

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strconv"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		fmt.Println("it is a get request")
		fruit := req.FormValue("fruit")
		fmt.Fprintf(w, "%s Got!", fruit)

	case "POST":

		//body, _ := ioutil.ReadAll(req.Body)
		//if len(body) > 0 {
		//	fmt.Println("--------------it is a post request--------------")
		//	fmt.Println(string(body))
		//	w.Write([]byte("posted success!"))
		//}

		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			fmt.Println(err)
		}
		bodyData := string(requestDump)
		r, _ := regexp.Compile("{[\\s\\S]*}")
		bodyJson := r.FindString(bodyData)
		fmt.Println("--------------------------------post body-------------------------------")
		fmt.Println(bodyJson)
		fmt.Println("--------------------------------json字段样例-----------------------------")
		jsonData := []byte(bodyJson)
		var v interface{}
		json.Unmarshal(jsonData, &v)
		data := v.(map[string]interface{})

		fmt.Println("Hash_value: ", data["hash"])
		fmt.Println("Is_recovered: ", data["is_recovered"])
		fmt.Println("当前时间：", time.Now())
		fmt.Println("-------------------------------------end---------------------------------\n\n\n")

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	fmt.Println("server serves on 8090...")
	http.ListenAndServe(":8090", nil)
}

func UnixToTime(e string) (datatime time.Time, err error) {
	data, err := strconv.ParseInt(e, 10, 64)
	datatime = time.Unix(data/1000, 0)
	return
}
