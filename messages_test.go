package twilio

import (
	"net/url"
	"os"
	"testing"

	"github.com/paulormart/assert"
)

func TestSendMessage(t *testing.T) {
	client := New(AccountSid(os.Getenv("TWILIO_ACCOUNT_SID")), AuthToken(os.Getenv("TWILIO_AUTH_TOKEN")))

	available, err := client.GetIncomingPhoneNumbersLocal(nil)
	if err != nil {
		t.Fatal(err)
	}

	formValues := url.Values{
		"To":   []string{os.Getenv("MOBILE")},
		"From": []string{available.IncomingPhoneNumbers[0].PhoneNumber},
		"Body": []string{"test send message"},
	}

	resp, err := client.SendMessage(formValues)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, resp.Sid != "")
}
