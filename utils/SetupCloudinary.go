package utils

import (
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/joho/godotenv"
)

var CloudinaryClient *cloudinary.Cloudinary

func InitCloudinary() {
	// Load your Cloudinary API credentials from a configuration file or environment variables.
	// config := config.LoadConfig() // You should have a function or package to load your configuration.
	err := godotenv.Load(".env")
	ErrorNotNill(err)
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")

	client, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		panic("Error creating Cloudinary client: " + err.Error())
	}

	CloudinaryClient = client
}
