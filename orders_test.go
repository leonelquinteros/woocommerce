package woocommerce

import (
	"net/url"
	"testing"
)

func TestGetOrders(t *testing.T) {
	c := getTestClient()

	params := url.Values{}
	params.Set("orderby", "id")
	params.Set("order", "desc")

	r, err := c.Orders().List(params)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Orders().Get(r[0].ID)
	if err != nil {
		t.Fatal(err)
	}
}
