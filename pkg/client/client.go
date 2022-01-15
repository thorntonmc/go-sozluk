package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

const (
	sozlukEndpoint = "https://sozluk.gov.tr"
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client for the Sozluk API
type Client struct {
	httpClient
	log.Logger
	debug    bool
	endpoint string
}

// Option defines an option for a client
type Option func(c *Client)

// OptionDebug enables debugging for the http client
func OptionDebug(b bool) func(c *Client) {
	return func(c *Client) {
		c.debug = b
	}
}

// NewClient Returns a New Sozluk Client
func NewClient(options ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{},
		Logger:     *log.New(),
		endpoint:   sozlukEndpoint,
	}

	for _, opt := range options {
		opt(c)
	}

	return c
}

//Debugf prints a formatted debug line
func (c *Client) Debugf(format string, v ...interface{}) {
	if c.debug {
		c.Logger.Debugf(fmt.Sprintf(format, v...))
	}
}

func (c *Client) Get(path string, values url.Values, response interface{}) error {
	req, err := http.NewRequest(http.MethodGet, c.endpoint+path, nil)

	if err != nil {
		return err
	}

	req.URL.RawQuery = values.Encode()

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respTxt, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(respTxt, response)

	if err != nil {
		return err
	}

	return nil
}
