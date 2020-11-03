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

type PoolWorld struct {
	restClient rest.RestClientInterface
}

func NewPoolWorld(restClient rest.RestClientInterface) *PoolWorld {
	return &PoolWorld{restClient}
}

// GET: Queries

type GetPoolWorldByEventParameters struct {
	EventId string
}

func (t *PoolWorld) GetPoolWorldByEvent(p *GetPoolWorldByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/PoolWorld/UseCase/GetPoolWorldByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListPoolWorldsForPoolParameters struct {
	PoolId string
}

func (t *PoolWorld) ListPoolWorldsForPool(p *ListPoolWorldsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/PoolWorld/UseCase/ListPoolWorldsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreatePoolWorldParameters struct {
	PoolId         string
	VirbelaWorldId string
}

func (t *PoolWorld) CreatePoolWorld(p *CreatePoolWorldParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`virbelaWorldId`, p.VirbelaWorldId)

	return t.restClient.Post(
		`/v2/PoolWorld/UseCase/CreatePoolWorld`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolWorld) CreatePoolWorldWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolWorld/UseCase/CreatePoolWorld`,
		data,
		nil,
		nil,
	)
}

type RemovePoolWorldParameters struct {
	PoolWorldId string
}

func (t *PoolWorld) RemovePoolWorld(p *RemovePoolWorldParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolWorldId`, p.PoolWorldId)

	return t.restClient.Post(
		`/v2/PoolWorld/UseCase/RemovePoolWorld`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolWorld) RemovePoolWorldWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolWorld/UseCase/RemovePoolWorld`,
		data,
		nil,
		nil,
	)
}

type UpdatePoolWorldParameters struct {
	PoolWorldId  string
	TitleMapping string
	TeamId       int64 // 0-51
}

func (t *PoolWorld) UpdatePoolWorld(p *UpdatePoolWorldParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolWorldId`, p.PoolWorldId)
	queryParameters.Add(`titleMapping`, p.TitleMapping)
	queryParameters.Add(`teamId`, strconv.FormatInt(p.TeamId, 10))

	return t.restClient.Post(
		`/v2/PoolWorld/UseCase/UpdatePoolWorld`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolWorld) UpdatePoolWorldWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolWorld/UseCase/UpdatePoolWorld`,
		data,
		nil,
		nil,
	)
}
