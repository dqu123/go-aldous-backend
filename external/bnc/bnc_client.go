package bnc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	BNCRapidAPIHost = "bravenewcoin.p.rapidapi.com"
)

type GetAssetsResponseBody struct {
	Content []BncGetAssetsContent `json:"content"`
}
type BncGetAssetsContent struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	SlugName string `json:"slugName"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Url      string `json:"url"`
}

func GetAssets() (*GetAssetsResponseBody, error) {
	rapidAPIKey := os.Getenv("RAPID_API_KEY")
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://%s/asset", BNCRapidAPIHost),
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-rapidapi-host", BNCRapidAPIHost)
	req.Header.Set("x-rapidapi-key", rapidAPIKey)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	switch res.StatusCode {
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("issue accessing %s", BNCRapidAPIHost)
	}
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resBody := &GetAssetsResponseBody{}
	err = json.Unmarshal(resBytes, resBody)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}

type GetAssetResponseBody struct {
	Content []BncGetAssetContent `json:"content"`
}
type BncGetAssetContent struct {
	ID            string `json:"id"`
	AssetID       string `json:"assetId"`
	Timestamp     string `json:"timestamp"`
	MarketCapRank string `json:"marketCapRank"`
}

func GetAsset() (*GetAssetResponseBody, error) {
	postAccessTokenResponse, err := postAccessToken()
	if err != nil {
		return nil, err
	}
	token := postAccessTokenResponse.AccessToken
	fmt.Println("Token", token)
	// TODO: pass Bearer Token
	return nil, nil
}
