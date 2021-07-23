package api

import (
	"testing"
)

func TestSMSValidationPasses(t *testing.T) {

	ins := []*SMS{
		{
			Sender:    "senmder",
			Recipient: "recipient",
			Message:   "msg",
		},
		{
			Sender:    "senmder",
			Recipient: "recipient",
			Message:   "long messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong message",
		},
		{
			Sender:    "61654448894",
			Recipient: "s",
			Message:   "msg",
		},
	}

	for _, in := range ins {
		if err := in.validate(); err != nil {
			t.Errorf("Expected all inputs to be valid but %+v gave error %s", in, err.Error())
		}
	}

}

func TestSMSValidationFails(t *testing.T) {

	ins := []*SMS{
		{
			Recipient: "recipient",
			Message:   "msg",
		},
		{
			Sender:  "senmder",
			Message: "long messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong message",
		},
		{
			Sender:  "senmder",
			Message: "super long messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagesuper long messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagesuper long messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagesuper long messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagesuper long messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong messagelong message",
		},
		{
			Sender:    "61654448894",
			Recipient: "s",
		},
	}

	for _, in := range ins {
		if err := in.validate(); err == nil {
			t.Errorf("Expected %+v to give error but instead got nothing", in)
		}
	}

}
