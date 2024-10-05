package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	PublishedAt time.Time          `json:"publishedAt" bson:"publishedAt"`
	User        string             `json:"user"`
	Workout     struct {
		Name      string `json:"name"`
		Exercises []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Category    string `json:"category"`
			Sets        []struct {
				Reps   int `json:"reps"`
				Weight int `json:"weight"`
			} `json:"sets"`
		} `json:"exercises"`
	} `json:"workout"`
}
