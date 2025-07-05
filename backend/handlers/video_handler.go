package handlers

import (
	"net/http"

	"github.com/jefersonprimer/youtube/backend/database"
	"github.com/jefersonprimer/youtube/backend/models"

	"github.com/gin-gonic/gin"
)

// CreateVideo cria um novo vídeo
func CreateVideo(c *gin.Context) {
	var video models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&video)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, video)
}

// GetVideoByID obtém um vídeo pelo ID
func GetVideoByID(c *gin.Context) {
	id := c.Param("id")
	var video models.Video
	if err := database.DB.First(&video, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vídeo não encontrado"})
		return
	}
	c.JSON(http.StatusOK, video)
}

// UpdateVideo atualiza um vídeo existente
func UpdateVideo(c *gin.Context) {
	id := c.Param("id")
	var video models.Video
	if err := database.DB.First(&video, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vídeo não encontrado"})
		return
	}

	var updatedVideo models.Video
	if err := c.ShouldBindJSON(&updatedVideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Atualize os campos do vídeo
	video.Title = updatedVideo.Title
	video.Description = updatedVideo.Description
	video.FileUrl = updatedVideo.FileUrl
	video.ThumbnailUrl = updatedVideo.ThumbnailUrl
	video.DurationSeconds = updatedVideo.DurationSeconds
	video.Visibility = updatedVideo.Visibility

	database.DB.Save(&video)
	c.JSON(http.StatusOK, video)
}

// DeleteVideo exclui um vídeo
func DeleteVideo(c *gin.Context) {
	id := c.Param("id")
	var video models.Video
	if err := database.DB.First(&video, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vídeo não encontrado"})
		return
	}

	database.DB.Delete(&video)
	c.JSON(http.StatusNoContent, nil)
}

// GetAllVideos obtém todos os vídeos
func GetAllVideos(c *gin.Context) {
	var videos []models.Video
	if err := database.DB.Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, videos)
}
