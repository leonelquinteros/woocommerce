[![GitHub release](https://img.shields.io/github/release/leonelquinteros/woocommerce.svg)](https://github.com/leonelquinteros/woocommerce)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/leonelquinteros/woocommerce?status.svg)](https://godoc.org/github.com/leonelquinteros/woocommerce)
[![Go Report Card](https://goreportcard.com/badge/github.com/leonelquinteros/woocommerce)](https://goreportcard.com/report/github.com/leonelquinteros/woocommerce)


# Woocommerce SDK 

Woocommerce REST API Client for Go

## Quick start

```
go get github.com/leonelquinteros/woocommerce
```

```go
package main

import "github.com/leonelquinteros/woocommerce"

func main() {
    // Grab config from environment variables
    conf := woocommerce.ClientConfig{
		APIHost:        os.Getenv("WC_API_HOST"),
		ConsumerKey:    os.Getenv("WC_API_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("WC_API_CONSUMER_SECRET"),
    }
    
    // Get Woocommerce client
    client := woocommerce.NewClient(conf)
    
    // List orders
    params := url.Values{}
	params.Set("orderby", "id")
	params.Set("order", "desc")
    orders, err := client.Orders().List(params)
    if err != nil {
        log.Fatal(err)
    }

    // Print results
    for _, order := range orders {
        log.Printf("%+v", order)
    }
}
```


## Available endpoints

### Customers

- [List costumers](https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-customers)
- [Get costumer](https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-a-customer)

```go
client.Customers().List(params)
client.Customers().Get(customerID)
```

### Orders

- [List orders](https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-orders)
- [Get order](https://woocommerce.github.io/woocommerce-rest-api-docs/#retrieve-an-order)
- [List orders notes](https://woocommerce.github.io/woocommerce-rest-api-docs/#list-all-order-notes)

```go
client.Orders().List(params)
client.Orders().Get(orderID)
client.Orders().ListOrderNotes(orderID)
```


## Tests

Most tests will depend on several environment variables to use as configuration for the Woocommerce client. 

Check the helper function at `env_test.go`:

```go
func getTestClient() Client {
	cc := ClientConfig{
		APIHost:        os.Getenv("WC_API_HOST"),
		ConsumerKey:    os.Getenv("WC_API_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("WC_API_CONSUMER_SECRET"),
		Debug:          true,
	}
	return NewClient(cc)
}
```

You have to set `WC_API_HOST`, `WC_API_CONSUMER_KEY` and `WC_API_CONSUMER_SECRET` 
to the corresponding configuration values for your environment.