package models

import "time"

type Plant struct {
	ID                string  `json:"id" gorm:"primaryKey"`
	Name              string  `json:"name"`
	Condition         string  `json:"condition"`
	Longitude         float32 `json:"longitude"`
	Latitude          float32 `json:"latitude"`
	ImageUrl          string  `json:"image_url"`
	PlantDictionaryID string
	PlantDictionary   PlantDictionary
	Treatment         []Treatment
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}

type PlantCreate struct {
	Code              string  `json:"code" validate:"required"`
	Condition         string  `json:"condition" validate:"required"`
	Longitude         float32 `json:"longitude" validate:"required,gte=-180,lte=180"`
	Latitude          float32 `json:"latitude" validate:"required,gte=-90,lte=90"`
	PlantDictionaryID string  `json:"plant_dictionary_id" validate:"required"`
}

type PlantView struct {
	ID        string              `json:"id" gorm:"primaryKey"`
	Name      string              `json:"name"`
	Condition string              `json:"condition"`
	Longitude float32             `json:"longitude"`
	Latitude  float32             `json:"latitude"`
	PlantDict PlantDictionaryView `json:"plant_dict,omitempty"`
	ImageUrl  string              `json:"image_url"`
}
