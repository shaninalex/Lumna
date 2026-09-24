package ws

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	authc "gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	"gitlab.com/shaninalex/lumna/app/platform/realtime"
)

const (
	writeWait = 1 * time.Second
)

type Feed interface {
	Attach(identityID int) (*realtime.Subscriber, func())
}

type Config struct {
	AllowedOrigins []string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RegisterWebsocketRoute(feed Feed, verifier authc.Verifier, router *gin.Engine) {
	private := router.Group("")
	private.Use(httpx.AuthMiddleware(verifier))
	private.GET("/ws", handle(feed, upgrader))
}

func handle(feed Feed, up websocket.Upgrader) gin.HandlerFunc {
	return func(c *gin.Context) {
		a, _ := actor.From(c.Request.Context())
		conn, err := up.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("Upgrade error: %v", err)
			return
		}
		sub, detach := feed.Attach(a.IdentityID)
		defer detach()

		go readPump(conn)
		writePump(conn, sub.C())
	}
}

func readPump(conn *websocket.Conn) {
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}
		fmt.Println(messageType, message)
	}
}

func writePump(conn *websocket.Conn, in <-chan realtime.Envelope) {
	ping := time.NewTicker(time.Second)

	defer ping.Stop()
	defer conn.Close()

	for {
		select {
		case e, ok := <-in:
			if !ok {
				_ = conn.WriteMessage(websocket.CloseMessage, nil)
				return
			}
			dto, ok := encode(e)
			if !ok {
				continue
			}
			_ = conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteJSON(dto); err != nil {
				return
			}
		case <-ping.C:
			_ = conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				return
			}
		}
	}
}
