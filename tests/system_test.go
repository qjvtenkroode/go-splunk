package tests

import (
	"fmt"
	"testing"
	"time"

	splunk "github.com/qjvtenkroode/go-splunk"
)

var serviceWaitTime = 30 * time.Second

func TestSystemRestart(t *testing.T) {
	c := splunk.Client{Username: "admin", Password: "12345678", Url: "https://localhost:8089"}
	sessionKey, err := c.Login()
	if err != nil {
		t.Errorf("Error creating Client: %s", err)
	}
	if sessionKey.Value == "" {
		t.Error("SessionKey is empty!")
	}

	resp, err := c.Restart()
	if err != nil {
		t.Errorf("Error restarting server: %s", err)
	}
	time.Sleep(serviceWaitTime)
}

func TestSystemRestartWeb(t *testing.T) {
	c := splunk.Client{Username: "admin", Password: "12345678", Url: "https://localhost:8089"}
	sessionKey, err := c.Login()
	if err != nil {
		t.Errorf("Error creating Client: %s", err)
	}
	if sessionKey.Value == "" {
		t.Error("SessionKey is empty!")
	}

	resp, err := c.RestartWeb()
	if err != nil {
		t.Errorf("Error restarting webui: %s", err)
	}
	time.Sleep(serviceWaitTime)
}
