package models

import (
	"os"
	"rest-duren-marsekal/utils"
	"time"

	_ "database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func SetDB() {
	err := godotenv.Load(".env")
	utils.ErrorNotNill(err)
	conn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	// Create a Logrus logger instance
	log := logrus.New()

	// Configure Logrus as the GORM logger
	db.Logger = logger.New(
		log,
		logger.Config{
			SlowThreshold: time.Second, // Adjust as needed
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	Migrate(db)
	DB = db

}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&PlantDictionary{}, &Plant{}, &Treatment{})
	utils.ErrorNotNill(err)

	// seeder for plant dictionary
	var plantDictCount int64
	db.Model(&PlantDictionary{}).Count(&plantDictCount)
	if plantDictCount == 0 {
		SeederPlantDictionary(db)
	}

	var plantCount int64
	db.Model(&Plant{}).Count(&plantCount)
	if plantCount == 0 {
		SeederPlant(db)
	}

	var plantTreatment int64
	db.Model(&Treatment{}).Count(&plantTreatment)
	if plantTreatment == 0 {
		SeederTreatment(db)
	}

}

func SeederPlantDictionary(db *gorm.DB) {
	data := []PlantDictionary{
		{
			ID:        "1",
			Name:      "musangking",
			Detail:    "musangking adalah tanaman dari malaysia",
			Care:      "temperature must above 27 celcius",
			ImageUrl:  "duren-marsekal/dict-plant/default",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			Name:      "bawor",
			Detail:    "bawor adalah tanaman asli indonesia",
			Care:      "kelembapan must in 75%",
			ImageUrl:  "duren-marsekal/dict-plant/default",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, v := range data {
		db.Create(&v)
	}
}

func SeederPlant(db *gorm.DB) {
	data := []Plant{
		{
			ID:                "1",
			Name:              "BWR-02",
			Condition:         "good",
			Longitude:         -7.010911,
			Latitude:          109.603322,
			ImageUrl:          "duren-marsekal/plant/default",
			PlantDictionaryID: "7aa4384a-62ed-4e5e-8ee6-a28dac6be3fe",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		{
			ID:                "2",
			Name:              "PTR-02",
			Condition:         "need attention",
			Longitude:         -7.007796,
			Latitude:          109.602516,
			ImageUrl:          "duren-marsekal/plant/default",
			PlantDictionaryID: "86a7a539-4a0e-4884-8248-95ced663eb24",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
	}

	for _, v := range data {
		db.Create(&v)
	}
}

func SeederTreatment(db *gorm.DB) {
	currentTime := time.Now()
	data := []Treatment{
		{
			ID:            "1",
			TypeTreatment: "Pemupukan",
			Detail:        "pemupukan dilakukan dengan NPK",
			PlantId:       "d0582cb2-01ca-4ad0-ba78-eff56521689d",
			IsDone:        true,
			DateDone:      &currentTime,
			DueDate:       &currentTime,
			ImageUrl:      "duren-marsekal/treatment/default",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            "2",
			TypeTreatment: "Pemupukan",
			Detail:        "pemupukan dilakukan dengan NPK",
			PlantId:       "d0582cb2-01ca-4ad0-ba78-eff56521689d",
			IsDone:        false,
			DateDone:      &currentTime,
			DueDate:       &currentTime,
			ImageUrl:      "duren-marsekal/treatment/default",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            "3",
			TypeTreatment: "Peruning",
			Detail:        "peruning pada cabang bawah",
			PlantId:       "d0582cb2-01ca-4ad0-ba78-eff56521689d",
			IsDone:        true,
			DateDone:      &currentTime,
			DueDate:       &currentTime,
			ImageUrl:      "duren-marsekal/treatment/default",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}
	for _, v := range data {
		db.Create(&v)
	}
}
