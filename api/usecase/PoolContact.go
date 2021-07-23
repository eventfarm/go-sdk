/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type PoolContact struct {
	restClient rest.RestClientInterface
}

func NewPoolContact(restClient rest.RestClientInterface) *PoolContact {
	return &PoolContact{restClient}
}

// GET: Queries

type ListPoolContactsForUserParameters struct {
	UserId          string
	PoolId          *string
	PoolContactType *string
}

func (t *PoolContact) ListPoolContactsForUser(p *ListPoolContactsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.PoolContactType != nil {
		queryParameters.Add(`poolContactType`, *p.PoolContactType)
	}

	return t.restClient.Get(
		`/v2/PoolContact/UseCase/ListPoolContactsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreatePoolContactParameters struct {
	PoolId          string
	Email           string
	PoolContactType string
}

func (t *PoolContact) CreatePoolContact(p *CreatePoolContactParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`email`, p.Email)
	queryParameters.Add(`poolContactType`, p.PoolContactType)

	return t.restClient.Post(
		`/v2/PoolContact/UseCase/CreatePoolContact`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolContact) CreatePoolContactWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolContact/UseCase/CreatePoolContact`,
		data,
		nil,
		nil,
	)
}

type RemovePoolContactParameters struct {
	PoolId        string
	PoolContactId string
	UserId        string
}

func (t *PoolContact) RemovePoolContact(p *RemovePoolContactParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`poolContactId`, p.PoolContactId)
	queryParameters.Add(`userId`, p.UserId)

	return t.restClient.Post(
		`/v2/PoolContact/UseCase/RemovePoolContact`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolContact) RemovePoolContactWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolContact/UseCase/RemovePoolContact`,
		data,
		nil,
		nil,
	)
}

type UpdateTypeForPoolContactParameters struct {
	PoolContactId   string
	PoolContactType string
}

func (t *PoolContact) UpdateTypeForPoolContact(p *UpdateTypeForPoolContactParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolContactId`, p.PoolContactId)
	queryParameters.Add(`poolContactType`, p.PoolContactType)

	return t.restClient.Post(
		`/v2/PoolContact/UseCase/UpdateTypeForPoolContact`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolContact) UpdateTypeForPoolContactWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolContact/UseCase/UpdateTypeForPoolContact`,
		data,
		nil,
		nil,
	)
}
