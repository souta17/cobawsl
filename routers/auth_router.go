package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/souta17/go/configs"
	"github.com/souta17/go/handlers"
	"github.com/souta17/go/repositories"
	"github.com/souta17/go/services"
)

func AuthRouter(api *gin.RouterGroup) {
	// logic yg berkomunikasi dengan database
	authRepository := repositories.NewAuthRepository(configs.DB)

	// logic yg menghendel request
	authService := services.NewAuthService(authRepository)

	// logic untuk mengeksekusi repositori sama service
	authHandler := handlers.NewAuthHandler(authService)
	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
}
