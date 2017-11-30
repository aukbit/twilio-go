package twilio

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	// DefaultName default pakage name
	DefaultName = "twilio-go"
	// DefaultEndpoint default twilio endpoint
	DefaultEndpoint = "https://api.twilio.com/2010-04-01/Accounts/"
)

// A Client defines parameters for making calls to the twilio REST API
type Client struct {
	accountSid, authToken string
	httpClient            *http.Client
}

// ResponseError is a representation of a twilio response error.
type ResponseError struct {
	Status   int    `json:"status"`    // HTTP specific error code
	Message  string `json:"message"`   // HTTP error message
	Code     int    `json:"code"`      // Twilio specific error code
	MoreInfo string `json:"more_info"` // Additional info from Twilio
}

// New creates a new http twilio client
func New(cfgs ...Config) *Client {
	return newClient(cfgs...)
}

// newClient will instantiate a new client with the given options
func newClient(cfgs ...Config) *Client {
	c := &Client{httpClient: http.DefaultClient}
	c.apply(cfgs...)
	return c
}

func (c *Client) apply(cfgs ...Config) {
	for _, cfg := range cfgs {
		cfg.apply(c)
	}
}

func (c *Client) get(out interface{}, url string, queryValues url.Values) *ResponseError {
	var _url string
	_url = DefaultEndpoint + url
	if queryValues != nil {
		_url += "?" + queryValues.Encode()
	}

	req, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		return &ResponseError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	req.SetBasicAuth(c.accountSid, c.authToken)

	return c.do(out, req)
}

func (c *Client) post(out interface{}, url string, formValues url.Values) *ResponseError {
	req, err := http.NewRequest("POST", DefaultEndpoint+url, strings.NewReader(formValues.Encode()))
	if err != nil {
		return &ResponseError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	req.SetBasicAuth(c.accountSid, c.authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return c.do(out, req)
}

func (c *Client) do(out interface{}, r *http.Request) *ResponseError {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return &ResponseError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var e ResponseError
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return &ResponseError{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		return &e
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return &ResponseError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}
