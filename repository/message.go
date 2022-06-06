package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/go-secret/models"
)

type MessageRepo interface {
	Add(ctx *gin.Context, msg *models.Message) error
	Get(ctx *gin.Context, msg models.Message) (models.Message, error)
	GetAndRead(ctx *gin.Context, msg models.Message) (models.Message, error)
	GetAll(ctx *gin.Context) ([]models.Message, error)
	Update(ctx *gin.Context, msg *models.Message) error
	Delete(ctx *gin.Context, msg *models.Message) error
}

type messageRepo struct {
	db gorm.DB
}

// Add a new Message
func (repo *messageRepo) Add(ctx *gin.Context, msg *models.Message) error {
	// check if message is valid
	if msg.Message == "" {
		return errors.New("message is empty")
	}
	// if username is empty, set it to anonymous
	if msg.UserName == "" {
		msg.UserName = "anonymous"
	}
	// check if message is already exists
	result := repo.db.Find(&msg, "message = ?", msg.Message)
	if result.Error != nil {
		return result.Error
	}
	if msg.ID > 0 {
		// then update the message to unread
		msg.Deleted = false
		// update the message
		result = repo.db.Save(&msg)
		if result.Error != nil {
			return result.Error
		}
		return errors.New("message already exists")
	}
	result = repo.db.Omit("ID").Create(&msg)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get a message by id
func (repo *messageRepo) Get(ctx *gin.Context, msg models.Message) (models.Message, error) {
	result := repo.db.Find(&msg, "id = ?", msg.ID)
	if result.Error != nil {
		return msg, result.Error
	}
	return msg, nil
}

// GetAndRead get a message by url
func (repo *messageRepo) GetAndRead(ctx *gin.Context, msg models.Message) (models.Message, error) {
	// get message by hash
	result := repo.db.Find(&msg, "hash = ?", msg.Hash)
	if result.Error != nil {
		return msg, result.Error
	}
	if msg.ID == 0 {
		return msg, errors.New("message not found")
	}
	if msg.Deleted {
		return msg, errors.New("message is deleted")
	}
	msg.Deleted = true
	result = repo.db.Save(&msg)
	if result.Error != nil {
		return msg, result.Error
	}
	return msg, nil
}

// GetAll get all messages
func (repo *messageRepo) GetAll(ctx *gin.Context) ([]models.Message, error) {
	var messages []models.Message
	result := repo.db.Find(&messages)
	if result.Error != nil {
		return messages, result.Error
	}
	return messages, nil
}

// Update a message
func (repo *messageRepo) Update(ctx *gin.Context, msg *models.Message) error {
	var existingMessage models.Message
	result := repo.db.Find(&existingMessage, "id = ?", msg.ID)
	if result.Error != nil {
		return result.Error
	}
	if existingMessage.ID == 0 {
		return errors.New("message not found")
	}
	if msg.Message != "" {
		existingMessage.Message = msg.Message
	}
	if msg.UserName != "" {
		existingMessage.UserName = msg.UserName
	}
	result = repo.db.Save(&existingMessage)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete a message
func (repo *messageRepo) Delete(ctx *gin.Context, msg *models.Message) error {
	// delete where id = msg.id
	// check if msg.ID is not empty
	if msg.ID != 0 {
		result := repo.db.Delete(&msg, "id = ?", msg.ID)
		if result.Error != nil {
			return result.Error
		}
	}
	if msg.Hash != "" {
		result := repo.db.Delete(&msg, "hash = ?", msg.Hash)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func NewMessageRepo(db *gorm.DB) MessageRepo {
	return &messageRepo{db: *db}
}