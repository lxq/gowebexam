/* example for https://gowebexamples.com/websockets/
* 本练习重点是：
* 1. websocket基本使用;
* 2. github.com/gorilla/websocket 库使用.
*
* @date 2019/8/26
* @author Neo Lin
* @email xqlin@qq.com
 */

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echo(w http.ResponseWriter, r *http.Request) {
	// conn 是一个与Client连接的WebSocket连接
	conn, _ := upgrader.Upgrade(w, r, nil) // 忽略错误
	for {
		// 读取信息，接收浏览器WebSocket.send()的数据
		msgT, msg, err := conn.ReadMessage()
		if nil != err {
			return
		}

		temp := string(msg)

		fmt.Printf("%s 发送了： %s\n", conn.RemoteAddr(), temp)

		temp = "我收到了：" + temp
		// 向浏览器WebSocket写入数据，由浏览器WebSocket.onmessage()进行接收
		if err = conn.WriteMessage(msgT, []byte(temp)); nil != err {
			return
		}
	}
}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":80", nil)
}
