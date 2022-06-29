package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/go-secret/models"
	"github.com/mrinjamul/go-secret/repository"
	"github.com/mrinjamul/go-secret/utils"
)

type Views interface {
	Index(ctx *gin.Context)
	NotFound(ctx *gin.Context)
	ShowMessage(ctx *gin.Context)
	AddMessage(ctx *gin.Context)
}

type views struct {
	messageRepo repository.MessageRepo
}

func (m *views) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Secret — Home",
	})
}

func (m *views) NotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "404.html", gin.H{
		"title": "Secret — Error",
	})
}

func (m *views) ShowMessage(ctx *gin.Context) {
	// get hash from parameter
	hash := ctx.Param("hash")
	message := models.Message{
		Hash: hash,
	}
	message, err := m.messageRepo.GetAndRead(ctx, message)

	if err != nil {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "Secret — 404",
		})
		log.Println(err)
		return
	}
	ctx.HTML(http.StatusOK, "show.html", gin.H{
		"title":    "Secret — Show",
		"username": message.UserName,
		"message":  message.Message,
		"count":    utils.TimeRequiredToRead(message.Message),
	})
}

func (m *views) AddMessage(ctx *gin.Context) {
	var msg models.Message
	// get the form values
	msg.UserName = ctx.PostForm("username")
	msg.Message = ctx.PostForm("message")
	if msg.Message == "" {
		ctx.HTML(http.StatusOK, "404.html", gin.H{
			"title": "Secret — Error",
		})
		return
	}
	// create a new secret
	msg, err := m.messageRepo.Add(ctx, &msg)
	if err != nil {
		ctx.HTML(http.StatusOK, "404.html", gin.H{
			"title": "Secret — Error",
		})
		return
	}
	// get hostname
	hostname := ctx.Request.Host
	ctx.HTML(http.StatusOK, "new.html", gin.H{
		"title": "Secret — New",
		"link":  "http://" + hostname + "/" + msg.Hash,
	})
}

func NewViews(messageRepo repository.MessageRepo) Views {
	return &views{
		messageRepo: messageRepo,
	}
}
