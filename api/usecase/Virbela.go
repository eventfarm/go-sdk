/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type Virbela struct {
	restClient rest.RestClientInterface
}

func NewVirbela(restClient rest.RestClientInterface) *Virbela {
	return &Virbela{restClient}
}

// GET: Queries

type CreateUserForVirbelaParameters struct {
	InvitationId string
	Password     string
}

func (t *Virbela) CreateUserForVirbela(p *CreateUserForVirbelaParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`password`, p.Password)

	return t.restClient.Get(
		`/v2/Virbela/UseCase/CreateUserForVirbela`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetUserStatusForVirbelaParameters struct {
	InvitationId string
}

func (t *Virbela) GetUserStatusForVirbela(p *GetUserStatusForVirbelaParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Get(
		`/v2/Virbela/UseCase/GetUserStatusForVirbela`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
