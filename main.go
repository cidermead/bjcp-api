package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
  "log"

  "github.com/joho/godotenv"

  "github.com/cidermead/bjcp-api/config"
	"github.com/cidermead/bjcp-api/controller"
	"github.com/cidermead/bjcp-api/include"
	"github.com/cidermead/bjcp-api/model"
)

var db *gorm.DB
var err error

// init is invoked before main()
func init() {
  // loads values from .env into the system
  if err := godotenv.Load(); err != nil {
    log.Print("No .env file found")
  }
}

func main() {
	config := config.InitConfig()

	db = include.InitDB()
	defer db.Close()
	db.AutoMigrate(&model.Post{}, &model.Tag{})

	router := gin.Default()
	router.Use(include.CORS())

	// Non-protected routes
	posts := router.Group("/posts")
	{
		posts.GET("/", controller.GetPosts)
		posts.GET("/:id", controller.GetPost)
		posts.POST("/", controller.CreatePost)
		posts.PUT("/:id", controller.UpdatePost)
		posts.DELETE("/:id", controller.DeletePost)
	}

	questions := router.Group("/questions")
	{
		// questions.GET("/view/:id", controller.GetQuestion)
		questions.GET("/", controller.GetRandom)
		questions.GET("/:exam/", controller.GetRandom)
		questions.GET("/:exam/:topic", controller.GetRandom)
	}

	styles := router.Group("/styles")
	{
		styles.GET("/range", controller.GetStyleRange)
		styles.GET("/question", controller.GetStyleQuestion)
	}

	categories := router.Group("/categories")
	{
		categories.GET("/:id", controller.GetCategory)
	}

	// Protected routes
	// For authorized access, group protected routes using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"username1": "password1",
		"username2": "password2",
		"username3": "password3",
	}))

	// /admin/dashboard endpoint is now protected
	authorized.GET("/dashboard", controller.Dashboard)

	router.Run(":" + config.Server.Port)
}
