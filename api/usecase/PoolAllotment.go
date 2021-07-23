/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/eventfarm/go-sdk/rest"
)

type PoolAllotment struct {
	restClient rest.RestClientInterface
}

func NewPoolAllotment(restClient rest.RestClientInterface) *PoolAllotment {
	return &PoolAllotment{restClient}
}

// GET: Queries

// POST: Commands

type CreatePoolAllotmentParameters struct {
	PoolId            string
	PoolAllotmentType string
	Amount            int64
	PoolAllotmentId   *string
}

func (t *PoolAllotment) CreatePoolAllotment(p *CreatePoolAllotmentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`poolAllotmentType`, p.PoolAllotmentType)
	queryParameters.Add(`amount`, strconv.FormatInt(p.Amount, 10))
	if p.PoolAllotmentId != nil {
		queryParameters.Add(`poolAllotmentId`, *p.PoolAllotmentId)
	}

	return t.restClient.Post(
		`/v2/PoolAllotment/UseCase/CreatePoolAllotment`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolAllotment) CreatePoolAllotmentWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolAllotment/UseCase/CreatePoolAllotment`,
		data,
		nil,
		nil,
	)
}

type SetPoolAllotmentAmountParameters struct {
	PoolAllotmentId string
	Amount          int64
}

func (t *PoolAllotment) SetPoolAllotmentAmount(p *SetPoolAllotmentAmountParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolAllotmentId`, p.PoolAllotmentId)
	queryParameters.Add(`amount`, strconv.FormatInt(p.Amount, 10))

	return t.restClient.Post(
		`/v2/PoolAllotment/UseCase/SetPoolAllotmentAmount`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolAllotment) SetPoolAllotmentAmountWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolAllotment/UseCase/SetPoolAllotmentAmount`,
		data,
		nil,
		nil,
	)
}
