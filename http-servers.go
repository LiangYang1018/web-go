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

		trigger_time := time.Unix(int64(data["trigger_time"].(float64)), 0)
		last_eval_time := time.Unix(int64(data["last_eval_time"].(float64)), 0)
		last_sent_time := time.Unix(int64(data["last_sent_time"].(float64)), 0)
		first_trigger_time := time.Unix(int64(data["first_trigger_time"].(float64)), 0)

		fmt.Println("trigger_time:       ", trigger_time)
		fmt.Println("last_eval_time:     ", last_eval_time)
		fmt.Println("last_sent_time:     ", last_sent_time)
		fmt.Println("first_trigger_time: ", first_trigger_time)

		var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
		fmt.Println("上海时间           :", time.Now().In(cstSh).Format("2006-01-02 15:04:05"))
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
