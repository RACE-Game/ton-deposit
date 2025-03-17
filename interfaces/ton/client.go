package ton

type Client struct {
	tonScanURL string
}

func New(tonScanURL string) *Client {
	return &Client{
		tonScanURL: tonScanURL,
	}
}
