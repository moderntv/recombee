// The recombee package implements recombee API specification.
// Does not include external packages so it should be easy to install.

package recombee

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Used for json slice encode.
type BatchRequest struct {
	Requests []Request `json:"requests"`
}

// Response from Recombee API to be checked.
type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// Returns all responses from batch API call.
type BatchResponse []struct {
	Code int `json:"code"`
	// Typically unmarshalled into [Recommendations] struct
	Json json.RawMessage `json:"json"`
}

// Client used for communication with Recombee API
type Client struct {
	// URI where the recombee API is exposed.
	baseURI string
	// DatabaseID points on storage where are API calls stores into or retrieves from.
	databaseID string
	// Authentication token created in Recombee admin
	token []byte

	// configurable via options
	requestTimeout  time.Duration
	maxBatchSize    int
	distinctRecomms bool //TODO implement

	httpClient *http.Client
}

// Returns new recombee API client.
func NewClient(baseURI string, databaseID string, token string, opts ...ClientOption) (c *Client) {
	c = &Client{
		baseURI:    baseURI,
		databaseID: databaseID,
		token:      []byte(token),

		requestTimeout: 5 * time.Minute,
		maxBatchSize:   10000, // API limit
	}
	for _, o := range opts {
		o(c)
	}
	c.httpClient = &http.Client{
		Timeout: c.requestTimeout,
	}

	return
}

func (c *Client) signURL(u *url.URL) (err error) {
	q := u.Query()

	// add timestamp
	q.Set("hmac_timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	u.RawQuery = q.Encode()

	// create signature
	var sum []byte
	mac := hmac.New(sha1.New, []byte(c.token))
	mac.Write([]byte(u.RequestURI()))
	sum = mac.Sum(nil)

	// add signature
	q.Set("hmac_sign", hex.EncodeToString(sum))
	u.RawQuery = q.Encode()

	return
}

func (c *Client) batchRequest(ctx context.Context, requests ...Request) (batchResponse BatchResponse, err error) {
	var batchRequest = BatchRequest{
		Requests: requests,
	}

	// URL
	var u *url.URL
	u, err = url.Parse(fmt.Sprintf("%s/%s/batch/", c.baseURI, c.databaseID))
	if err != nil {
		return
	}
	err = c.signURL(u)
	if err != nil {
		return
	}

	// request
	var body = bytes.NewBuffer(nil)
	err = json.NewEncoder(body).Encode(batchRequest)
	if err != nil {
		return
	}
	var request *http.Request
	request, err = http.NewRequestWithContext(ctx, http.MethodPost, u.String(), body)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	// response
	var response *http.Response
	response, err = c.httpClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		var resp Response
		err = json.NewDecoder(response.Body).Decode(&resp)
		if err != nil {
			return
		}
		err = fmt.Errorf("server returned '%s': %s", http.StatusText(response.StatusCode), resp.Message)
		return
	}

	err = json.NewDecoder(response.Body).Decode(&batchResponse)
	if err != nil {
		return
	}

	return
}

// Requests appropriate entity that is given by requests.
func (c *Client) Request(ctx context.Context, requests ...Request) (responses BatchResponse, err error) {
	if len(requests) == 0 {
		return
	}

	responses = make(BatchResponse, 0, len(requests))
	for i := 0; i < len(requests); i += c.maxBatchSize {
		var batch []Request
		size := min(len(requests)-i, c.maxBatchSize)
		batch = requests[i : i+size]

		var br BatchResponse
		br, err = c.batchRequest(ctx, batch...)
		if err != nil {
			return
		}

		responses = append(responses, br...)
	}

	return
}
