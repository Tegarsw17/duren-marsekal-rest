package models

import (
	"mime/multipart"
	"time"
)

type PlantDictionary struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Detail    string    `json:"detail"`
	Care      string    `json:"care"`
	ImageUrl  string    `json:"image_url"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

type PlantDictionaryView struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Detail   string `json:"detail"`
	Care     string `json:"care"`
	Code     string `json:"code"`
	ImageUrl string `json:"image_url"`
}

type PlantDictionaryCreate struct {
	Name   string `json:"name" validate:"required"`
	Detail string `json:"detail" validate:"required"`
	Care   string `json:"care" validate:"required"`
	Code   string `json:"code" validate:"required"`
}

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}
