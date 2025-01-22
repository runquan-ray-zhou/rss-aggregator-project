package main

import (
	"encoding/json"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
}
