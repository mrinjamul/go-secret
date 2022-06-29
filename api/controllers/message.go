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
)

type Message interface {
	Add(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAndRead(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
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

// NewMessage returns a new message controller
func NewMessage(messageRepo repository.MessageRepo) Message {
	return &message{
		messageRepo: messageRepo,
	}
}
