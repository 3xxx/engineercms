package controllers

import (
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"
	"github.com/holys/initials-avatar"
	// "io"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

var chat_clients = make(map[*websocket.Conn]bool) // connected clients
var chat_broadcast = make(chan ChatMessage)       // broadcast channel

// Configure the upgrader
var chat_upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type ChatMessage struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type ChatController struct {
	web.Controller
}

func init() {
	// Create a simple file server
	// fs := http.FileServer(http.Dir("../public"))
	// http.Handle("/", fs)

	// Configure websocket route
	// http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8000 and log any errors
	// log.Println("http server started on :8000")
	// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}

// @Title get chat
// @Description get chat
// @Success 200 {object} models.UserTemple
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /chat [get]
// 聊天室页面
func (c *ChatController) Chat() {
	c.Data["IsChat"] = true
	c.TplName = "chat.tpl"
}

// @Title get wschat
// @Description get wschat
// @Success 200 {object} models.GetChat
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /wschat [get]
// 用户连接后建立websocket长连接
func (c *ChatController) HandleConnections() {
	// Upgrade initial GET request to a websocket
	ws, err := chat_upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	chat_clients[ws] = true

	for {
		var msg ChatMessage
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(chat_clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		chat_broadcast <- msg
	}
	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "info": "SUCCESS", "title": ""}
	c.ServeJSON()
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-chat_broadcast
		// Send it out to every client that is currently connected
		for client := range chat_clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(chat_clients, client)
			}
		}
	}
}

// @Title get avatar
// @Description get avatar
// @Param text path string  true "The text"
// @Success 200 {object} models.Avatar
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /avatar/:text [get]
// 用户头像，用流stream的方式
func (c *ChatController) Avatar() {
	// 秦修改了源码，支持字的大小，下面第二个参数是字的大小
	// a := avatar.New("./static/fonts/Hiragino_Sans_GB_W3.ttf", 26.0) //./resource/fonts/Hiragino_Sans_GB_W3.ttf
	a := avatar.New("./static/fonts/Hiragino_Sans_GB_W3.ttf")
	text := c.Ctx.Input.Param(":text")
	// beego.Info(text)
	strData, err := url.QueryUnescape(text) //
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(strData)
	b, err := a.DrawToBytes(strData, 32) //背景的大小
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(b)
	// w http.ResponseWriter, r *http.Request
	// io.Copy(c.Ctx.ResponseWriter, b) // stream实现了io.reader接口
	c.Ctx.Output.Body(b) //流stream的方式
	// now `b` is image data which you can write to file or http stream.
	// 写入文件
	// var f *os.File
	// f, err = os.Create(filename) //创建文件
	// if err != nil {
	// 	logs.Error(err)
	// }
	// _, err = f.Write(b)
}

// golang文件流下载
// https://www.bilibili.com/video/BV1Yp4y197Ly
type JsonData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func download(response http.ResponseWriter, request *http.Request) {
	user := JsonData{
		Name: "dewei",
		Age:  33,
	}
	mjson, _ := json.Marshal(user)
	jsonData := string(mjson)
	b := []byte(jsonData)
	fileName := "user.json"
	response.Header().Add("content-type", "application/octet-stream;charset=utf-8")
	response.Header().Add("Content-Disposition", "attachement;filename=\""+fileName+"\"")
	response.Write(b)
}

// https://www.cnblogs.com/haima/p/13442194.html
// golang之UrlEncode编码/UrlDecode解码
// var urlStr string = "傻了吧:%:%@163& .html.html"
// escapeUrl := url.QueryEscape(urlStr)
// fmt.Println("编码:",escapeUrl)

// enEscapeUrl, _ := url.QueryUnescape(escapeUrl)
// fmt.Println("解码:",enEscapeUrl)
// 输出
// 编码: %E5%82%BB%E4%BA%86%E5%90%A7%3A%25%3A%25%40163%26+.html.html
// 解码: 傻了吧:%:%@163& .html.html
