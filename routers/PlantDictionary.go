package routers

import (
	"rest-duren-marsekal/controlers"

	"github.com/gin-gonic/gin"
)

func PlantDictionaryRouter(r *gin.Engine) {
	r.GET("/plant-dictionary", controlers.GetAllPlantDictionary)
	r.GET("/plant-dictionary/:id_plant_dictionary", controlers.GetPlantDictionaryById)
	r.POST("/plant-dictionary", controlers.CreatePlantDictionary)
	r.PUT("/plant-dictionary/:id_plant_dictionary", controlers.UpdatePlantDictionaryById)
	r.DELETE("/plant-dictionary/:id_plant_dictionary", controlers.DeletePlantDictionaryById)
	r.POST("/plant-dictionary/:id_plant_dictionary/upload-images", controlers.UploadImagePlantDictionary)
}
