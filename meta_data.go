package woocommerce

// MetaData object
type MetaData struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// LinksResponse data
type LinksResponse struct {
	Self []struct {
		Href string `json:"href"`
	} `json:"self"`
	Collection []struct {
		Href string `json:"href"`
	} `json:"collection"`
}
