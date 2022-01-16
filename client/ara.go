package client

import (
	"net/url"
)

const (
	path   = "/gts"
	araKey = "ara"
)

func (c *Client) Ara(s string) ([]Kelime, error) {
	v := make(url.Values)
	k := &[]Kelime{}

	v.Set(araKey, s)

	err := c.Get(path, v, k)

	if err != nil {
		return *k, err
	}

	return *k, nil
}
