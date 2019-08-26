/**
 * example for https://gowebexamples.com/static-files/
 * 本练习重点是：
 * 1. 使用文件服务提供URL与Server端目录的灵活对应;
 * 2. 即URL的路径可以与本地目录不一致，通过http.Dir()来设定.
 *
 * @date 2019/8/29
 * @author Neo Lin
 * @email xqlin@qq.com
 */

package main

import "net/http"

func main() {
	// http.Dir()指定的是Server端的本地目录
	fs := http.FileServer(http.Dir("assets/"))

	// http.Handle()第1个参数指定的是URL相对地址
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
