package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

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

// OptionEndpoint overrides the endpoint for the http client
func OptionEndpoint(e string) func(c *Client) {
	return func(c *Client) {
		c.endpoint = e
	}
}

var httpDefaultClient = &http.Client{
	Timeout: 30 * time.Second,
}

// NewClient Returns a New Sozluk Client
func NewClient(options ...Option) *Client {
	c := &Client{
		httpClient: httpDefaultClient,
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

func (c *Client) get(ctx context.Context, path string, values url.Values, response interface{}) error {
	araError := &AraError{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoint+path, nil)

	if err != nil {
		return err
	}

	req.URL.RawQuery = values.Encode()

	c.Debugf("outbound request to %v", c.endpoint+path)
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("non 200 response code")
	}

	defer resp.Body.Close()

	respTxt, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(respTxt, araError)

	// We need to successfully unmarshal into the error struct to determine
	// if there was an issue, such as no results returned, as the sozluk
	// always returns a response code of 200 if the request was not malformed
	if err == nil {
		return errors.New(araError.Err)
	}

	err = json.Unmarshal(respTxt, response)

	if err != nil {
		return err
	}

	return nil
}
