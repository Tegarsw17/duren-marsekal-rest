package routers

import (
	"rest-duren-marsekal/controlers"

	"github.com/gin-gonic/gin"
)

func PlantRouter(r *gin.Engine) {
	r.GET("/plant", controlers.GetAllPlant)
	r.POST("/plant", controlers.CreatePlant)
	r.GET("/plant/:id_plant", controlers.GetPlantById)
	r.PUT("/plant/:id_plant", controlers.UpdatePlantById)
	r.DELETE("/plant/:id_plant", controlers.DeletePlantById)
	r.POST("/plant/:id_plant/upload-images", controlers.UploadImagePlant)
}
