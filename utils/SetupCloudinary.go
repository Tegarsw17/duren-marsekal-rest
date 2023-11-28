package utils

import (
	"os"

	"github.com/cloudinary/cloudinary-go"
)

var CloudinaryClient *cloudinary.Cloudinary

func InitCloudinary() {

	cloudinaryURL := os.Getenv("CLOUDINARY_URL")

	client, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		panic("Error creating Cloudinary client: " + err.Error())
	}

	CloudinaryClient = client
}
