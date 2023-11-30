package main

import (
	"net/http"
	"rest-duren-marsekal/models"
	"rest-duren-marsekal/routers"
	"rest-duren-marsekal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	models.SetDB()
	utils.InitCloudinary()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // Change this to the origin(s) you want to allow
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"origin","Content-Type"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, utils.ResponsJson{
			Error:   false,
			Message: "Welcome to Duren Marsekal API this is provide for handle management in the garden of durian",
		})
	})

	routers.PlantDictionaryRouter(r)
	routers.PlantRouter(r)
	routers.TreatmentRouter(r)
	r.Run("localhost:8080")
}
