package woocommerce

import "testing"

func TestGetOrders(t *testing.T) {
	c := getTestClient()
	r, err := c.Orders().List(nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Orders().Get(r[0].ID)
	if err != nil {
		t.Fatal(err)
	}
}
