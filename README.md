[![GitHub release](https://img.shields.io/github/release/leonelquinteros/woocommerce.svg)](https://github.com/leonelquinteros/woocommerce)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/leonelquinteros/woocommerce?status.svg)](https://godoc.org/github.com/leonelquinteros/woocommerce)
[![Build Status](https://travis-ci.org/leonelquinteros/woocommerce.svg?branch=master)](https://travis-ci.org/leonelquinteros/woocommerce)
[![Go Report Card](https://goreportcard.com/badge/github.com/leonelquinteros/woocommerce)](https://goreportcard.com/report/github.com/leonelquinteros/gotext)


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
		APIPrefix:      os.Getenv("WC_API_PREFIX"),
		APIVersion:     os.Getenv("WC_API_VERSION"),
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