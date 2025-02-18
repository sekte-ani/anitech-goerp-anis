package utils

import (
	"encoding/json"
	"log"
)

func DecodeImages(imagesStr string) ([]string, error) {
	var images []string
	if err := json.Unmarshal([]byte(imagesStr), &images); err != nil {
		log.Printf("Error unmarshaling images: %v", err)
		return nil, err
	}
	return images, nil
}
