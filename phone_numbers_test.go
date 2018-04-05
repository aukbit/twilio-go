package twilio

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/paulormart/assert"
)

func TestAvailablePhoneNumbers(t *testing.T) {
	client := New(AccountSid(os.Getenv("TWILIO_ACCOUNT_SID")), AuthToken(os.Getenv("TWILIO_AUTH_TOKEN")))

	resp, err := client.GetAvailablePhoneNumbersLocalByCountry("GB", url.Values{"InLocality": []string{"Brigton"}})
	// resp, err := client.GetAvailablePhoneNumbersLocalByCountry("GB", nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "/2010-04-01/Accounts/"+os.Getenv("TWILIO_ACCOUNT_SID")+"/AvailablePhoneNumbers/GB/Local.json?MmsEnabled=false&SmsEnabled=true&VoiceEnabled=true", resp.Uri)
	assert.Equal(t, true, len(resp.AvailablePhoneNumbers) > 0)
}

func TestIncomingPhoneNumbers(t *testing.T) {
	client := New(AccountSid(os.Getenv("TWILIO_ACCOUNT_SID")), AuthToken(os.Getenv("TWILIO_AUTH_TOKEN")))

	resp, err := client.GetIncomingPhoneNumbersLocal(nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "/2010-04-01/Accounts/"+os.Getenv("TWILIO_ACCOUNT_SID")+"/IncomingPhoneNumbers/Local.json?PageSize=50&Page=0", resp.Uri)
	assert.Equal(t, true, len(resp.IncomingPhoneNumbers) > 0)
}

func TestCreateAnAvailableNumber(t *testing.T) {
	client := New(AccountSid(os.Getenv("TWILIO_ACCOUNT_SID")), AuthToken(os.Getenv("TWILIO_AUTH_TOKEN")))
	available, err := client.GetAvailablePhoneNumbersLocalByCountry("GB", url.Values{"InLocality": []string{"City of London"}})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("available number: %v", available.AvailablePhoneNumbers[0].PhoneNumber)
	formValues := url.Values{
		"PhoneNumber":  []string{available.AvailablePhoneNumbers[0].PhoneNumber},
		"FriendlyName": []string{"TEST"},
		"SmsMethod":    []string{"POST"},
		"SmsUrl":       []string{"http://google.com"},
		"VoiceMethod":  []string{"POST"},
		"VoiceUrl":     []string{"http://google.com"},
	}
	incoming, err := client.CreateIncomingPhoneNumbersLocal(formValues)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, incoming.Sid != "")
}

func TestDeleteAllIncomingNumbers(t *testing.T) {
	client := New(AccountSid(os.Getenv("TWILIO_ACCOUNT_SID")), AuthToken(os.Getenv("TWILIO_AUTH_TOKEN")))
	l, err := client.GetIncomingPhoneNumbersLocal(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("incoming numbers: %v", l.IncomingPhoneNumbers)
	for _, n := range l.IncomingPhoneNumbers {
		client.DeleteIncomingPhoneNumber(n.Sid)
		if err != nil {
			t.Fatal(err)
		}
	}
}
