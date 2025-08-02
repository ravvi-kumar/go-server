package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "created a new workout")
}

func (wh *WorkoutHandler) GetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")
	if params == "" {
		http.NotFound(w, r)
		return
	}
	workoutID, err := strconv.ParseInt(params, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "workout id: %d", workoutID)
}
