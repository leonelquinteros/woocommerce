package woocommerce

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

// ClientConfig is used to instantiate a new Woocommerce API Client with the defined values.
// To construct requests to https://example.com/wp-json/wc/v3/products set the client configuration as follows:
//
//   ClientConfig {
//	   APIHost:    "https://example.com",
//	   APIPrefix:  "/wp-json/wc",
//	   APIVersion: "/v3",
//   }
//
type ClientConfig struct {
	APIHost        string
	APIPrefix      string
	APIVersion     string
	ConsumerKey    string
	ConsumerSecret string
	UseBasicAuth   bool
	Debug          bool
}

// GetAPIEndpoint takes an API method name and returns the full URL for that endpoint based on the current configuration.
func (cc ClientConfig) GetAPIEndpoint(method string) string {
	base, err := url.Parse(cc.APIHost)
	if err != nil {
		return ""
	}
	base.Path = path.Join(base.Path, cc.APIPrefix, cc.APIVersion, method)
	return base.String()
}

// Client is the main SDK entry point for all methods.
// It wraps all available SDK objects into a single access point.
type Client struct {
	config ClientConfig
}

// NewClient constructs a new woocommerce.Client object and sets the given configuration.
func NewClient(cc ClientConfig) Client {
	return Client{
		config: cc,
	}
}

// Request performs a request and unmarshals JSON response into given response object.
func (c Client) Request(method, endpoint string, params url.Values, response interface{}) error {
	// Get full endpoint URL
	uri := c.config.GetAPIEndpoint(endpoint)

	// Parse params
	if params == nil {
		params = url.Values{}
	}

	// Set Auth
	if !c.config.UseBasicAuth {
		params.Set("consumer_key", c.config.ConsumerKey)
		params.Set("consumer_secret", c.config.ConsumerSecret)
	}

	// Add params to URI
	encodedParams := params.Encode()
	if encodedParams != "" {
		uri += "?" + params.Encode()
	}

	// Create request
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return err
	}

	// Set Auth
	if c.config.UseBasicAuth {
		req.SetBasicAuth(c.config.ConsumerKey, c.config.ConsumerSecret)
	}

	// Debug
	if c.config.Debug {
		log.Printf("NEW REQUEST TO %s", uri)
	}

	// Perform request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Debug
	if c.config.Debug {
		log.Printf("RESPONSE FROM %s: \n%s", uri, body)
	}

	// Handle API errors
	if resp.StatusCode >= 400 {
		errResp := ErrorResponse{}
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return err
		}

		err = errors.New(errResp.Code + ": " + errResp.Message)
	} else {
		err = json.Unmarshal(body, response)
	}
	return err
}

// Customers returns a Customers API client
func (c Client) Customers() Customers {
	return Customers{
		Client: c,
	}
}

// Orders returns a Orders API client
func (c Client) Orders() Orders {
	return Orders{
		Client: c,
	}
}
