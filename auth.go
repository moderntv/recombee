package recombee

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"
)

type Auth interface {
	Authenticate(request *http.Request) error
}

type HmacSignatureAuth struct {
	secret []byte
}

type BearerTokenAuth struct {
	token string
}

// NewHmacSignatureAuth creates authentication using token created in Recombee admin for the HMAC signature.
func NewHmacSignatureAuth(secret []byte) *HmacSignatureAuth {
	return &HmacSignatureAuth{secret: secret}
}

func (a *HmacSignatureAuth) Authenticate(request *http.Request) error {
	u := request.URL
	q := u.Query()
	q.Set("hmac_timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	u.RawQuery = q.Encode()
	mac := hmac.New(sha1.New, a.secret)
	mac.Write([]byte(u.RequestURI()))
	sum := mac.Sum(nil)
	q.Set("hmac_sign", hex.EncodeToString(sum))
	u.RawQuery = q.Encode()
	return nil
}

func NewBearerTokenAuth(token string) *BearerTokenAuth {
	return &BearerTokenAuth{token: token}
}

func (a *BearerTokenAuth) Authenticate(request *http.Request) error {
	request.Header.Set("Authorization", "Bearer "+a.token)
	return nil
}
