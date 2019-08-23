/**
 * example for https://gowebexamples.com/hello-world/
 * @date 2019/8/23
 * @author Neo Lin
 * @email xqlin@qq.com
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handler 处理来自浏览器、HTTP客户端、API等的HTTP连接
	// Handler的签名：func(w http.ResponseWriter, r *http.Request)
	// http.ResponseWriter 写text/html的地方
	// http.Request 包含所有http请求的信息
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, 你的请求：%s\n", r.URL.Path)
	})

	// 启动 http server
	http.ListenAndServe(":80", nil)
}
