package twilio

// Config is used to set configuration options for the client.
type Config interface {
	apply(*Client)
}

// configFunc wraps a func so it satisfies the Config interface.
type configFunc func(*Client)

func (f configFunc) apply(s *Client) {
	f(s)
}

// AccountSid twilio account sid
func AccountSid(accountSid string) Config {
	return configFunc(func(c *Client) {
		c.accountSid = accountSid
	})
}

// AuthToken twilio authorization token
func AuthToken(authToken string) Config {
	return configFunc(func(c *Client) {
		c.authToken = authToken
	})
}
