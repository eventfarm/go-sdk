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

type PoolImage struct {
	restClient rest.RestClientInterface
}

func NewPoolImage(restClient rest.RestClientInterface) *PoolImage {
	return &PoolImage{restClient}
}

// GET: Queries

type GetPoolImageParameters struct {
	PoolImageId string
}

func (t *PoolImage) GetPoolImage(p *GetPoolImageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolImageId`, p.PoolImageId)

	return t.restClient.Get(
		`/v2/PoolImage/UseCase/GetPoolImage`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListPoolImagesForPoolParameters struct {
	PoolId        string
	SortDirection *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
	ImageFilter   *[]string
}

func (t *PoolImage) ListPoolImagesForPool(p *ListPoolImagesForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.ImageFilter != nil {
		for i := range *p.ImageFilter {
			queryParameters.Add(`imageFilter[]`, (*p.ImageFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/PoolImage/UseCase/ListPoolImagesForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreatePoolImageParameters struct {
	PoolId        string
	Image         string
	PoolImageType string
	PoolImageId   *string
}

func (t *PoolImage) CreatePoolImage(p *CreatePoolImageParameters) (r *http.Response, err error) {
	// TODO
	return
}

func (t *PoolImage) CreatePoolImageWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolImage/UseCase/CreatePoolImage`,
		data,
		nil,
		nil,
	)
}
