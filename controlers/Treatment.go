package controlers

import (
	"fmt"
	"net/http"
	"rest-duren-marsekal/models"
	"rest-duren-marsekal/service"
	"rest-duren-marsekal/utils"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetTreatmentPlant(c *gin.Context) {
	var dt []models.TreatmentView
	var datas models.Plant

	graph := c.Query("graph")
	treatment := c.Query("treatment")
	id_plant := c.Param("id_plant")
	is_done := c.Query("is_done")

	if graph == "true" {
		query := fmt.Sprintf(`
		SELECT
			TO_CHAR(DATE_TRUNC('month', date_done), 'Month') AS month,
			COUNT(CASE WHEN type_treatment = 'Pemupukan' THEN 1 ELSE NULL END) AS count_pemupukan,
			COUNT(CASE WHEN type_treatment = 'Penyiraman' THEN 1 ELSE NULL END) AS count_penyiraman,
			COUNT(CASE WHEN type_treatment = 'Kocor' THEN 1 ELSE NULL END) AS count_kocor,
			COUNT(CASE WHEN type_treatment = 'Peruning' THEN 1 ELSE NULL END) AS count_peruning,
			COUNT(CASE WHEN type_treatment = 'Semprot' THEN 1 ELSE NULL END) AS count_semprot,
			COUNT(CASE WHEN type_treatment = 'Bersih Gulma' THEN 1 ELSE NULL END) AS count_bersih_gulma
		FROM treatments
		WHERE (type_treatment = 'Pemupukan' OR type_treatment = 'Penyiraman' OR type_treatment = 'Kocor' OR type_treatment = 'Peruning' OR type_treatment = 'Semprot' OR type_treatment = 'Bersih Gulma') 
			AND plant_id = '%s' AND is_done=TRUE
		GROUP BY DATE_TRUNC('month', date_done)
		ORDER BY DATE_TRUNC('month', date_done);
		`, id_plant)

		var result []struct {
			Month            string `json:"month"`
			CountPemupukan   int    `json:"count_pemupukan"`
			CountPenyiraman  int    `json:"count_penyiraman"`
			CountKocor       int    `json:"count_kocor"`
			CountPeruning    int    `json:"count_peruning"`
			CountSemprot     int    `json:"count_semprot"`
			CountBersihGulma int    `json:"count_bersih_gulma"`
		}

		models.DB.Raw(query).Scan(&result)

		if result == nil {
			c.JSON(http.StatusNotFound, utils.ResponsJson{
				Error:   true,
				Message: "Data not found",
			})
			return
		}

		for i := range result {
			result[i].Month = strings.TrimSpace(result[i].Month)
		}

		c.JSON(http.StatusOK, utils.ResponsJsonStruct{
			Error:   false,
			Message: "count monthly",
			Data:    result,
		})
		return
	}

	query := models.DB.Model(&datas).Where("id = ?", id_plant)

	if treatment != "" {
		query = query.Preload("PlantDictionary").Preload("Treatment", "type_treatment=?", treatment)
	} else {
		query = query.Preload("PlantDictionary").Preload("Treatment")
	}

	if is_done == "false" {
		if treatment != "" {
			query = query.Preload("Treatment", "is_done = ? AND type_treatment = ?", false, treatment)
		} else {
			query = query.Preload("PlantDictionary").Preload("Treatment", "is_done=?", false)
		}
	}

	query.Find(&datas)

	for _, dts := range datas.Treatment {

		dt = append(dt, models.TreatmentView{
			ID:            dts.ID,
			TypeTreatment: dts.TypeTreatment,
			Detail:        dts.Detail,
			PlantId:       dts.PlantId,
			IsDone:        dts.IsDone,
			DateDone:      dts.DateDone,
			DueDate:       dts.DueDate,
			ImageUrl:      urlImage + dts.ImageUrl,
		})
	}

	if dt == nil {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data not found",
		})
		return
	}

	dataView := models.TreatmentPlantView{
		ID:            datas.ID,
		Name:          datas.Name,
		Condition:     datas.Condition,
		Longitude:     datas.Longitude,
		Latitude:      datas.Latitude,
		TreatmentView: dt,
	}

	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "All treatment from " + datas.Name,
		Data:    dataView,
	})
}

func CreateTreatmentPlant(c *gin.Context) {
	var payload models.TreatmentCreate
	var data models.Treatment
	id_plant := c.Param("id_plant")
	request_treatment := c.Query("req_treat")
	// log.Print(request_treatment)

	err := c.ShouldBind(&payload)
	utils.ErrorNotNill(err)

	id := uuid.NewV4().String()

	data.ID = id
	data.TypeTreatment = payload.TypeTreatment
	data.Detail = payload.Detail
	data.PlantId = id_plant
	data.ImageUrl = "duren-marsekal/treatment/default"

	if request_treatment == "true" {
		data.IsDone = false
		data.DueDate = utils.GenerateDate(payload.DueDate)
		data.DateDone = nil

		result := models.DB.Create(&data)

		if result.RowsAffected != 0 {
			c.JSON(http.StatusCreated, utils.ResponsJsonString{
				Error:   false,
				Message: "Data success created",
				Data:    data.ID,
			})
			return
		}
		return
	}

	data.IsDone = true
	data.DateDone = utils.GenerateDate(payload.DateDone)
	data.DueDate = nil

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

func GetTreatmentPlantById(c *gin.Context) {
	var data models.Plant
	var dt []models.TreatmentView
	// var dt models.Treatment

	id_plant := c.Param("id_plant")
	id_treatment := c.Param("id_treatment")

	models.DB.Preload("Treatment", "id=?", id_treatment).First(&data, "id=?", id_plant)
	if data.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data is Not Found",
		})
		return
	}

	fmt.Print(data.Treatment)
	for _, dts := range data.Treatment {
		dt = append(dt, models.TreatmentView{
			ID:            dts.ID,
			TypeTreatment: dts.TypeTreatment,
			Detail:        dts.Detail,
			PlantId:       dts.PlantId,
			IsDone:        dts.IsDone,
			DateDone:      dts.DateDone,
			DueDate:       dts.DueDate,
			ImageUrl:      urlImage + dts.ImageUrl,
		})
	}

	if dt == nil {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   false,
		Message: "Data found",
		Data:    dt,
	})

}

func UpdateTreatmentPlantById(c *gin.Context) {
	var payload models.TreatmentCreate
	var data models.Treatment

	id_plant := c.Param("id_plant")
	id_treatment := c.Param("id_treatment")

	err := c.ShouldBind(&payload)
	utils.ErrorNotNill(err)

	models.DB.First(&data, "id=? and plant_id=?", id_treatment, id_plant)

	if data.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data is Not Found",
		})
		return
	}

	data.IsDone = payload.IsDone
	data.DueDate = utils.GenerateDate(payload.DateDone)
	data.Detail = payload.Detail
	data.TypeTreatment = payload.TypeTreatment

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

func DeleteTreatmentPlantById(c *gin.Context) {
	var data models.Treatment

	id_plant := c.Param("id_plant")
	id_treatment := c.Param("id_treatment")

	models.DB.First(&data, "id=? and plant_id=?", id_treatment, id_plant)
	if data.ID == "" {
		c.JSON(http.StatusNotFound, utils.ResponsJson{
			Error:   true,
			Message: "Data is Not Found",
		})
		return
	}

	models.DB.Where("id=? and plant_id=?", id_treatment, id_plant).Delete(&data)

	c.JSON(http.StatusOK, utils.ResponsJsonString{
		Error:   false,
		Message: "data found",
		Data:    data.ID + " succes delete",
	})
}

func UploadImageTreatment(c *gin.Context) {
	var data models.Treatment

	id_plant := c.Param("id_plant")
	id_treatment := c.Param("id_treatment")

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

	folderName := "duren-marsekal/treatment/"
	codeFolder := "T"

	pathUrl, err := service.UploadImage(c, header.Filename, file, folderName, codeFolder)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	models.DB.First(&data, "id=? and plant_id=?", id_treatment, id_plant)

	data.ImageUrl = pathUrl.PublicID

	models.DB.Save(&data)

	c.JSON(http.StatusOK, utils.ResponsJson{
		Error:   false,
		Message: pathUrl.SecureURL,
	})
}
