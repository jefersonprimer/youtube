package handlers

import (
	"net/http"

	"github.com/jefersonprimer/youtube/backend/database"
	"github.com/jefersonprimer/youtube/backend/models"

	"github.com/gin-gonic/gin"
)

// CreateComment cria um novo comentário
func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

// GetCommentByID obtém um comentário pelo ID
func GetCommentByID(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comentário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, comment)
}

// UpdateComment atualiza um comentário existente
func UpdateComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comentário não encontrado"})
		return
	}

	var updatedComment models.Comment
	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.CommentText = updatedComment.CommentText
	comment.ParentCommentID = updatedComment.ParentCommentID

	database.DB.Save(&comment)
	c.JSON(http.StatusOK, comment)
}

// DeleteComment exclui um comentário
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	var comment models.Comment
	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comentário não encontrado"})
		return
	}

	database.DB.Delete(&comment)
	c.JSON(http.StatusNoContent, nil)
}

// GetCommentsByVideoID obtém todos os comentários de um vídeo
func GetCommentsByVideoID(c *gin.Context) {
	videoID := c.Param("video_id")
	var comments []models.Comment
	if err := database.DB.Where("video_id = ?", videoID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}
