# Recombee
[Recombee API](https://docs.recombee.com/api.html) client implements the recombee API specification in go.

* The client is supposed to be used using batch requests using `Request(...)` method. 
* The batch requests are efficient and saves outgoing traffic. However, you can still use the `Request(...)` method per each request. 
* It will make a single call with single response but using `/batch` endpoint.

# Getting started
As recombee serves recommendations based on the user interactions, first of all there have to be inserted items and interactions. Afterwards the user asks for his/hers recommendation. 
To import recombee API client insert

```import "github.com/moderntv/recombee"```

to your code and run `go [build|run|test]` which automatically fetch dependency.

## Example
```
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/moderntv/recombee"
)

type Recombee struct {
	client *recombee.Client
}

func main() {
	databaseId := "dev"
	token := "yourtoken"
	r := &Recombee{
		client: recombee.NewClient("https://rapi.recombee.com", databaseId, token),
	}
	var seriesID int64 = 123
	var userID int64 = 1
	myRequests := []recombee.Request{
		recombee.AddItem("123"),
		recombee.AddRating(
			strconv.FormatInt(userID, 10),
			strconv.FormatInt(seriesID, 10),
			float64(0.8),
			recombee.WithCascadeCreate(),
		),
	}

	responses, err := r.client.Request(context.Background(), myRequests...)
	if err != nil {
		panic(err)
	}

	for _, resp := range responses {
		if resp.Code != http.StatusCreated && resp.Code != http.StatusConflict {
			log.Print(resp.Code)
			log.Print(resp.Json)
		}
	}

	// Ask for single recommendation
	resp, err := r.client.Request(context.Background(), recombee.RecommendItemsToUser(
		strconv.FormatInt(userID, 10),
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

	fmt.Println(recommendations)
}
```

