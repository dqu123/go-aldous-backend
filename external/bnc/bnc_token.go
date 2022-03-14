package bnc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PostAccessTokenRequestBody struct {
	Audience  string `json:"audience"`
	ClientID  string `json:"client_id"`
	GrantType string `json:"grant_type"`
}

type PostAccessTokenResponseBody struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func postAccessToken() (*PostAccessTokenResponseBody, error) {
	// Get client ID from environment
	clientID := os.Getenv("BNC_CLIENT_ID")
	rapidAPIKey := os.Getenv("RAPID_API_KEY")
	requestBody := PostAccessTokenRequestBody{
		Audience:  "https://api.bravenewcoin.com",
		ClientID:  clientID,
		GrantType: "client_credentials",
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Get token
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://%s/oauth/token", BNCRapidAPIHost),
		bytes.NewBuffer(requestBodyBytes),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-rapidapi-host", BNCRapidAPIHost)
	req.Header.Set("x-rapidapi-key", rapidAPIKey)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// check for response error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// read response data
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// print request `Content-Type` header
	requestContentType := res.Request.Header.Get("Content-Type")
	fmt.Println("Request content-type:", requestContentType)

	responseBody := &PostAccessTokenResponseBody{}
	err = json.Unmarshal(resBytes, responseBody)
	if err != nil {
		return nil, err
	}

	// print response body
	fmt.Printf("%s\n", resBytes)
	return responseBody, nil
}
