/**
 * example for https://gowebexamples.com/forms/
 * 本练习重点是：
 * 1. form信息POST;
 *
 * @date 2019/8/26
 * @author Neo Lin
 * @email xqlin@qq.com
 */

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 联系信息
type Contact struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// form 方法判断
		if http.MethodPost != r.Method {
			tmpl.Execute(w, nil)
			return
		}

		contact := Contact{
			// 获取Form数据
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// 处理form中的数据
		fmt.Println("contact: %v", contact)

		// 传递结果
		// 临时struct，符合interface.
		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":80", nil)
}
