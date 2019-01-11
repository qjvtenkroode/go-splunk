package tests

import (
	"fmt"
	"testing"

	splunk "github.com/qjvtenkroode/go-splunk"
)

func TestClient(t *testing.T) {
	c := splunk.Client{Username: "admin", Password: "12345678", Url: "https://localhost:8089"}
	sessionKey, err := c.Login()
	if err != nil {
		t.Errorf("Error creating Client: %s", err)
	}
	fmt.Println(sessionKey)
	if sessionKey.Value == "" {
		t.Error("SessionKey is empty!")
	}
}
