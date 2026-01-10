package ui

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/bmj2728/catfetch/pkg/shared/api"
	"github.com/bmj2728/catfetch/pkg/shared/catdb"
	"github.com/bmj2728/catfetch/pkg/shared/metadata"
)

func HandleButtonClick(db *catdb.CatDB) (image.Image, *metadata.CatMetadata, error) {
	img, md, err := api.RequestRandomCat(30*time.Second, db)
	if err != nil {
		log.Printf("Error fetching image: %v", err)
		return nil, nil, err
	}

	return img, md, nil
}
