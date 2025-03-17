package ton

import "net/http"

type Client struct {
	tonScanURL string

	client *http.Client
}

func New(tonScanURL string) *Client {
	return &Client{
		tonScanURL: tonScanURL,
		client:     http.DefaultClient,
	}
}
