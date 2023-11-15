package models

import "time"

type Plant struct {
	ID                string `json:"id" gorm:"primaryKey"`
	Name              string `json:"name"`
	Condition         string `json:"condition"`
	Longitude         string `json:"longitude"`
	Latitude          string `json:"latitude"`
	ImageUrl          string `json:"image_url"`
	PlantDictionaryID string
	PlantDictionary   PlantDictionary
	Treatment         []Treatment
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
}

type PlantCreate struct {
	Condition         string `json:"condition"`
	Longitude         string `json:"longitude"`
	Latitude          string `json:"latitude"`
	PlantDictionaryID string `json:"plant_dictionary_id"`
}

type PlantView struct {
	ID        string              `json:"id" gorm:"primaryKey"`
	Name      string              `json:"name"`
	Condition string              `json:"condition"`
	Longitude string              `json:"longitude"`
	Latitude  string              `json:"latitude"`
	PlantDict PlantDictionaryView `json:"plant_dict,omitempty"`
	ImageUrl  string              `json:"image_url"`
}
