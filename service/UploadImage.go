package service

import (
	"rest-duren-marsekal/utils"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context, fileName string, file interface{}, folderName string, codeFolder string) (pathUrl *uploader.UploadResult, err error) {
	fileName = utils.GenerateName(fileName)

	uploadResult, err := utils.CloudinaryClient.Upload.Upload(
		c,
		file,
		uploader.UploadParams{
			Folder:   folderName,
			PublicID: codeFolder + "-" + fileName,
		},
	)

	return uploadResult, err
}
