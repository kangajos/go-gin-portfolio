package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kangajos/go-gin-portfolio.git/app/models"
)

type PortfolioController struct {
}

func (pc *PortfolioController) Index(ctx *gin.Context) {
	portfolio := models.Portfolio{}
	ctx.HTML(http.StatusOK, "pages/home/portfolio", portfolio.Get())
}
