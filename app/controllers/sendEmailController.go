package controllers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
)

type SendEmailController struct {
}

func (sec SendEmailController) Send(ctx *gin.Context) {
	// name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	toList := []string{email}
	// subject := ctx.PostForm("subject")
	message := ctx.PostForm("message")

	from := os.Getenv("MAIL_FROM_ADDRESS")
	password := os.Getenv("MAIL_PASSWORD")

	host := os.Getenv("MAIL_HOST")

	port := os.Getenv("MAIL_PORT")

	msg := message
	auth := smtp.PlainAuth("", from, password, host)
	body := []byte(msg)

	// fmt.Println("Host:", host)
	// fmt.Println("Port:", port)
	// fmt.Println("From:", from)
	// fmt.Println("Auth:", auth)
	// fmt.Println("To:", toList)
	// fmt.Println("Message:", message)
	// fmt.Println("Body:", body)

	if err := smtp.SendMail(host+":"+port, auth, from, toList, body); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"status": "success"})
}
