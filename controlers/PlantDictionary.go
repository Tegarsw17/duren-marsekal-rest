package controlers

import (
	"net/http"
	"rest-duren-marsekal/models"
	"rest-duren-marsekal/service"
	"rest-duren-marsekal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

var validate *validator.Validate

var urlImage string = "https://res.cloudinary.com/daw1nuqgv/image/upload/f_auto,q_auto/v1/"

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func GetAllPlantDictionary(c *gin.Context) {
	// additional add filter
	var data []models.PlantDictionary
	var dataView []models.PlantDictionaryView
	result := models.DB.Find(&data)
	// Data not found
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data Not Found",
		})
		return
	}
	for _, pdx := range data {
		dataView = append(dataView, models.PlantDictionaryView{
			ID:       pdx.ID,
			Name:     pdx.Name,
			Detail:   pdx.Detail,
			Care:     pdx.Care,
			Code:     pdx.Code,
			ImageUrl: urlImage + pdx.ImageUrl,
		})

	}
	// Get All data Success
	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "All Data Plant Dictionary",
		Data:    dataView,
	})
}

// create dict
func CreatePlantDictionary(c *gin.Context) {
	var payload models.PlantDictionaryCreate
	var data models.PlantDictionary
	var count int64

	err := c.ShouldBind(&payload)
	utils.ErrorNotNill(err)

	err = validate.Struct(payload)
	if err != nil {
		var dataError []string
		for _, err := range err.(validator.ValidationErrors) {
			dataError = append(dataError, (err.Field() + ":" + err.Tag()))
		}
		c.JSON(http.StatusBadRequest, utils.ResponsJsonArray{
			Error:   true,
			Message: "Invalid input",
			Data:    dataError,
		})
		return
	}

	models.DB.Model(&data).Where("name = ?", payload.Name).Count(&count)

	if count > 0 {
		c.JSON(http.StatusBadRequest, utils.ResponsJsonString{
			Error:   true,
			Message: "Invalid input",
			Data:    "Name already used",
		})
		return
	}

	id := uuid.NewV4().String()

	data.ID = id
	data.Name = payload.Name
	data.Detail = payload.Detail
	data.Care = payload.Care
	data.ImageUrl = "duren-marsekal/dict-plant/default"
	data.Code = payload.Code

	result := models.DB.Create(&data)

	if result.RowsAffected != 0 {
		c.JSON(http.StatusCreated, utils.ResponsJsonString{
			Error:   false,
			Message: "Data created success",
			Data:    data.ID,
		})
		return
	}
	c.JSON(http.StatusBadRequest, utils.ResponsJson{
		Error:   true,
		Message: "Data is Invalid",
	})

}

// get by id

func GetPlantDictionaryById(c *gin.Context) {
	var data models.PlantDictionary
	id_plant_dictionary := c.Param("id_plant_dictionary")

	models.DB.First(&data, "id=?", id_plant_dictionary)

	if data.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data is Not Found",
		})
		return
	}

	dataView := models.PlantDictionaryView{
		ID:       data.ID,
		Name:     data.Name,
		Detail:   data.Detail,
		Care:     data.Care,
		Code:     data.Code,
		ImageUrl: urlImage + data.ImageUrl,
	}

	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "data found",
		Data:    dataView,
	})
}

// update by id
func UpdatePlantDictionaryById(c *gin.Context) {
	var payload models.PlantDictionaryCreate
	var data models.PlantDictionary

	id_plant_dictionary := c.Param("id_plant_dictionary")

	err := c.ShouldBind(&payload)
	utils.ErrorNotNill(err)

	err = validate.Struct(payload)
	if err != nil {
		var dataError []string
		for _, err := range err.(validator.ValidationErrors) {
			dataError = append(dataError, (err.Field() + ":" + err.Tag()))
		}
		c.JSON(http.StatusBadRequest, utils.ResponsJsonArray{
			Error:   true,
			Message: "Invalid input",
			Data:    dataError,
		})
		return
	}

	models.DB.First(&data, "id=?", id_plant_dictionary)

	if data.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data is Not Found",
		})
		return
	}

	data.Name = payload.Name
	data.Detail = payload.Detail
	data.Care = payload.Care
	data.Code = payload.Code

	// result := models.DB.Save(&data)
	models.DB.Save(&data)

	// log.Print(result.Error)

	c.JSON(http.StatusOK, utils.ResponsJsonString{
		Error:   false,
		Message: "Data is Updated",
		Data:    data.ID + " is Updated",
	})

}

// delete by id
func DeletePlantDictionaryById(c *gin.Context) {
	var data models.PlantDictionary

	id_plant_dictionary := c.Param("id_plant_dictionary")

	models.DB.First(&data, "id=?", id_plant_dictionary)

	if data.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data is Not Found",
		})
		return
	}

	models.DB.Where("id=?", id_plant_dictionary).Delete(&data)

	c.JSON(http.StatusOK, utils.ResponsJsonString{
		Error:   false,
		Message: "data found",
		Data:    data.ID + " succes delete",
	})

}

func UploadImagePlantDictionary(c *gin.Context) {
	var data models.PlantDictionary
	id_plant_dictionary := c.Param("id_plant_dictionary")
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if file == nil {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "file not found",
		})
		return
	}

	folderName := "duren-marsekal/dict-plant/"
	codeFolder := "PD"

	pathUrl, err := service.UploadImage(c, header.Filename, file, folderName, codeFolder)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	models.DB.First(&data, "id=?", id_plant_dictionary)

	data.ImageUrl = pathUrl.PublicID

	models.DB.Save(&data)

	c.JSON(http.StatusOK, utils.ResponsJson{
		Error:   false,
		Message: pathUrl.SecureURL,
	})
}
