package models

import "time"

type Treatment struct {
	ID            string     `json:"id" gorm:"primaryKey"`
	TypeTreatment string     `json:"type_treatment"`
	Detail        string     `json:"detail"`
	PlantId       string     `json:"plant_id"`
	IsDone        bool       `json:"is_done"`
	DateDone      *time.Time `json:"date_done,omitempty"`
	DueDate       *time.Time `json:"due_date,omitempty"`
	ImageUrl      string     `json:"image_url"`
	CreatedAt     time.Time  `gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime"`
}

type TreatmentCreate struct {
	TypeTreatment string `json:"type_treatment" validate:"required"`
	Detail        string `json:"detail"`
	IsDone        bool   `json:"is_done"`
	DueDate       string `json:"due_date,omitempty"`
	DateDone      string `json:"date_done,omitempty"`
}

type TreatmentView struct {
	ID            string     `json:"id" gorm:"primaryKey"`
	TypeTreatment string     `json:"type_treatment"`
	Detail        string     `json:"detail"`
	PlantId       string     `json:"plant_id"`
	IsDone        bool       `json:"is_done"`
	DateDone      *time.Time `json:"date_done,omitempty"`
	DueDate       *time.Time `json:"due_date,omitempty"`
	ImageUrl      string     `json:"image_url"`
}

type TreatmentPlantView struct {
	ID            string          `json:"id" gorm:"primaryKey"`
	Name          string          `json:"name"`
	Condition     string          `json:"condition"`
	Longitude     float32         `json:"longitude"`
	Latitude      float32         `json:"latitude"`
	TreatmentView []TreatmentView `json:"treatment"`
}
