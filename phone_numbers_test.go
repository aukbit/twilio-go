package twilio

import (
	"net/url"
	"os"
	"testing"

	"github.com/paulormart/assert"
)

func TestAvailablePhoneNumbers(t *testing.T) {
	client := New(AccountSid(os.Getenv("TWILIO_ACCOUNT_SID")), AuthToken(os.Getenv("TWILIO_AUTH_TOKEN")))

	resp, err := client.GetAvailablePhoneNumbersLocalByCountry("GB", url.Values{"InLocality": []string{"City of London"}})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "/2010-04-01/Accounts/"+os.Getenv("TWILIO_ACCOUNT_SID")+"/AvailablePhoneNumbers/GB/Local.json?InLocality=City+of+London&MmsEnabled=false&SmsEnabled=true&VoiceEnabled=true", resp.Uri)
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
