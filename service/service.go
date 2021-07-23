package service

import (
	"fmt"
	"strconv"

	"github.com/burstsms/golang-test/rpc"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const Name = "SMS"
const Port = 11702

// the reciever of rpc calls
type SMS struct {
	name   string
	apiKey string
}

type Service struct {
	receiver *SMS
}

type Client struct {
	rpc.Client
}

func (s *Service) Name() string {
	return s.receiver.name
}

func (s *Service) Receiver() interface{} {
	return s.receiver
}

func NewService(apiKey string) rpc.Service {
	return &Service{
		receiver: &SMS{
			name:   Name,
			apiKey: apiKey,
		},
	}
}

func NewClient(host string, port int) *Client {
	return &Client{
		Client: rpc.Client{
			ServiceAddress: host + ":" + strconv.Itoa(port),
			ServiceName:    Name,
		},
	}
}

// rpc server methods

type SendSMSParams struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

type SendSMSReply struct {
	SMS struct {
		ID        string  `json:"_id"`
		Cost      float64 `json:"cost"`
		SMSCount  int     `json:"sms_count"`
		Sender    string  `json:"sender"`
		Recipient string  `json:"recipient"`
		Message   string  `json:"message"`
	} `json:"sms"`
}

// implement sms service method here

func (s *SMS) SendSMS(args *SendSMSParams, retVal *SendSMSReply) error {
	url := "https://api.transmitmessage.com/v1/sms/message"
	method := "POST"

	payload, err := json.Marshal(args)
	if err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("x-api-key", s.apiKey)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed with status: %d", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	resp := &SendSMSReply{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}
	retVal.SMS = resp.SMS

	return nil
}

func (c *Client) SendSMS(sender, recipient, message string) (reply *SendSMSReply, err error) {
	reply = &SendSMSReply{}
	err = c.Call("SendSMS", &SendSMSParams{
		Sender:    sender,
		Recipient: recipient,
		Message:   message,
	}, reply)
	return
}
