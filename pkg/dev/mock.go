package dev

import (
	"Flowershop-GoBackend/pkg/db"
	"Flowershop-GoBackend/pkg/models"
)

func InitializeMockData() {
	//Mock data for Flowers
	var flowers []*models.Item
	flowers = append(flowers, &models.Item{ID: 0, Type: "rose", Description: "Rose Arrangement", Price: "30.00", Image: "assets/flowers/rosearrangement.jpg"})
	flowers = append(flowers, &models.Item{ID: 0, Type: "rose", Description: "Rose Arrangement", Price: "30.00", Image: "assets/flowers/rosearrangement.jpg"})
	flowers = append(flowers, &models.Item{ID: 0, Type: "daisy", Description: "Daisy Arrangement", Price: "30.00", Image: "assets/flowers/daisyarrangement.jpg"})
	flowers = append(flowers, &models.Item{ID: 0, Type: "daisy", Description: "Daisy Arrangement", Price: "30.00", Image: "assets/flowers/daisyarrangement.jpg"})
	flowers = append(flowers, &models.Item{ID: 0, Type: "lily", Description: "Lily Arrangement", Price: "30.00", Image: "assets/flowers/lilyarrangement.jpg"})
	flowers = append(flowers, &models.Item{ID: 0, Type: "lily", Description: "Lily Arrangement", Price: "30.00", Image: "assets/flowers/lilyarrangement.jpg"})
	flowers = append(flowers, &models.Item{ID: 0, Type: "carnation", Description: "Carnation Arrangement", Price: "30.00", Image: "assets/flowers/carnationarrangement.jpg"})
	flowers = append(flowers, &models.Item{ID: 0, Type: "carnation", Description: "Carnation Arrangement", Price: "30.00", Image: "assets/flowers/carnationarrangement.jpg"})

	for _, s := range flowers {
		db.CreateItem(s.Type, s.Description, s.Price, s.Image)
	}
}
