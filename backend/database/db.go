package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("DATABASE_URL não definida")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar com o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")

	// Opcional: AutoMigrate para criar as tabelas se não existirem
	// Certifique-se de importar seus modelos aqui
	// err = DB.AutoMigrate(&models.User{}, &models.Video{}, &models.Comment{}, &models.LikeDislike{}, &models.Subscription{})
	// if err != nil {
	// 	log.Fatalf("Falha ao migrar os modelos do banco de dados: %v", err)
	// }
}
