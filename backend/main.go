package main

import (
	"log"
	"os"

	"github.com/jefersonprimer/youtube/backend/database"
	"github.com/jefersonprimer/youtube/backend/handlers"
	"github.com/jefersonprimer/youtube/backend/models" // Importe todos os seus modelos aqui para o AutoMigrate (se for usar)

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Para carregar variáveis de ambiente de um arquivo .env
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum arquivo .env encontrado, usando variáveis de ambiente do sistema.")
	}

	// Conecta ao banco de dados
	database.ConnectDatabase()

	// Opcional: AutoMigrate para criar/atualizar tabelas no banco de dados
	// Descomente e adicione todos os seus modelos aqui se precisar
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Video{},
		&models.Comment{},
		&models.LikeDislike{},
		&models.Subscription{},
	)
	if err != nil {
		log.Fatalf("Falha ao migrar os modelos do banco de dados: %v", err)
	}

	r := gin.Default()

	// Rotas para Usuários
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", handlers.CreateUser)
		userRoutes.GET("/:id", handlers.GetUserByID)
		userRoutes.PUT("/:id", handlers.UpdateUser)
		userRoutes.DELETE("/:id", handlers.DeleteUser)
		userRoutes.GET("/", handlers.GetAllUsers) // Nova rota para buscar todos os usuários
	}

	// Rotas para Vídeos
	videoRoutes := r.Group("/videos")
	{
		videoRoutes.POST("/", handlers.CreateVideo)
		videoRoutes.GET("/:id", handlers.GetVideoByID)
		videoRoutes.PUT("/:id", handlers.UpdateVideo)
		videoRoutes.DELETE("/:id", handlers.DeleteVideo)
		videoRoutes.GET("/", handlers.GetAllVideos) // Nova rota para buscar todos os vídeos
	}

	// Rotas para Comentários
	commentRoutes := r.Group("/comments")
	{
		commentRoutes.POST("/", handlers.CreateComment)
		commentRoutes.GET("/:id", handlers.GetCommentByID)
		commentRoutes.PUT("/:id", handlers.UpdateComment)
		commentRoutes.DELETE("/:id", handlers.DeleteComment)
		commentRoutes.GET("/video/:video_id", handlers.GetCommentsByVideoID) // Comentários por vídeo
	}

	// Rotas para Curtidas/Não Curtidas
	likesDislikesRoutes := r.Group("/likes-dislikes")
	{
		likesDislikesRoutes.POST("/", handlers.CreateLikeDislike)
		likesDislikesRoutes.GET("/:id", handlers.GetLikeDislikeByID)
		likesDislikesRoutes.DELETE("/:id", handlers.DeleteLikeDislike)
		likesDislikesRoutes.GET("/video/:video_id", handlers.GetVideoLikesDislikes)       // Likes/Dislikes por vídeo
		likesDislikesRoutes.GET("/comment/:comment_id", handlers.GetCommentLikesDislikes) // Likes/Dislikes por comentário
	}

	// Rotas para Inscrições
	subscriptionRoutes := r.Group("/subscriptions")
	{
		subscriptionRoutes.POST("/", handlers.CreateSubscription)
		subscriptionRoutes.GET("/:id", handlers.GetSubscriptionByID)
		subscriptionRoutes.DELETE("/:id", handlers.DeleteSubscription)
		subscriptionRoutes.GET("/user/:user_id", handlers.GetSubscriptionsByUser)         // Inscrições de um usuário
		subscriptionRoutes.GET("/channel/:channel_id", handlers.GetSubscribersForChannel) // Inscritos de um canal
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão
	}
	log.Printf("Servidor Gin rodando na porta %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Falha ao iniciar o servidor Gin: %v", err)
	}
}
