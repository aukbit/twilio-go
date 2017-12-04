package twilio

import (
	"fmt"
	"net/url"
)

// Messages Endpoint

// Messages is a representation of a twilio Response
type Messages struct {
	Messages        []*Messages `json:"messages,omitempty"`
	Uri             string      `json:"uri,omitempty"`
	FirstPageUri    string      `json:"first_page_uri,omitempty"`
	End             int         `json:"end,omitempty"`
	PreviousPageUri string      `json:"previous_page_uri,omitempty"`
	PageSize        int         `json:"page_size,omitempty"`
	Start           int         `json:"start,omitempty"`
	NextPageUri     string      `json:"next_page_uri,omitempty"`
	Page            int         `json:"page,omitempty"`
}

type Message struct {
	Sid                 string            `json:"sid,omitempty"`
	AccountSid          string            `json:"account_sid,omitempty"`
	DateCreated         string            `json:"date_created,omitempty"`
	DateUpdated         string            `json:"date_updated,omitempty"`
	DateSent            string            `json:"date_sent,omitempty"`
	To                  string            `json:"to,omitempty"`
	From                string            `json:"from,omitempty"`
	MessagingServiceSid string            `json:"messaging_service_sid,omitempty"`
	Body                string            `json:"body,omitempty"`
	Status              string            `json:"status,omitempty"`
	NumSegments         string            `json:"num_segments,omitempty"`
	NumMedia            string            `json:"num_media,omitempty"`
	Direction           string            `json:"direction,omitempty"`
	ApiVersion          string            `json:"api_version,omitempty"`
	Price               string            `json:"price,omitempty"`
	PriceUnit           string            `json:"price_unit,omitempty"`
	ErrorCode           string            `json:"error_code,omitempty"`
	ErrorMessage        string            `json:"error_message,omitempty"`
	Uri                 string            `json:"uri,omitempty"`
	SubresourceUris     map[string]string `json:"subresource_uris,omitempty"`
}

// SendMessage make POST request on twilio api
// endpoint: /Accounts/[AccountSid]/IncomingPhoneNumbers/Local
func (c *Client) SendMessage(formValues url.Values) (*Message, *ResponseError) {

	var i Message
	if err := c.post(&i, fmt.Sprintf("%s/Messages.json", c.accountSid), formValues); err != nil {
		return nil, err
	}
	return &i, nil

}
