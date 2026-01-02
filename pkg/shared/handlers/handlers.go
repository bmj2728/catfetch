package handlers

import (
	"bytes"
	"image"
	"log"
	"time"

	"go-gui/pkg/shared/api"
)

func HandleButtonClick() (image.Image, error) {
	data, err := api.RequestCat(30 * time.Second)
	if err != nil {
		log.Printf("Error fetching image: %v", err)
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Printf("Error decoding image: %v", err)
		return nil, err
	}
	return img, nil
}
