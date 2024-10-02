package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Workout struct {
	WorkoutID int       `json:"workout_id"`
	UserID    int       `json:"user_id"`
	Date      time.Time `json:"date"`
	Duration  int       `json:"duration"` // продолжительность в минутах
	Notes     string    `json:"notes"`
}
type WorkoutDetails struct {
	Date   time.Time `json:"date"`
	Notes  string    `json:"notes"`
	Set    string    `json:"set_id"`
	Reps   string    `json:"reps"`
	Weight string    `json:"weight"`
	Name   string    `json:"name"`
}

var db *sql.DB

func main() {
	var err error
	// Подключение к базе данных
	db, err = sql.Open("pgx", "postgres://ioojejfj:J7v5I1sg2kfqk3wJ0okGt7mgthQls43C@dumbo.db.elephantsql.com:5432/ioojejfj")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router := gin.Default()

	// Обработчики маршрутов
	router.POST("/users", createUser)
	router.GET("/users/:id", getUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
	router.GET("/users/:id/workouts", getUserWorkouts)

	// Обработчики для тренировок
	router.POST("/workouts", createWorkout)
	router.GET("/workouts/:id", getWorkout)
	router.PUT("/workouts/:id", updateWorkout)
	router.DELETE("/workouts/:id", deleteWorkout)
	router.GET("/workouts/:id/:work_id", getWorkoutDetails)

	router.Run("localhost:8000")
}

// Создание нового пользователя
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("INSERT INTO Users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// Получение пользователя по ID
func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	err := db.QueryRow("SELECT user_id, username, email, password FROM Users WHERE user_id = $1", id).Scan(&user.UserID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Обновление информации о пользователе
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE Users SET username = $1, email = $2, password = $3 WHERE user_id = $4", user.Username, user.Email, user.Password, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Удаление пользователя
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM Users WHERE user_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Создание новой тренировки
func createWorkout(c *gin.Context) {
	var workout Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	workout.Date = time.Now()
	_, err := db.Exec("INSERT INTO Workouts (user_id, duration, notes, date) VALUES ($1, $2, $3, $4)",
		workout.UserID, workout.Duration, workout.Notes, workout.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, workout)
}

// Получение тренировки по ID
func getWorkout(c *gin.Context) {
	id := c.Param("id")
	var workout Workout
	err := db.QueryRow("SELECT workout_id, user_id, date, duration, notes FROM Workouts WHERE workout_id = $1", id).Scan(
		&workout.WorkoutID, &workout.UserID, &workout.Date, &workout.Duration, &workout.Notes)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout not found"})
		return
	}

	c.JSON(http.StatusOK, workout)
}

// Обновление информации о тренировке
func updateWorkout(c *gin.Context) {
	id := c.Param("id")
	var workout Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Exec("UPDATE Workouts SET user_id = $1, workout_name = $2, date = $3, duration = $4, description = $5 WHERE workout_id = $6",
		workout.UserID, workout.Duration, workout.Notes, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, workout)
}

// Удаление тренировки
func deleteWorkout(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM Workouts WHERE workout_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func getWorkoutDetails(c *gin.Context) {
	userid := c.Param("id")
	workoutid := c.Param("work_id")
	results := []WorkoutDetails{}

	rows, err := db.Query("SELECT we2.date, we2.notes, we.set_id, we.reps, we.weight, e.name  from Workout_Exercises we join workouts we2 on we2.workout_id = we.workout_id join exercises e on we.exercise_id = e.exercise_id where we2.workout_id = $2 and we2.user_id  = $1 order by we2.workout_id",
		userid, workoutid)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		workout := WorkoutDetails{}
		err = rows.Scan(&workout.Date, &workout.Notes, &workout.Set, &workout.Reps, &workout.Weight, &workout.Name)
		if err != nil {
			log.Println(err)
		}
		results = append(results, workout)
	}
	c.JSON(http.StatusOK, results)
}
func getUserWorkouts(c *gin.Context) {
	userid := c.Param("id")
	results := []Workout{}

	rows, err := db.Query("select workout_id, date, duration, notes from workouts w where user_id =$1",
		userid)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		workout := Workout{}
		err = rows.Scan(&workout.WorkoutID, &workout.Date, &workout.Duration, &workout.Notes)
		if err != nil {
			log.Println(err)
		}
		results = append(results, workout)
	}
	c.JSON(http.StatusOK, results)
}
