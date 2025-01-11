package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	PublishedAt time.Time
	User        string `json:"user" bson:"user"`
	MuscleGroup string `json:"muscle_group" bson:"muscle_group"`
	Coach       bool   `json:"coach" bson:"coach"`
	Workout     struct {
		SetsCount int `json:"sets_count" bson:"sets_count"`
		Exercises []struct {
			Name string `json:"name" bson:"name"`
			Sets []struct {
				Reps   *int     `json:"reps"`
				Weight *float32 `json:"weight"`
			} `json:"sets"`
		} `json:"exercises"`
		Cardio []struct {
			Type     string  `json:"type"`
			Speed    float32 `json:"speed"`
			Distance float32 `json:"distance"`
			Time     int     `json:"time"`
			Calories int     `json:"calories"`
		} `json:"cardio"`
	} `json:"workout"`
}
