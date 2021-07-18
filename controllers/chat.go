package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/holys/initials-avatar"
	// "io"
	"log"
	"net/http"
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
	beego.Controller
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
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
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
	a := avatar.New("./static/fonts/Hiragino_Sans_GB_W3.ttf", 26.0) //./resource/fonts/Hiragino_Sans_GB_W3.ttf
	text := c.Ctx.Input.Param(":text")
	b, _ := a.DrawToBytes(text, 32) //背景的大小
	// beego.Info(text)
	// w http.ResponseWriter, r *http.Request
	// io.Copy(c.Ctx.ResponseWriter, b) // stream实现了io.reader接口
	c.Ctx.Output.Body(b) //流stream的方式
	// now `b` is image data which you can write to file or http stream.
}
