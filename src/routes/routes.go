package routes

import (
	"gorestapi/src/controllers"
	"gorestapi/src/middlewares"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// SET Storage session
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("gorestSession", store))

	r.POST("login", controllers.LogIn)
	protectEndpoint := r.Group("/api", middlewares.AuthorizeJWT())
	{
		protectEndpoint.GET("users", controllers.GetUsers)
		protectEndpoint.POST("user", controllers.CreateUser)
		protectEndpoint.GET("user/:id", controllers.GetUserByID)
		protectEndpoint.GET("userinfo", controllers.GetUserInfo)
		protectEndpoint.PUT("user/:id", controllers.UpdateUser)
		protectEndpoint.DELETE("user/:id", controllers.DeleteUser)
	}
	publicEndpoint := r.Group("/public/api")
	{
		publicEndpoint.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello world")
		})
	}
	return r
}
