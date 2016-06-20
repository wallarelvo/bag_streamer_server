package main

import (
    "github.com/gin-gonic/gin"
    "github.com/olahol/melody"
)

type BagStreamer struct {
    r *gin.Engine
    m *melody.Melody
}

func MakeBagStreamer() BagStreamer {
    return BagStreamer{gin.Default(), melody.New()}
}

func (bs BagStreamer) WsHandler(c *gin.Context) {
    bs.m.HandleRequest(c.Writer, c.Request)
}

func (bs BagStreamer) MessageHandler(s *melody.Session, msg []byte) {
    bs.m.Broadcast(msg)
}

func (bs BagStreamer) Run(host, port string) {
    bs.r.Run(host + ":" + port)
}

func main() {
    bs := MakeBagStreamer()
    bs.r.GET("/ws", bs.WsHandler)
    bs.m.HandleMessage(bs.MessageHandler)
    bs.Run("", "5000")
}
