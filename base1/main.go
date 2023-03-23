package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe("localhost:9999", nil))
}


// 最基础的功能展示， 这里函数携带的参数是根据http库里面定义的
func sayHello(w http.ResponseWriter, r *http.Request){
	// 在页面输出展示json
	obj := make(map[string]interface{}, 0)
	obj["username"] = "xiawuyue"
	obj["password"] = "xiaoqizhou"

	// 这里是设置response 的响应头
	w.Header().Set("Content-Type", "application/json")
	// 这里是设置响应头的状态码  ok 就是 200
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(obj); err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Write([]byte("hello world"))
}