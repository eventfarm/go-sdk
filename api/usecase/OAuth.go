package usecase

import (
	"net/http"
	"net/url"

	"bitbucket.ef.network/go/sdk"
)

type OAuth struct {
	restClient sdk.RestClientInterface
}

func NewOAuth(restClient sdk.RestClientInterface) *OAuth {
	return &OAuth{restClient}
}

// GET: Queries
// @param string OauthAccessTokenId

type GetOAuthAccessTokenParameters struct {
	OauthAccessTokenId string
}

func (t *OAuth) GetOAuthAccessToken(p *GetOAuthAccessTokenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`oauthAccessTokenId`, p.OauthAccessTokenId)

	return t.restClient.Get(
		`/v2/OAuth/UseCase/GetOAuthAccessToken`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string Email

type CreateGhostAccessTokenParameters struct {
	Email string
}

func (t *OAuth) CreateGhostAccessToken(p *CreateGhostAccessTokenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, p.Email)

	return t.restClient.Post(
		`/v2/OAuth/UseCase/CreateGhostAccessToken`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Identifier

type RevokeAccessTokenParameters struct {
	Identifier string
}

func (t *OAuth) RevokeAccessToken(p *RevokeAccessTokenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`identifier`, p.Identifier)

	return t.restClient.Post(
		`/v2/OAuth/UseCase/RevokeAccessToken`,
		&queryParameters,
		nil,
		nil,
	)
}
