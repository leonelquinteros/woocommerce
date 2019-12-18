package woocommerce

import (
	"testing"
)

func TestClientConfig(t *testing.T) {
	check := "https://example.com/wp-json/wc/v3/products"
	cc := ClientConfig{
		APIHost:    "https://example.com",
		APIPrefix:  "wp-json/wc",
		APIVersion: "v3",
	}

	r := cc.GetAPIEndpoint("products")
	if r != check {
		t.Errorf("Expected '%s', got '%s'", check, r)
	}

	// Add slashes
	cc.APIHost = "https://example.com/"
	cc.APIPrefix = "/wp-json/wc"
	cc.APIVersion = "/v3"

	r = cc.GetAPIEndpoint("products")
	if r != check {
		t.Errorf("Expected '%s', got '%s'", check, r)
	}
}
