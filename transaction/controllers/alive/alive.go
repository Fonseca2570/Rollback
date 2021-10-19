package alive

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Alive(c *gin.Context) {
	c.JSON(http.StatusOK, AliveResponse{
		"Alive",
	})
}

type AliveResponse struct {
	Msg string `json:"message"`
}
