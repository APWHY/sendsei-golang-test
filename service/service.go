package service

import (
	"strconv"

	"github.com/burstsms/golang-test/rpc"
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
	Sender    string
	Recipient string
	Message   string
}

type SendSMSReply struct {
}

// implement sms service method here

func (c *Client) SendSMS(sender, recipient, message string) (reply *SendSMSReply, err error) {
	reply = &SendSMSReply{}
	err = c.Call("SendSMS", &SendSMSParams{
		Sender:    sender,
		Recipient: recipient,
		Message:   message,
	}, reply)
	return
}
