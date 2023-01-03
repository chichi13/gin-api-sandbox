package controllers

import (
	"gin-api/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

//	@Summary		Responds if the server is alive
//	@Description	Responds if the server is alive
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	schemas.BaseResponse	"BaseResponse"
//	@Router			/health [get]
//	@Tags			system
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, schemas.BaseResponse{
		Success: true,
		Message: "Healthy",
	})
}
