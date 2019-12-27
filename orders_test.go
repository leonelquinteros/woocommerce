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

	if len(r) > 0 {
		_, err = c.Orders().Get(r[0].ID)
		if err != nil {
			t.Error(err)
		}

		ons, err := c.Orders().ListOrderNotes(r[0].ID)
		if err != nil {
			t.Error(err)
		}
		if len(ons) > 0 {
			_, err = c.Orders().GetOrderNote(r[0].ID, ons[0].ID)
		}
	}
}
