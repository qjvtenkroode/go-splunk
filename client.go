package splunk

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Client represents the a client through which the Splunk API is used
type Client struct {
	Username   string
	Password   string
	Url        string
	sessionKey SessionKey
}

// SessionKey represents the JSON object returned after a succesful login
type SessionKey struct {
	Value string `json:"sessionKey"`
}

// Login authenticates the user and returns a SessionKey
func (c Client) Login() (SessionKey, error) {
	data := make(url.Values)
	data.Add("username", c.Username)
	data.Add("password", c.Password)
	data.Add("output_mode", "json")

	resp, err := c.Post(fmt.Sprintf("%s/services/auth/login", c.Url), &data)
	if err != nil {
		return SessionKey{}, err
	}

	bytes := []byte(resp)
	var sessionKey SessionKey
	err = json.Unmarshal(bytes, &sessionKey)
	c.sessionKey = sessionKey
	return sessionKey, err
}
