# Recombee

A [Recombee API](https://docs.recombee.com/api.html) client written in Go. 

[![GoDoc][GoDoc-Image]][GoDoc-Url]

[GoDoc-Url]: https://pkg.go.dev/github.com/moderntv/recombee
[GoDoc-Image]: https://img.shields.io/badge/GoDoc-reference-007d9c

## Features

Is capable of managing:

* Items
    * Add item
    * Delete item
    * List items
    * Delete more items
* Item Properties
    * Add item property
    * Delete item property
    * Get item property info
    * List item properties
    * Set item values
    * Get item values
    * Update more items
* Users
    * Add user
    * Delete user
    * Merge users
    * List users
* User Properties
    * Add user property
    * Delete user property
    * Get user property info
    * List user properties
    * Set user values
    * Get user values
* User-Item Interactions
    * Add detail view
    * Delete detail view
    * List item detail views
    * List user detail views
    * Add purchase
    * Delete purchase
    * List item purchases
    * List user purchases
    * Add rating
    * Delete rating
    * List item ratings
    * List user ratings
    * Add cart addition
    * Delete cart addition
    * List item cart additions
    * List user cart additions
    * Add bookmark
    * Delete bookmark
    * List item bookmarks
    * List user bookmarks
    * Set view portion
    * Delete view portion
    * List item view portions
    * List user view portions
* Recommendations
    * Recommending items
    * Recommending item segments
    * Recommending users 
* Search
    * Search items 
    * Search items segments
    * Synonyms
* Series
    * Add series
    * Delete series
    * List series
    * List series items
    * Insert to series
    * Remove from series
* Segmentations definition
    * General
    * Property based
    * Manual ReQL
    * Auto ReQL
* Miscellaneous
    * Reset database
    * Batch

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
