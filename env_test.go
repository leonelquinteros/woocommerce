package woocommerce

import (
	"os"
)

func getTestClient() Client {
	cc := ClientConfig{
		APIHost:        os.Getenv("WC_API_HOST"),
		ConsumerKey:    os.Getenv("WC_API_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("WC_API_CONSUMER_SECRET"),
		Debug:          true,
	}
	return NewClient(cc)
}
