package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqoi/todo-apis-go/server/controllers/views"
)

func WriteResponse(c *gin.Context, resp *views.Response) {
	c.JSON(resp.StatusCode, resp)
}
