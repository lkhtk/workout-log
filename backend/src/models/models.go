package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	PublishedAt time.Time
	UserID      primitive.ObjectID `bson:"user_id"`
	MuscleGroup string             `json:"muscle_group" bson:"muscle_group"`
	Coach       bool               `json:"coach" bson:"coach"`
	Workout     struct {
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

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Picture   string             `bson:"picture"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
