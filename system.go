package splunk

import (
	"fmt"
	"net/url"
)

func (c Client) Restart() (string, error) {
	data := make(url.Values)
	data.Add("output_mode", "json")

	resp, err := c.Post(fmt.Sprintf("%s/services/server/control/restart", c.Url), &data)
	return resp, err
}

func (c Client) RestartWeb() (string, error) {
	data := make(url.Values)
	data.Add("output_mode", "json")

	resp, err := c.Post(fmt.Sprintf("%s/services/server/control/restart_webui", c.Url), &data)
	return resp, err
}
