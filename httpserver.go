/**
 * example for https://gowebexamples.com/http-server/
 * 本练习重点是：
 * 1. 动态请求的处理；
 * 2. 静态资源，这里以http文件系统为例.
 *
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
	// http.Request 包含了所有动态请求的数据，如：
	// * r.URL.Query().Get("name") 获取通过GET方法传输的名称为"name"的参数值.
	// * r.FormValue("email") 获取通过POST方法传输的HTML FORM的名称为"email"的参数值.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "欢迎来到我的网站！")
	})

	// http.Dir 使用了Web端的文件系统.
	// http.FileServer 是一个Handler，实现了以参数http.Dir为根目录的web文件系统响应.
	fs := http.FileServer(http.Dir("static/"))
	// TODO: http.StripPrefix 是一个Handler，需要深入理解......
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":80", nil)
}
