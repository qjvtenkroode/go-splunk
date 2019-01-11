package splunk

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c Client) Post(url string, data *url.Values) (string, error) {
	return c.Do(url, "POST", data)
}

func (c Client) Get(url string, data *url.Values) (string, error) {
	return c.Do(url, "GET", data)
}

func (c Client) Do(url string, method string, data *url.Values) (string, error) {
	client := httpClient()

	var payload io.Reader
	if data != nil {
		payload = bytes.NewBufferString(data.Encode())
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return "", err
	}
	c.addHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return string(body), err
}

func (c Client) addHeader(req *http.Request) {
	if c.sessionKey.Value != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Splunk %s", c.sessionKey.Value))
	} else {
		req.SetBasicAuth(c.Username, c.Password)
	}
}

func httpClient() *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: transport}
	return client
}
