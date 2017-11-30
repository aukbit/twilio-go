package twilio

import (
	"fmt"
	"net/url"
)

// PhoneNumbersResponse is a representation of a twilio Response
type PhoneNumbersResponse struct {
	AvailablePhoneNumbers []*PhoneNumber `json:"available_phone_numbers,omitempty"`
	Uri                   string         `json:"uri,omitempty"`
}

type PhoneNumber struct {
	FriendlyName        string        `json:"friendly_name,omitempty"`
	PhoneNumber         string        `json:"phone_number,omitempty"`
	Lata                string        `json:"lata,omitempty"`
	RateCenter          string        `json:"rate_center,omitempty"`
	Latitude            string        `json:"latitude,omitempty"`
	Longitude           string        `json:"longitude,omitempty"`
	Locality            string        `json:"locality,omitempty"`
	Region              string        `json:"region,omitempty"`
	PostalCode          string        `json:"postal_code,omitempty"`
	IsoCountry          string        `json:"iso_country,omitempty"`
	AddressRequirements string        `json:"address_requirements,omitempty"`
	Beta                bool          `json:"beta,omitempty"`
	Capabilities        *Capabilities `json:"capabilities,omitempty"`
}

type Capabilities struct {
	Voice bool `json:"voice,omitempty"`
	SMS   bool `json:"sms,omitempty"`
	MMS   bool `json:"mms,omitempty"`
	Fax   bool `json:"fax,omitempty"`
}

// GetAvailablePhoneNumbersLocalByCountry make GET request on twilio api
func (c *Client) GetAvailablePhoneNumbersLocalByCountry(countryIsoCode string, queryValues url.Values) (*PhoneNumbersResponse, *ResponseError) {
	// SmsEnabled
	defaultQueryValues := url.Values{
		"SmsEnabled":   []string{"true"},
		"VoiceEnabled": []string{"true"},
		"MmsEnabled":   []string{"false"},
	}
	for k, v := range queryValues {
		defaultQueryValues[k] = v
	}
	var p PhoneNumbersResponse
	if err := c.get(&p, fmt.Sprintf("%s/AvailablePhoneNumbers/%s/Local.json", c.accountSid, countryIsoCode), defaultQueryValues); err != nil {
		return nil, err
	}
	return &p, nil
}
