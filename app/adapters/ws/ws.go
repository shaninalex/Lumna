package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gitlab.com/shaninalex/lumna/app/adapters/ws/hub"
	"gitlab.com/shaninalex/lumna/app/core"
)

var upgrader = websocket.Upgrader{
	// Allow all origins for development; restrict this in production.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RegisterWebsocketRoute(resolve core.Resolve, router *gin.Engine) {
	h := hub.New()

	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("Upgrade error: %v", err)
			return
		}
		h.AddClient(conn)
		defer h.RemoveClient(conn)

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				break
			}
			fmt.Println(messageType, message)
		}
	})
}
