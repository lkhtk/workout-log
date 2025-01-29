package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	PublishedAt time.Time
	UserID      primitive.ObjectID `json:"-" bson:"user_id"`
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
			Type      string  `json:"type"`
			HeartRate int     `json:"heart"`
			Speed     float32 `json:"speed"`
			Distance  float32 `json:"distance"`
			Time      int     `json:"time"`
			Calories  int     `json:"calories"`
		} `json:"cardio"`
	} `json:"workout"`
}
type Measurement struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	UserID          primitive.ObjectID `json:"-" bson:"user_id"`
	MeasurementDate time.Time          `json:"measurement_date" bson:"measurement_date"`
	BodyFat         *float32           `json:"body_fat" bson:"body_fat"`
	BodyWeight      *float32           `json:"body_weight" bson:"body_weight"`
	Neck            *float32           `json:"neck" bson:"neck"`
	Chest           *float32           `json:"chest" bson:"chest"`
	Waist           *float32           `json:"waist" bson:"waist"`
	Hips            *float32           `json:"hips" bson:"hips"`
	UpperArm        *float32           `json:"upperarm" bson:"upperarm"`
	Forearm         *float32           `json:"forearm" bson:"forearm"`
	Thighs          *float32           `json:"thighs" bson:"thighs"`
	Calves          *float32           `json:"calves" bson:"calves"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Picture   string             `bson:"picture"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
