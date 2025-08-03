package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Raziur306/kon-dol/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetIncidents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := []model.Incident{
		{
			ID:       primitive.NewObjectID(),
			Title:    "Violence during protest",
			Location: "Dhaka",
			Party:    model.PartyBNP,
			Date:     "2025-08-01",
		},
		{
			ID:       primitive.NewObjectID(),
			Title:    "March disrupted by unknown group",
			Location: "Chittagong",
			Party:    model.PartyUnknown,
			Date:     "2025-07-30",
		},
	}

	json.NewEncoder(w).Encode(data)
}
