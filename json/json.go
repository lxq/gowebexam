/**
 * example for https://gowebexamples.com/json/
 * 本练习重点是：
 * 1. encoding/json 库的使用;
 * 2. 如何读写http的json数据.
 *
 * @date 2019/8/26
 * @author Neo Lin
 * @email xqlin@qq.com
 */

// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// User 用户信息.
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

var (
	tmpl = template.Must(template.ParseFiles("json.html"))
)

func enc(w http.ResponseWriter, r *http.Request) {
	neo := User{
		Firstname: "Neo",
		Lastname:  "Lin",
		Age:       40,
	}

	json.NewEncoder(w).Encode(neo)
}

func dec(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost != r.Method {
		tmpl.Execute(w, nil)
		return
	}

	// 以下获取数据
	var usr User
	// 前端Form数据无法通过这种方式获取，只能通过 Request.FormValue()获取 .
	json.NewDecoder(r.Body).Decode(&usr)

	fmt.Fprintf(w, "%s %s 有%d岁了！", usr.Firstname, usr.Lastname, usr.Age)
}

func main() {
	// template

	http.HandleFunc("/enc", enc)
	http.HandleFunc("/dec", dec)

	http.ListenAndServe(":80", nil)
}
