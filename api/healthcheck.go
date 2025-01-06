package api

import "github.com/gin-gonic/gin"

func (server *Server) ping(ctx *gin.Context) {
	ctx.String(200, "pong")
}
