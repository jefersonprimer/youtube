package handlers

import (
	"net/http"

	"github.com/jefersonprimer/youtube/backend/database"
	"github.com/jefersonprimer/youtube/backend/models"

	"github.com/gin-gonic/gin"
)

// CreateSubscription cria uma nova inscrição
func CreateSubscription(c *gin.Context) {
	var subscription models.Subscription
	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prevenir que um usuário se inscreva no próprio canal
	if subscription.SubscriberUserID == subscription.ChannelUserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Não é possível se inscrever no próprio canal"})
		return
	}

	result := database.DB.Create(&subscription)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, subscription)
}

// GetSubscriptionByID obtém uma inscrição pelo ID
func GetSubscriptionByID(c *gin.Context) {
	id := c.Param("id")
	var subscription models.Subscription
	if err := database.DB.First(&subscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inscrição não encontrada"})
		return
	}
	c.JSON(http.StatusOK, subscription)
}

// DeleteSubscription exclui uma inscrição
func DeleteSubscription(c *gin.Context) {
	id := c.Param("id")
	var subscription models.Subscription
	if err := database.DB.First(&subscription, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inscrição não encontrada"})
		return
	}

	database.DB.Delete(&subscription)
	c.JSON(http.StatusNoContent, nil)
}

// GetSubscriptionsByUser obtém as inscrições de um usuário (em quais canais ele se inscreveu)
func GetSubscriptionsByUser(c *gin.Context) {
	userID := c.Param("user_id")
	var subscriptions []models.Subscription
	if err := database.DB.Where("subscriber_user_id = ?", userID).Find(&subscriptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}

// GetSubscribersForChannel obtém os inscritos de um canal
func GetSubscribersForChannel(c *gin.Context) {
	channelID := c.Param("channel_id")
	var subscriptions []models.Subscription
	if err := database.DB.Where("channel_user_id = ?", channelID).Find(&subscriptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}
