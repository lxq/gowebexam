/**
 * example for https://gowebexamples.com/sessions/
 * 本练习重点是：
 * 1. session cookie数据存储;
 * 2. gorilla/sessions库的使用: https://github.com/gorilla/sessions.
 *
 * 通过3个URL模拟是否已登录授权.
 *
 * @date 2019/8/26
 * @author Neo Lin
 * @email xqlin@qq.com
 */

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key 必须是16/24/32位长度的AES-128/AES-192/AES-256
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

const (
	cookieName = "cookie-name"
	authStr    = "authenticated"
)

func secret(w http.ResponseWriter, r *http.Request) {
	ses, _ := store.Get(r, cookieName)

	// 检查是否授权
	if auth, ok := ses.Values[authStr].(bool); !ok || !auth {
		http.Error(w, "禁止访问", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	ses, _ := store.Get(r, cookieName)

	// ...

	// 设置授权
	ses.Values[authStr] = true
	ses.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	ses, _ := store.Get(r, cookieName)

	ses.Values[authStr] = false
	ses.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":80", nil)
}
