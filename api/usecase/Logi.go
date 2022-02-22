/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type Logi struct {
	restClient rest.RestClientInterface
}

func NewLogi(restClient rest.RestClientInterface) *Logi {
	return &Logi{restClient}
}

// GET: Queries

type GetLogiURLForEventParameters struct {
	EventId string
	UserId  string
}

func (t *Logi) GetLogiURLForEvent(p *GetLogiURLForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`userId`, p.UserId)

	return t.restClient.Get(
		`/v2/Logi/UseCase/GetLogiURLForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
