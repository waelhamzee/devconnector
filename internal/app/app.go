package app

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/waelhamzee/devconnector/internal/config"
	"github.com/waelhamzee/devconnector/internal/http"
	a "github.com/waelhamzee/devconnector/pkg/auth"
	comm "github.com/waelhamzee/devconnector/pkg/comment"
	conn "github.com/waelhamzee/devconnector/pkg/connection"
	p "github.com/waelhamzee/devconnector/pkg/post"
	u "github.com/waelhamzee/devconnector/pkg/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Run() {
	godotenv.Load()
	cfg := config.Load()
	db, err := gorm.Open(sqlite.Open("devconnector.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&u.User{})
	db.AutoMigrate(&comm.Comment{})
	db.AutoMigrate(&conn.Connection{})
	db.AutoMigrate(&p.Post{})

	router := http.NewRouter(cfg)
	api := router.Group("/api")
	a.RegisterRoutes(api, db)
	u.RegisterRoutes(api, db)
	p.RegisterRoutes(api, db)
	comm.RegisterRoutes(api, db)
	conn.RegisterRoutes(api, db)

	log.Printf("Starting server on port %s", cfg.Port)
	router.Run(":" + cfg.Port)
}
