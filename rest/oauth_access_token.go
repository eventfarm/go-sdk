package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type OAuthAccessToken struct {
	tokenType    string
	accessToken  string
	expiresIn    int64
	refreshToken string
}

type OAuthAccessTokenResponse struct {
	TokenType       string `json:"token_type"`
	ExpiresIn       int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
}

type EventFarmAccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType string  `json:"token_type"`
	ExpiresIn int64  `json:"expires_in"`
}

func NewOAuthAccessTokenFromResponse(r *http.Response) (oAuthAccessToken *OAuthAccessToken, err error) {
	if r.StatusCode != 200 {
		err = errors.New(`Unable to load access token`)
		LogResponse(r)
		return
	}

	bodyBuff, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	eventFarmAccessTokenResponse := EventFarmAccessTokenResponse{}
	err = json.Unmarshal(bodyBuff, &eventFarmAccessTokenResponse)
	if err != nil {
		return
	}

	if err != nil {
		return
	}

	oAuthAccessToken = &OAuthAccessToken{
		tokenType:    eventFarmAccessTokenResponse.TokenType,
		expiresIn:    eventFarmAccessTokenResponse.ExpiresIn,
		accessToken:  eventFarmAccessTokenResponse.AccessToken,
	}

	return
}

type OAuthToken struct {
	token     string
	userId    string
	expiresIn int64
}

func (t OAuthAccessToken) getHeaderString() string {
	return t.tokenType + ` ` + t.accessToken
}

func (t *OAuthAccessToken) isExpired() bool {
	return time.Now().Unix() > t.expiresIn
}
