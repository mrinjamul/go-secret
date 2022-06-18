package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/go-secret/models"
	"github.com/mrinjamul/go-secret/repository"
	"github.com/mrinjamul/go-secret/utils"
)

type Message interface {
	Add(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAndRead(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	// views
	Index(ctx *gin.Context)
	NotFound(ctx *gin.Context)
	ShowMessage(ctx *gin.Context)
	AddMessage(ctx *gin.Context)
}

type message struct {
	messageRepo repository.MessageRepo
}

// Add adds a new message
func (m *message) Add(ctx *gin.Context) {
	var message models.Message
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &message)
	if err != nil {
		log.Fatal(err)
	}
	message, err = m.messageRepo.Add(ctx, &message)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "failed to create message",
			"error":   err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, message)
	}
}

// Get gets a message
func (m *message) Get(ctx *gin.Context) {
	// Get id from parameter
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	message := models.Message{
		ID: id,
	}

	message, err = m.messageRepo.Get(ctx, message)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, message)
	}
}

// GetAndRead gets a message and mark it as read
func (m *message) GetAndRead(ctx *gin.Context) {
	// get hash from parameter
	hash := ctx.Param("hash")
	message := models.Message{
		Hash: hash,
	}
	message, err := m.messageRepo.GetAndRead(ctx, message)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, message)
	}
}

// GetAll gets all messages
func (m *message) GetAll(ctx *gin.Context) {
	messages, err := m.messageRepo.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, messages)
}

// Update updates a message
func (m *message) Update(ctx *gin.Context) {
	// Get id from parameter
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	var message models.Message
	err = json.Unmarshal(bytes, &message)
	if err != nil {
		log.Fatal(err)
	}
	message.ID = id
	err = m.messageRepo.Update(ctx, &message)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, message)
	}
}

// Delete deletes a message
func (m *message) Delete(ctx *gin.Context) {
	var message models.Message
	bytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &message)
	if err != nil {
		log.Fatal(err)
	}
	err = m.messageRepo.Delete(ctx, &message)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Message deleted",
		})
	}
}

// For views

func (m *message) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Secret — Home",
	})
}

func (m *message) NotFound(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "404.html", gin.H{
		"title": "Secret — Error",
	})
}

func (m *message) ShowMessage(ctx *gin.Context) {
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

func (m *message) AddMessage(ctx *gin.Context) {
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

// NewMessage returns a new message controller
func NewMessage(messageRepo repository.MessageRepo) Message {
	return &message{
		messageRepo: messageRepo,
	}
}
