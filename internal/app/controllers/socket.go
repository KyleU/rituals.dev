package controllers

import (
	"net/http"

	"github.com/kyleu/rituals.dev/internal/app/web"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Socket(w http.ResponseWriter, r *http.Request) {
	ctx := web.ExtractContext(w, r)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ctx.Logger.Info("unable to upgrade connection to websocket")
		return
	}

	connID, err := ctx.App.Socket.Register(ctx.Profile.UserID, c)
	if err != nil {
		ctx.Logger.Info("unable to register websocket connection")
		return
	}

	ctx.App.Socket.ReadLoop(connID)
}
