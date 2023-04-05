# Recombee

A [Recombee API](https://docs.recombee.com/api.html) client written in Go. 

[![GoDoc][GoDoc-Image]][GoDoc-Url]

[GoDoc-Url]: https://pkg.go.dev/github.com/nats-io/nats.go
[GoDoc-Image]: https://img.shields.io/badge/GoDoc-reference-007d9c

## Features

Is capable of managing:

* Items and Item Properties
* Users
* User-Item Interactions
* Recommendations
* Series

Uses [batch mode](https://docs.recombee.com/api.html#batch) by default to send requests to the API.

## Basic usage

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	
	"github.com/moderntv/recombee"
)

const (
	apiUri     = "https://rapi.recombee.com"
	databaseId = "dev"
	token      = "api-token"
)

func main() {
	userID := "23"
	seriesID := "6821"
	rating := 0.8
	myRequests := []recombee.Request{
		recombee.AddItem("1215"),
		recombee.AddRating(
			userID,
			seriesID,
			rating,
			recombee.WithCascadeCreate(),
		),
	}

	client := recombee.NewClient(apiUri, databaseId, token)
	responses, err := client.Request(context.Background(), myRequests...)
	if err != nil {
		panic(err)
	}

	for _, resp := range responses {
		if resp.Code != http.StatusCreated && resp.Code != http.StatusConflict {
			fmt.Printf("%s: %s\n", http.StatusText(resp.Code), resp.Json)
		}
	}

	// Ask for single recommendation
	resp, err := client.Request(context.Background(), recombee.RecommendItemsToUser(
		userID,
		1,
	))
	if err != nil {
		panic(err)
	}

	var recommendations recombee.Recommendations
	err = json.Unmarshal(resp[0].Json, &recommendations)
	if err != nil {
		return
	}

	fmt.Printf("%#v\n", recommendations)
}
```

## License

Unless otherwise noted, the source files are distributed under the MIT license found in the LICENSE file.
