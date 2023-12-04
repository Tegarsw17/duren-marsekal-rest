package routers

import (
	"rest-duren-marsekal/controlers"

	"github.com/gin-gonic/gin"
)

func TreatmentRouter(r *gin.Engine) {
	r.GET("/plant/:id_plant/treatment", controlers.GetTreatmentPlant)
	r.POST("/plant/:id_plant/treatment", controlers.CreateTreatmentPlant)
	r.GET("/plant/:id_plant/treatment/:id_treatment", controlers.GetTreatmentPlantById)
	r.PUT("/plant/:id_plant/treatment/:id_treatment", controlers.UpdateTreatmentPlantById)
	r.DELETE("/plant/:id_plant/treatment/:id_treatment", controlers.DeleteTreatmentPlantById)
	r.POST("/plant/:id_plant/treatment/:id_treatment", controlers.UploadImageTreatment)

	r.GET("/treatment", controlers.GetAllTreatmentPlant)
}
