package gapiauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Auth represents Google API authentication object.
type Auth struct {
	client          *http.Client
	oauth           *oauth2.Config
	tokenPath       string
	credentialsPath string
	scopes          []string
}

// NewAuth creates an `Auth` object.
func NewAuth(credentialsPath, tokenPath string, scopes []string) *Auth {
	return &Auth{
		tokenPath:       tokenPath,
		credentialsPath: credentialsPath,
		scopes:          scopes,
	}
}

// GetClient gets an HTTP client that can be used to interact with the Google
// API.
//
// If the HTTP client is unset, it is built using the OAuth 2 config.
func (a *Auth) GetClient() (*http.Client, error) {
	if a.client == nil {
		client, err := getClient(a.credentialsPath, a.tokenPath, a.scopes)
		if err != nil {
			return nil, fmt.Errorf("Error: Could not get client: %v", err)
		}
		a.client = client
	}
	return a.client, nil
}

func getClient(credentialsPath, tokenPath string, scopes []string) (*http.Client, error) {
	buf, err := ioutil.ReadFile(credentialsPath)
	if err != nil {
		return nil, err
	}

	cfg, err := google.ConfigFromJSON(buf, scopes...)
	if err != nil {
		return nil, fmt.Errorf("Error: Could not read config from %s: %v", credentialsPath, err)
	}

	tok, err := tokenFromFile(tokenPath)
	if err != nil {
		tok, err = getTokenFromWeb(cfg)
		if err != nil {
			return nil, fmt.Errorf("Error: Could not get token from web: %v", err)
		}
		saveToken(tokenPath, tok)
	}
	return cfg.Client(context.Background(), tok), nil
}

func tokenFromFile(tokenPath string) (*oauth2.Token, error) {
	fh, err := os.Open(tokenPath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(fh).Decode(tok)

	return tok, err
}

func getTokenFromWeb(cfg *oauth2.Config) (*oauth2.Token, error) {
	url := cfg.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser, then type the authorisation code:\n%v\n", url)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("Error: Unable to retrieve authentication code from console: %v", err)
	}

	tok, err := cfg.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, fmt.Errorf("Error: Unable to retrive authentication token from web: %v", err)
	}

	return tok, nil
}

func saveToken(tokenPath string, token *oauth2.Token) error {
	fh, err := os.OpenFile(tokenPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer fh.Close()

	return json.NewEncoder(fh).Encode(token)
}
