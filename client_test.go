package woocommerce

import (
	"testing"
)

func TestClientConfig(t *testing.T) {
	check := "https://example.com/wp-json/wc/v3/products"
	cc := ClientConfig{
		APIHost: "https://example.com/wp-json/wc/v3",
	}

	r := cc.GetAPIEndpoint("products")
	if r != check {
		t.Errorf("Expected '%s', got '%s'", check, r)
	}
}
