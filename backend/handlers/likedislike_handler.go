package handlers

import (
	"net/http"

	"github.com/jefersonprimer/youtube/backend/database"
	"github.com/jefersonprimer/youtube/backend/models"

	"github.com/gin-gonic/gin"
)

// CreateLikeDislike cria uma nova curtida/não curtida
func CreateLikeDislike(c *gin.Context) {
	var likeDislike models.LikeDislike
	if err := c.ShouldBindJSON(&likeDislike); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificação para garantir que ou video_id ou comment_id é fornecido, mas não ambos
	if (likeDislike.VideoID == nil && likeDislike.CommentID == nil) || (likeDislike.VideoID != nil && likeDislike.CommentID != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Deve fornecer VideoID ou CommentID, mas não ambos"})
		return
	}

	result := database.DB.Create(&likeDislike)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, likeDislike)
}

// GetLikeDislikeByID obtém uma curtida/não curtida pelo ID
func GetLikeDislikeByID(c *gin.Context) {
	id := c.Param("id")
	var likeDislike models.LikeDislike
	if err := database.DB.First(&likeDislike, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curtida/Não Curtida não encontrada"})
		return
	}
	c.JSON(http.StatusOK, likeDislike)
}

// DeleteLikeDislike exclui uma curtida/não curtida
func DeleteLikeDislike(c *gin.Context) {
	id := c.Param("id")
	var likeDislike models.LikeDislike
	if err := database.DB.First(&likeDislike, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curtida/Não Curtida não encontrada"})
		return
	}

	database.DB.Delete(&likeDislike)
	c.JSON(http.StatusNoContent, nil)
}

// GetVideoLikesDislikes obtém todas as curtidas/não curtidas de um vídeo
func GetVideoLikesDislikes(c *gin.Context) {
	videoID := c.Param("video_id")
	var likesDislikes []models.LikeDislike
	if err := database.DB.Where("video_id = ?", videoID).Find(&likesDislikes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, likesDislikes)
}

// GetCommentLikesDislikes obtém todas as curtidas/não curtidas de um comentário
func GetCommentLikesDislikes(c *gin.Context) {
	commentID := c.Param("comment_id")
	var likesDislikes []models.LikeDislike
	if err := database.DB.Where("comment_id = ?", commentID).Find(&likesDislikes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, likesDislikes)
}
