package rest

import (
	"net/http"
	"net/url"
	"time"
    "encoding/base64"

)

const version = `7.2.4`

var authAccessToken *OAuthAccessToken
var tokenExpiresAt time.Time

type EventFarmRestConnection struct {
	LoginUrl string
	ApiUrl   string
}

type EventFarmRestClient struct {
	restClient            RestClientInterface
	accessTokenRestClient RestClientInterface
	clientId              string
	clientSecret          string
	audience              string
	oAuthAccessToken      *OAuthAccessToken
}

func basicAuth(username, password string) string {
  auth := username + ":" + password
  return base64.StdEncoding.EncodeToString([]byte(auth))
}


func NewEventFarmRestClient(
	restClient RestClientInterface,
	accessTokenRestClient RestClientInterface,
	clientId string,
	clientSecret string,
	audience string,
	oAuthAccessToken *OAuthAccessToken,
) *EventFarmRestClient {

	return &EventFarmRestClient{
		restClient:            restClient,
		accessTokenRestClient: accessTokenRestClient,
		clientId:              clientId,
		clientSecret:          clientSecret,
		audience:              audience,
		oAuthAccessToken:      oAuthAccessToken,
	}
}

func (t *EventFarmRestClient) Get(
	url string,
	queryParameters *url.Values,
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	newHeaders, err := t.getAuthorizationHeader(headers)
	if err != nil {
		return
	}

	return t.restClient.Get(
		url,
		queryParameters,
		newHeaders,
		timeout,
	)
}

func (t *EventFarmRestClient) Post(
	url string,
	formParameters *url.Values,
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	newHeaders, err := t.getAuthorizationHeader(headers)
	if err != nil {
		return
	}

	return t.restClient.Post(
		url,
		formParameters,
		newHeaders,
		timeout,
	)
}

func (t *EventFarmRestClient) PostJSON(
	url string,
	data *map[string]interface{},
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	newHeaders, err := t.getAuthorizationHeader(headers)
	if err != nil {
		return
	}

	return t.restClient.PostJSON(
		url,
		data,
		newHeaders,
		timeout,
	)
}

func (t *EventFarmRestClient) PostMultipart(
	url string,
	multipart map[string]string,
	headers map[string]string,
	timeout *time.Duration,
) (resp *http.Response, err error) {

	newHeaders, err := t.getAuthorizationHeader(headers)
	if err != nil {
		return
	}

	return t.restClient.PostMultipart(
		url,
		multipart,
		newHeaders,
		timeout,
	)
}

func (t *EventFarmRestClient) getAuthorizationHeader(headers map[string]string) (newHeaders map[string]string, err error) {
	newHeaders = make(map[string]string)

	for k, v := range headers {
		newHeaders[k] = v
	}

	oAuthAccessToken, err := t.getOAuthAccessToken()

	if err != nil {
		return
	}

	newHeaders[`Authorization`] = oAuthAccessToken.getHeaderString()

	return
}

func (t *EventFarmRestClient) getOAuthAccessToken() (oAuthAccessToken *OAuthAccessToken, err error) {
	if authAccessToken == nil {
		t.oAuthAccessToken, err = t.getAccessTokenWithClientCredentialsGrant()
		authAccessToken = t.oAuthAccessToken
		tokenExpiresAt = time.Now().Add(time.Second * time.Duration(authAccessToken.expiresIn-2))

		if err != nil {
			return
		}
	} else {
		t.oAuthAccessToken = authAccessToken
	}

	if time.Now().After(tokenExpiresAt) {
		t.oAuthAccessToken, err = t.getAccessTokenWithClientCredentialsGrant()
		authAccessToken = t.oAuthAccessToken
		tokenExpiresAt = time.Now().Add(time.Second * time.Duration(authAccessToken.expiresIn-2))
		if err != nil {
			return
		}
	}

	oAuthAccessToken = t.oAuthAccessToken
	return oAuthAccessToken, err
}

func (t *EventFarmRestClient) getAccessTokenWithClientCredentialsGrant() (oAuthAccessToken *OAuthAccessToken, err error) {
	headers := map[string]string{
    		"Authorization":    "Basic "+ basicAuth(t.clientId, t.clientSecret),
    }
	resp, err := t.accessTokenRestClient.PostJSON(
		"/oauth2/token?grant_type=client_credentials&scope=target-entity:"+t.audience+":read,write",
		nil,
		headers,
		nil,
	)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	oAuthAccessToken, err = NewOAuthAccessTokenFromResponse(resp)
	if err != nil {
		return
	}

	return
}
