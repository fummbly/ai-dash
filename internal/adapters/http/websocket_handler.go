package http

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func (h *ResponseHandler) WebsocketResponse(c echo.Context) error {
	upgrader := websocket.Upgrader{}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	defer ws.Close()

	for {
		err := ws.WriteMessage(websocket.TextMessage, []byte(`<div id="chat_room" hx-swap-oob="beforeend"><p>Hello Client</p></div>`))
		if err != nil {
			c.Logger().Error("Failed to write WS message", "err", err)
		}

		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error("failed to read websocket message", "error", err)
		}

		fmt.Printf("%s\n", msg)

	}
}
