package handlers

import (
	"encoding/json"
	"net/http"
)

type Enrollment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ListEnrollments(w http.ResponseWriter, _ *http.Request) {
	items := []Enrollment{{ID: 1, Name: "sample"}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
