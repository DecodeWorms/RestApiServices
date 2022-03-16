package util

import "net/http"

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//https://jamboard.google.com/d/10-cyxldvFtC5H4vrVFWCuseNWQ-YBTGsf3YtiJhFzYE/edit?usp=meet_whiteboard
