/**
 * example for https://gowebexamples.com/routes-using-gorilla-mux/
 * 本练习重点是：
 * 1. 使用库：https://github.com/gorilla/ 实现router；
 * 2. net/http 对于处理复杂请求不是做得很好.
 *
 * @date 2019/8/23
 * @author Neo Lin
 * @email xqlin@qq.com
 */
package main

import (
	"fmt"
	"net/http"

	// gorilla/mux 可深入学习和使用.
	"github.com/gorilla/mux"
)

func main() {
	// 根Router，注册到Server中。
	// Router分发HTTP请求给对应的Handler。
	r := mux.NewRouter()

	// Router中注册Handler
	// {title}、{page}是点位符，用于获取动态的URL路径。
	// mux.Vars 以http.Request作为参数，并返回参数map。
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "你在阅读《%s》的第%s页。", title, page)
	})

	http.ListenAndServe(":80", r)
}
