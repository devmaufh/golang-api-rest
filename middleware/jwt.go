package middleware

import (
	// "github.com/devmaufh/golang-api-rest/services/"

	"net/http"
	"strings"

	"github.com/devmaufh/golang-api-rest/services"
	"github.com/gin-gonic/gin"
)

//AuthorizeJWT validates token in request header
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := ctx.GetHeader("Authorization")
		splitToken := strings.Split(authHeader, BearerSchema)
		if len(splitToken) != 2 {
			ctx.AbortWithStatusJSON(401, map[string]string{"error": "Bearer token not in proper format"})
		}
		reqToken := strings.TrimSpace(splitToken[1])
		token, _ := services.JWTAuthService().ValidateToken(reqToken)
		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error": "You don't have access to this resource",
			})
		}
	}

}
