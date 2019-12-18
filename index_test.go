package woocommerce

import (
	"testing"
)

func TestIndex(t *testing.T) {
	c := getTestClient()
	_, err := c.Index()
	if err != nil {
		t.Fatal(err)
	}
}
