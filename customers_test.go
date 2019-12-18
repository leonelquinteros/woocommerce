package woocommerce

import (
	"net/url"
	"testing"
)

func TestGetCustomers(t *testing.T) {
	c := getTestClient()

	params := url.Values{}
	params.Set("orderby", "id")
	params.Set("order", "desc")

	r, err := c.Customers().List(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(r) > 0 {
		_, err = c.Customers().Get(r[0].ID)
		if err != nil {
			t.Fatal(err)
		}
	}
}
