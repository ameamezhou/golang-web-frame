package compare

import (
	"encoding/json"
	"fmt"
	"net/http"
	xiawuyue "xiawuyue/base3/xiawuyue_base3"
)

func test(w http.ResponseWriter, req *http.Request){
	// 封装前
	obj := map[string]interface{}{
		"name": "xiaoqizhou",
		"password": "1234",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(obj); err != nil {
		fmt.Println("编码错误")
		// 后续开发一个 log 包
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// 封装后
	c := xiawuyue.NewContext(w, req)
	c.Json(http.StatusOK, xiawuyue.Z{
		"name": c.PostForm("name"),
		"password": c.PostForm("password"),
	})
}
