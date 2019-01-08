package token_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/aukbit/twilio-go/token"
	"github.com/paulormart/assert"
)

func TestExample(t *testing.T) {
	tk := token.New("ACxxxxxxxxxxxx", "SKxxxxxxxxxxxx", "xxxxxxxxxxxxxx", "test@example.com", time.Hour)
	grant := token.NewVideoGrant("a-video-room")
	tk.AddGrant(grant)
	jwt, _ := tk.JWT()
	fmt.Println(jwt) // A string encoded with the given values.
	assert.Equal(t, true, jwt != "")
}
