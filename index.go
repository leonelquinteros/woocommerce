package woocommerce

// IndexResponse data
type IndexResponse struct {
	Namespace string
	Routes    map[string]interface{}
}

// Index method returns all available API endpoints for the current host.
func (c Client) Index() (IndexResponse, error) {
	var ir IndexResponse
	err := c.Request("GET", "", nil, &ir)
	return ir, err
}
