package api

import (
	"net/http"
	"time"

	"github.com/PolakoJambo/Final-project/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "id is required"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{"error": "task not found"})
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeJSON(w, map[string]string{"error": "failed to delete task"})
			return
		}
		writeJSON(w, map[string]any{})
		return
	}

	now := time.Now()
	next, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	err = db.UpdateDate(id, next)
	if err != nil {
		writeJSON(w, map[string]string{"error": "failed to update date"})
		return
	}

	writeJSON(w, map[string]any{})
}
