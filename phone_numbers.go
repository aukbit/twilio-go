package twilio

import (
	"fmt"
	"net/url"
)

type Capabilities struct {
	Voice bool `json:"voice,omitempty"`
	SMS   bool `json:"sms,omitempty"`
	MMS   bool `json:"mms,omitempty"`
	Fax   bool `json:"fax,omitempty"`
}

// AvailablePhoneNumbers Endpoint

// AvailablePhoneNumbers is a representation of a twilio Response
type AvailablePhoneNumbers struct {
	AvailablePhoneNumbers []*AvailablePhoneNumber `json:"available_phone_numbers,omitempty"`
	Uri                   string                  `json:"uri,omitempty"`
}

type AvailablePhoneNumber struct {
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

// GetAvailablePhoneNumbersLocalByCountry make GET request on twilio api
// endpoint: /Accounts/[AccountSid]/AvailablePhoneNumbers/[CountryCode]/Local
func (c *Client) GetAvailablePhoneNumbersLocalByCountry(countryIsoCode string, queryValues url.Values) (*AvailablePhoneNumbers, *ResponseError) {

	defaultQueryValues := url.Values{
		"SmsEnabled":   []string{"true"},
		"VoiceEnabled": []string{"true"},
		"MmsEnabled":   []string{"false"},
	}
	for k, v := range queryValues {
		defaultQueryValues[k] = v
	}
	var p AvailablePhoneNumbers
	if err := c.get(&p, fmt.Sprintf("%s/AvailablePhoneNumbers/%s/Local.json", c.accountSid, countryIsoCode), defaultQueryValues); err != nil {
		return nil, err
	}
	return &p, nil

}

// IncomingPhoneNumbers Endpoint

// IncomingPhoneNumbers is a representation of a twilio Response
type IncomingPhoneNumbers struct {
	IncomingPhoneNumbers []*IncomingPhoneNumber `json:"incoming_phone_numbers,omitempty"`
	Uri                  string                 `json:"uri,omitempty"`
	FirstPageUri         string                 `json:"first_page_uri,omitempty"`
	End                  int                    `json:"end,omitempty"`
	PreviousPageUri      string                 `json:"previous_page_uri,omitempty"`
	PageSize             int                    `json:"page_size,omitempty"`
	Start                int                    `json:"start,omitempty"`
	NextPageUri          string                 `json:"next_page_uri,omitempty"`
	Page                 int                    `json:"page,omitempty"`
}

type IncomingPhoneNumber struct {
	Sid                   string        `json:"sid,omitempty"`
	AccountSid            string        `json:"account_sid,omitempty"`
	FriendlyName          string        `json:"friendly_name,omitempty"`
	PhoneNumber           string        `json:"phone_number,omitempty"`
	VoiceUrl              string        `json:"voice_url,omitempty"`
	VoiceMethod           string        `json:"voice_method,omitempty"`
	VoiceFallbackUrl      string        `json:"voice_fallback_url,omitempty"`
	VoiceFallbackMethod   string        `json:"voice_fallback_method,omitempty"`
	VoiceCallerIdLookup   bool          `json:"voice_caller_id_lookup,omitempty"`
	DateCreated           string        `json:"date_created,omitempty"`
	DateUpdated           string        `json:"date_updated,omitempty"`
	SmsUrl                string        `json:"sms_url,omitempty"`
	SmsMethod             string        `json:"sms_method,omitempty"`
	SmsFallbackUrl        string        `json:"sms_fallback_url,omitempty"`
	SmsFallbackMethod     string        `json:"sms_fallback_method,omitempty"`
	AddressRequirements   string        `json:"address_requirements,omitempty"`
	Beta                  bool          `json:"beta,omitempty"`
	Capabilities          *Capabilities `json:"capabilities,omitempty"`
	VoiceReceive_mode     string        `json:"voice_receive_mode,omitempty"`
	StatusCallback        string        `json:"status_callback,omitempty"`
	AtatusCallback_method string        `json:"status_callback_method,omitempty"`
	ApiVersion            string        `json:"api_version,omitempty"`
	VoiceApplicationSid   string        `json:"voice_application_sid,omitempty"`
	SmsApplicationSid     string        `json:"sms_application_sid,omitempty"`
	Origin                string        `json:"origin,omitempty"`
	TrunkSid              string        `json:"trunk_sid,omitempty"`
	EmergencyStatus       string        `json:"emergency_status,omitempty"`
	EmergencyAddressSid   string        `json:"emergency_address_sid,omitempty"`
	AddressSid            string        `json:"address_sid,omitempty"`
	IdentitySid           string        `json:"identity_sid,omitempty"`
	Uri                   string        `json:"uri,omitempty"`
}

// GetIncomingPhoneNumbersLocal make GET request on twilio api
// endpoint: /Accounts/[AccountSid]/IncomingPhoneNumbers/Local
func (c *Client) GetIncomingPhoneNumbersLocal(queryValues url.Values) (*IncomingPhoneNumbers, *ResponseError) {

	var i IncomingPhoneNumbers
	if err := c.get(&i, fmt.Sprintf("%s/IncomingPhoneNumbers/Local.json", c.accountSid), queryValues); err != nil {
		return nil, err
	}
	return &i, nil

}

// CreateIncomingPhoneNumbersLocal make GET request on twilio api
// endpoint: /Accounts/[AccountSid]/IncomingPhoneNumbers/Local
func (c *Client) CreateIncomingPhoneNumbersLocal(formValues url.Values) (*IncomingPhoneNumber, *ResponseError) {

	var i IncomingPhoneNumber
	if err := c.post(&i, fmt.Sprintf("%s/IncomingPhoneNumbers/Local.json", c.accountSid), formValues); err != nil {
		return nil, err
	}
	return &i, nil

}
