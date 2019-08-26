/**
 * example for https://gowebexamples.com/basic-middleware/
 * 本练习重点是：
 * 1. 初步学习自定义实现的Router;
 *
 * @date 2019/8/26
 * @author Neo Lin
 * @email xqlin@qq.com
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":80", nil)
}
