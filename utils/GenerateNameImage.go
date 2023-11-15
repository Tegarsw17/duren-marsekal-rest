package utils

import (
	"strings"
	"time"
)

func GenerateName(fileName string) string {
	currentDate := time.Now()
	dateStr := currentDate.Format("02012006150405")
	extensions := []string{".jpg", ".jpeg", ".png", ".pdf"}

	for _, ext := range extensions {
		fileName = strings.TrimSuffix(fileName, ext)
	}
	return fileName + "-" + dateStr
}
