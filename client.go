package terraspoof

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// HostURL - Default MajesticCloud URL
const HostURL string = "http://localhost:8080/api/v1"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	//Token      string
	//Auth       AuthStruct
}

//// AuthStruct -
//type AuthStruct struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//}
//
//// AuthResponse -
//type AuthResponse struct {
//	UserID   int    `json:"user_id`
//	Username string `json:"username`
//	Token    string `json:"token"`
//}

func NewClient(host *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If username or password not provided, return empty client
	//if username == nil || password == nil {
	//	return &c, nil
	//}

	//c.Auth = AuthStruct{
	//	Username: *username,
	//	Password: *password,
	//}
	//
	//ar, err := c.SignIn()
	//if err != nil {
	//	return nil, err
	//}
	//
	//c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	//token := c.Token

	//if authToken != nil {
	//	token = *authToken
	//}
	//
	//req.Header.Set("Authorization", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
