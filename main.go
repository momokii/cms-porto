package main

import (
	"cms-porto/database"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/unrolled/secure"
)

func main() {
	// * header protection
	secureMiddleware := secure.New(secure.Options{
		//ContentSecurityPolicy: "default-src 'self'",
		FrameDeny:          true,
		FeaturePolicy:      "camera 'none'; microphone 'none'",
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
	})

	err := godotenv.Load() // * local env load
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// * init mongo connection
	database.InitMongoDB()

	r := gin.Default()
	config := cors.DefaultConfig() // * use cors Default Config so we can modify it if needed later
	r.Use(cors.New(config))        // * cors
	r.Use(func(c *gin.Context) {
		secureMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
		c.Next()
	})

	// * routing

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}

}
