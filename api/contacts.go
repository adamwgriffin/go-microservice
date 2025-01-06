package api

import (
	"errors"
	"net/http"

	db "github.com/adamwgriffin/go-microservice/db/sqlc"
	"github.com/gin-gonic/gin"
)

type getContactRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getContact(ctx *gin.Context) {
	var req getContactRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	contact, err := server.store.GetContact(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, contact)
}
