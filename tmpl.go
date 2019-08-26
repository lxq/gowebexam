/**
 * example for https://gowebexamples.com/templates/
 * 本练习重点是：
 * 1. html模板: html/template 库;
 *
 * @date 2019/8/29
 * @author Neo Lin
 * @email xqlin@qq.com
 */

package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoData{
			PageTitle: "TODO列表",
			Todos: []Todo{
				{Title: "任务1", Done: true},
				{Title: "任务2", Done: false},
				{Title: "任务3", Done: true},
			},
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
