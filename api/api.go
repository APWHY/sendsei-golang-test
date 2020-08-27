package api

import (
	"encoding/json"
	"net/http"
)

type Hello struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	msg := Hello{
		Message: "here be dragons",
	}

	json.NewEncoder(w).Encode(msg)
}

// return this struct for a POST /sms
type SMS struct {
	ID        string  `json:"_id"`
	Cost      float64 `json:"cost"`
	SMSCount  int     `json:"sms_count"`
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Message   string  `json:"message"`
}

// implement send sms handler here
