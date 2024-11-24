package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	PublishedAt time.Time
	User        string `json:"user"`
	Workout     struct {
		MuscleGroup string `json:"muscle_group"`
		SetsCount   int    `json:"sets_count"`
		Exercises   []struct {
			Name string `json:"name"`
			Sets []struct {
				Reps   int `json:"reps"`
				Weight int `json:"weight"`
			} `json:"sets"`
		} `json:"exercises"`
		Cardio []struct {
			Type     string  `json:"type"`
			Speed    int     `json:"speed"`
			Distance float64 `json:"distance"`
			Time     int     `json:"time"`
			Calories int     `json:"calories"`
		} `json:"cardio"`
	} `json:"workout"`
}
