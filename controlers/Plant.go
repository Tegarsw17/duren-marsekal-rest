package controlers

import (
	"net/http"
	"rest-duren-marsekal/models"
	"rest-duren-marsekal/service"
	"rest-duren-marsekal/utils"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// var urlImage string = "https://res.cloudinary.com/daw1nuqgv/image/upload/f_auto,q_auto/v1/"

// var urlImage string = "https://res.cloudinary.com/daw1nuqgv/image/upload/v1698663228/"

func GetAllPlant(c *gin.Context) {
	var model []models.Plant
	var modelView []models.PlantView

	result := models.DB.Preload("PlantDictionary").Find(&model)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data Not Found",
		})
		return
	}

	for _, mod := range model {
		modelView = append(modelView, models.PlantView{
			ID:        mod.ID,
			Name:      mod.Name,
			Condition: mod.Condition,
			Longitude: mod.Longitude,
			Latitude:  mod.Latitude,
			PlantDict: models.PlantDictionaryView{
				ID:       mod.PlantDictionary.ID,
				Name:     mod.PlantDictionary.Name,
				Detail:   mod.PlantDictionary.Detail,
				Care:     mod.PlantDictionary.Care,
				ImageUrl: urlImage + mod.PlantDictionary.ImageUrl,
			},
			ImageUrl: urlImage + mod.ImageUrl,
		})
	}

	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "All Data Plant",
		Data:    modelView,
	})
}

func CreatePlant(c *gin.Context) {
	var payload models.PlantCreate
	var data models.Plant
	var dataPlantDict models.PlantDictionary

	err := c.ShouldBind(&payload)
	utils.ErrorNotNill(err)

	models.DB.First(&dataPlantDict, "id=?", payload.PlantDictionaryID)

	id := uuid.NewV4().String()

	data.ID = id
	data.Name = dataPlantDict.Name + " 01"
	data.Condition = payload.Condition
	data.Longitude = payload.Longitude
	data.Latitude = payload.Latitude
	data.PlantDictionaryID = payload.PlantDictionaryID
	data.ImageUrl = "duren-marsekal/plant/default"

	result := models.DB.Create(&data)

	if result.RowsAffected != 0 {
		c.JSON(http.StatusCreated, utils.ResponsJsonString{
			Error:   false,
			Message: "Data success created",
			Data:    data.ID,
		})
		return
	}
	c.JSON(http.StatusBadRequest, utils.ResponsJson{
		Error:   true,
		Message: "Data is Invalid",
	})
}

func GetPlantById(c *gin.Context) {
	var data models.Plant
	id_plant := c.Param("id_plant")

	models.DB.Preload("PlantDictionary").First(&data, "id=?", id_plant)

	if data.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data is Not Found",
		})
		return
	}

	dataView := models.PlantView{
		ID:        data.ID,
		Name:      data.Name,
		Condition: data.Condition,
		Longitude: data.Longitude,
		Latitude:  data.Latitude,
		PlantDict: models.PlantDictionaryView{
			ID:       data.PlantDictionary.ID,
			Name:     data.PlantDictionary.Name,
			Detail:   data.PlantDictionary.Detail,
			Care:     data.PlantDictionary.Care,
			ImageUrl: urlImage + data.PlantDictionary.ImageUrl,
		},
		ImageUrl: urlImage + data.ImageUrl,
	}

	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "data found",
		Data:    dataView,
	})
}

func UpdatePlantById(c *gin.Context) {
	var payload models.PlantCreate
	var data models.Plant

	id_plant := c.Param("id_plant")

	err := c.ShouldBind(&payload)
	utils.ErrorNotNill(err)

	models.DB.Preload("PlantDictionary").First(&data, "id=?", id_plant)

	if data.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data is Not Found",
		})
		return
	}

	data.Condition = payload.Condition
	data.Longitude = payload.Longitude
	data.Latitude = payload.Latitude
	data.PlantDictionaryID = payload.PlantDictionaryID

	result := models.DB.Save(&data)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, utils.ResponsJson{
			Error:   true,
			Message: "Invalid Input",
		})
		return
	}

	c.JSON(http.StatusNotFound, utils.ResponsJsonString{
		Error:   false,
		Message: "Data is Updated",
		Data:    data.ID + " is Updated",
	})
}

func DeletePlantById(c *gin.Context) {
	var data models.Plant

	id_plant := c.Param("id_plant")
	models.DB.First(&data, "id=?", id_plant)

	if data.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data is Not Found",
		})
		return
	}

	models.DB.Where("id=?", id_plant).Delete(&data)

	c.JSON(http.StatusOK, utils.ResponsJsonString{
		Error:   false,
		Message: "data found",
		Data:    data.ID + " succes delete",
	})
}

func UploadImagePlant(c *gin.Context) {
	var data models.Plant
	id_plant := c.Param("id_plant")
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if file == nil {
		c.JSON(http.StatusBadRequest, utils.ResponsJson{
			Error:   true,
			Message: "Input Invalid",
		})
		return
	}

	folderName := "duren-marsekal/plant/"
	codeFolder := "P"

	pathUrl, err := service.UploadImage(c, header.Filename, file, folderName, codeFolder)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	models.DB.First(&data, "id=?", id_plant)

	data.ImageUrl = pathUrl.PublicID

	models.DB.Save(&data)

	c.JSON(http.StatusOK, utils.ResponsJson{
		Error:   false,
		Message: pathUrl.SecureURL,
	})
}
