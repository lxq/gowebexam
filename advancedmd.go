/**
 * example for https://gowebexamples.com/advanced-middleware/
 * 本练习重点是：
 * 1. 学习自定义蹭件的实现;
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
	"time"
)

// Middleware 自定义中间件的实现函数
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging 记录所有请求信息
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			s := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(s))
			}()

			f(w, r)
		}
	}
}

// Method 确定只能调用带特定方法的URL，否则返回404错误.
func Method(m string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != m {
				// http.StatusText()定义了http的返回状态信息.
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w, r)
		}
	}
}

// CallChain 调用链
func CallChain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Hello test func.
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Golang!")
}

func main() {
	// http.HandleFunc是一个函数，http.HandlerFunc是一个函数类型.
	http.HandleFunc("/", CallChain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":80", nil)
}
