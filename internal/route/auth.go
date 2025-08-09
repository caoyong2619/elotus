package route

import (
	"net/http"

	"github.com/caoyong2619/elotus/internal/route/form"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	GinEngine *gin.IRouter
}

func (a *Auth) Register(c *gin.Context) {
	var data form.Resigster
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, Error(CodeSuccess, err.Error()))
		return
	}

}
