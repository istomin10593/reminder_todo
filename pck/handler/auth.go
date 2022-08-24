package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/istomin10593/reminder_todo"
)

func (h *Handler) sighUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

}

func (h *Handler) signIn(c *gin.Context) {

}
