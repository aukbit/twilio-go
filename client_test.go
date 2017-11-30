package twilio

import (
	"os"
	"testing"

	"github.com/paulormart/assert"
)

type FakeResponse struct{}

func TestClientResponseError(t *testing.T) {
	client := New(AccountSid(os.Getenv("TWILIO_ACCOUNT_SID")), AuthToken(os.Getenv("TWILIO_AUTH_TOKEN")))

	var f FakeResponse

	err := client.get(&f, "wrong", nil)
	if err == nil {
		t.Fatalf("an error should have been returned")
	}
	assert.Equal(t, 500, err.Status)
	assert.Equal(t, 0, err.Code)
	assert.Equal(t, true, len(err.Message) > 0)
}
