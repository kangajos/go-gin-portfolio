package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kangajos/go-gin-portfolio.git/app/controllers"
)

func Portfolio(router *gin.Engine) {

	portfolionController := controllers.PortfolioController{}
	router.GET("/", portfolionController.Index)

	sendEmailController := controllers.SendEmailController{}
	router.POST("/send-message", sendEmailController.Send)
}
