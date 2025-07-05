package handlers

import (
	"net/http"

	"github.com/jefersonprimer/youtube/backend/database" // Substitua pelo caminho correto
	"github.com/jefersonprimer/youtube/backend/models"   // Substitua pelo caminho correto

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt" // Para hash de senha
)

// CreateUser cria um novo usuário
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash da senha antes de salvar
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao hash da senha"})
		return
	}
	user.PasswordHash = string(hashedPassword)

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetUserByID obtém um usuário pelo ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser atualiza um usuário existente
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Se a senha for fornecida, faça o hash
	if updatedUser.PasswordHash != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao hash da senha"})
			return
		}
		user.PasswordHash = string(hashedPassword)
	}

	// Atualize os outros campos
	user.Username = updatedUser.Username
	user.Email = updatedUser.Email
	user.ProfilePictureURL = updatedUser.ProfilePictureURL
	user.ChannelName = updatedUser.ChannelName
	user.Description = updatedUser.Description

	database.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser exclui um usuário
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusNoContent, nil) // 204 No Content para exclusão bem-sucedida
}

// GetAllUsers obtém todos os usuários
func GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
