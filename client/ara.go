package client

import (
	"fmt"
	"net/url"
)

const (
	path   = "/gts"
	araKey = "ara"
)

// AraError represents an error returned from the Sozluk Ara method
type AraError struct {
	// Err represents the instance of the error returned,
	// e.g: word not found
	Err string `json:"error"`
}

// Ara returns a slice of Kelimelar/Kelimes for a given search
func (c *Client) Ara(s string) ([]Kelime, error) {
	v := make(url.Values)
	k := &[]Kelime{}

	v.Set(araKey, s)

	err := c.get(path, v, k)

	if err != nil {
		return *k, fmt.Errorf("failed to find word: %w", err)
	}

	return *k, nil
}
