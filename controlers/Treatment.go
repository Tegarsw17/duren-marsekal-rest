package controlers

import (
	"fmt"
	"log"
	"net/http"
	"rest-duren-marsekal/models"
	"rest-duren-marsekal/service"
	"rest-duren-marsekal/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func GetTreatmentPlant(c *gin.Context) {
	var dt []models.TreatmentView
	var datas models.Plant

	graph := c.Query("graph")
	treatment := c.Query("treatment")
	id_plant := c.Param("id_plant")

	// queries := c.Request.URL.Query()
	// log.Println("All Queries:", queries)

	log.Println(treatment)

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
			AND plant_id = '%s'
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

	if treatment != "" {
		models.DB.Preload("PlantDictionary").Preload("Treatment", "type_treatment=?", treatment).Find(&datas, "id=?", id_plant)
	} else {
		models.DB.Preload("PlantDictionary").Preload("Treatment").Find(&datas, "id=?", id_plant)
	}

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

	dataView := models.TreatmentPlantView{
		ID:            datas.ID,
		Name:          datas.Name,
		Condition:     datas.Condition,
		Longitude:     datas.Longitude,
		Latitude:      datas.Latitude,
		TreatmentView: dt,
	}
	log.Print(graph)
	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   true,
		Message: "All treatment from " + datas.ID,
		Data:    dataView,
	})
}

func CreateTreatmentPlant(c *gin.Context) {
	var payload models.TreatmentCreate
	var data models.Treatment
	id_plant := c.Param("id_plant")

	err := c.ShouldBind(&payload)
	utils.ErrorNotNill(err)

	id := uuid.NewV4().String()

	data.ID = id
	data.TypeTreatment = payload.TypeTreatment
	data.Detail = payload.Detail
	data.PlantId = id_plant
	data.IsDone = false
	data.DateDone = time.Date(2023, time.March, 12, 12, 30, 0, 0, time.Local)
	data.DueDate = time.Now()
	data.ImageUrl = "duren-marsekal/treatment/default"

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

	c.JSON(http.StatusOK, utils.ResponsJsonStruct{
		Error:   true,
		Message: "Success Bro",
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
	data.DueDate = payload.DueDate
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
