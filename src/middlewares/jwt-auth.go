package middlewares

import (
	"gorestapi/src/services"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		tokenString := authHeader[len(BearerSchema):]

		token, err := services.ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[ID]: ", claims["user_id"])
			log.Println("Claims[Name]: ", claims["username"])
			log.Println("Claims[Admin]: ", claims["user_is_admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])

			session := sessions.Default(c)
			session.Set("user_id", claims["user_id"])
			session.Save()
		} else {
			log.Println("error:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
