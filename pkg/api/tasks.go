package api

import (
	"net/http"

	"github.com/PolakoJambo/Final-project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50)
	if err != nil {
		writeJSON(w, map[string]string{"error": "database error"})
		return
	}

	writeJSON(w, TasksResp{Tasks: tasks})
}
