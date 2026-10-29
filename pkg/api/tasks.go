package api

import (
	"log"
	"net/http"

	"github.com/PolakoJambo/Final-project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		log.Printf("DB error: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	writeJSON(w, http.StatusOK, TasksResp{Tasks: tasks})
}
