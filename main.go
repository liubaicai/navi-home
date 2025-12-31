package main

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/liubaicai/navi-home/handlers"
	"github.com/liubaicai/navi-home/models"
)

func main() {
	// Initialize database
	models.InitDB()

	// Set gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Session configuration
	secretKey := os.Getenv("SECRET_KEY_BASE")
	if secretKey == "" {
		secretKey = "default-secret-key-for-development"
	}
	store := cookie.NewStore([]byte(secretKey))
	r.Use(sessions.Sessions("navi_session", store))

	// Load templates
	r.LoadHTMLGlob("templates/*")

	// Static files
	r.Static("/static", "./static")
	r.Static("/icons", "./public/icons")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")
	r.StaticFile("/ie.png", "./public/ie.png")

	// Routes
	r.GET("/", handlers.Index)
	r.GET("/search", handlers.Search)
	r.POST("/set_catalogs", handlers.SetCatalogs)
	r.GET("/test_icon_url", handlers.TestIconURL)

	// Auth routes
	r.GET("/users/sign_in", handlers.SignInPage)
	r.POST("/users/sign_in", handlers.SignIn)
	r.GET("/users/sign_up", handlers.SignUpPage)
	r.POST("/users/sign_up", handlers.SignUp)
	r.GET("/users/sign_out", handlers.SignOut)
	r.GET("/users/edit", handlers.EditUserPage)
	r.POST("/users/edit", handlers.EditUser)

	// Get port from environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
