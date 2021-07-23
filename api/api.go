package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/burstsms/golang-test/service"
)

type Hello struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	msg := Hello{
		Message: "here be dragons?",
	}

	json.NewEncoder(w).Encode(msg)
}

const smsMaxLen = 160 * 3 // and here I was thinking it was 255
// return this struct for a POST /sms
type SMS struct {
	ID        string  `json:"_id"`
	Cost      float64 `json:"cost"`
	SMSCount  int     `json:"sms_count"`
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Message   string  `json:"message"`
}

func (s SMS) validate() error {
	if s.Sender == "" {

		return fmt.Errorf("Sender is required")
	}
	if s.Recipient == "" {

		return fmt.Errorf("Recipient is required")
	}
	if s.Message == "" {

		return fmt.Errorf("Message is required")
	}

	if len(s.Message) > smsMaxLen {

		return fmt.Errorf("Message is too long!")
	}
	return nil
}

// implement send sms handler here

func SendHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		HelloHandler(w, r)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorHandler(w, http.StatusBadRequest, err)
		return
	}
	userApiKey := r.Header.Get("x-api-key")
	_ = userApiKey // I'm not sure what I'm supposed to be doing with this? We already specify the key when launching, so I'm discarding this
	args := &SMS{}
	err = json.Unmarshal(body, args)
	if err != nil {
		errorHandler(w, http.StatusBadRequest, err)
		return
	}

	if err = args.validate(); err != nil {
		errorHandler(w, http.StatusBadRequest, err)
		return
	}

	c := service.NewClient("localhost", service.Port)
	resp, err := c.SendSMS(args.Sender, args.Recipient, args.Message)
	if err != nil {
		errorHandler(w, http.StatusBadRequest, err)
		return
	}

	ret := SMS{ // just to ensure the return type is the same as the one specified in this file
		ID:        resp.SMS.ID,
		Cost:      resp.SMS.Cost,
		SMSCount:  resp.SMS.SMSCount,
		Sender:    resp.SMS.Sender,
		Recipient: resp.SMS.Recipient,
		Message:   resp.SMS.Message,
	}

	json.NewEncoder(w).Encode(ret)
}

func errorHandler(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "Error: %s\n", err.Error())
}
