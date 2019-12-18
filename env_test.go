package woocommerce

import (
	"os"
)

func getTestClient() Client {
	cc := ClientConfig{
		APIHost:        os.Getenv("WC_API_HOST"),
		APIPrefix:      os.Getenv("WC_API_PREFIX"),
		APIVersion:     os.Getenv("WC_API_VERSION"),
		ConsumerKey:    os.Getenv("WC_API_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("WC_API_CONSUMER_SECRET"),
		Debug:          true,
	}
	return NewClient(cc)
}
