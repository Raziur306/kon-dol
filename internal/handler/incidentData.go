package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Raziur306/kon-dol/internal/db"
	"github.com/Raziur306/kon-dol/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func GetIncidents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _, defaultCollection := db.ConnectDB()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := defaultCollection.Find(ctx, bson.M{}) // empty filter = all documents
	if err != nil {
		log.Fatal("Failed to find documents:", err)
	}
	defer cursor.Close(ctx)

	var incidents []model.GPTResponse
	if err := cursor.All(ctx, &incidents); err != nil {
		log.Fatal("Cursor decoding failed:", err)
	}

	json.NewEncoder(w).Encode(incidents)
}
