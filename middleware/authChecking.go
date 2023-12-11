package middleware

import (
	"cms-porto/database"
	"cms-porto/models"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsAuthenticated(c *gin.Context) {
	var User models.UserModel

	// * header checking
	headerToken := c.GetHeader("Authorization")
	if headerToken == "" {
		ErrorResponse(c, http.StatusUnauthorized, "authorization header required")
		return
	}

	// * decode token
	TokenDecode, _ := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	id, _ := primitive.ObjectIDFromHex(TokenDecode.Claims.(jwt.MapClaims)["user_id"].(string))
	err := database.Collections.UserCollection.FindOne(context.Background(), models.UserModel{Id: id}).Decode(&User)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ErrorResponse(c, http.StatusUnauthorized, "invalid token")
			return
		}
		ErrorResponse(c, http.StatusInternalServerError, "something went wrong when check auth data")
		return
	}

	c.Set("user_id", User.Id)
	fmt.Println(User.Id)
	c.Next()
}
