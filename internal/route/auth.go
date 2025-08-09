package route

import (
	"net/http"

	"github.com/caoyong2619/elotus/internal/route/form"
	"github.com/caoyong2619/elotus/internal/services"
	"github.com/gin-gonic/gin"
)

func NewAuth(authService *services.AuthService) *Auth {
	return &Auth{
		authService: authService,
	}
}

type Auth struct {
	ginEngine   gin.IRouter
	authService *services.AuthService
}

// register
func (a *Auth) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data form.Resigster
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, Error(CodeSuccess, err.Error()))
			return
		}

		err := a.authService.Register(data.Username, data.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Error(CodeError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, Success(CodeSuccess))
	}
}

func (a *Auth) Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data form.Login
		if err := context.BindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, Error(CodeSuccess, err.Error()))
			return
		}

		token, err := a.authService.Login(data.Username, data.Password)
		if err != nil {
			context.JSON(http.StatusInternalServerError, Error(CodeError, err.Error()))
		}

		context.JSON(http.StatusOK, Success(gin.H{"token": token}))
	}
}
