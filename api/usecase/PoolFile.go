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

type PoolFile struct {
	restClient rest.RestClientInterface
}

func NewPoolFile(restClient rest.RestClientInterface) *PoolFile {
	return &PoolFile{restClient}
}

// GET: Queries

type GetPoolAssetParameters struct {
	PoolAssetId string
}

func (t *PoolFile) GetPoolAsset(p *GetPoolAssetParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolAssetId`, p.PoolAssetId)

	return t.restClient.Get(
		`/v2/PoolFile/UseCase/GetPoolAsset`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetPoolDocumentParameters struct {
	PoolDocumentId string
}

func (t *PoolFile) GetPoolDocument(p *GetPoolDocumentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolDocumentId`, p.PoolDocumentId)

	return t.restClient.Get(
		`/v2/PoolFile/UseCase/GetPoolDocument`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListPoolFilesForPoolParameters struct {
	PoolId                 string
	SortDirection          *string
	Page                   *int64 // >= 1
	ItemsPerPage           *int64 // 1-100
	PoolFileCategoryFilter *[]string
	Query                  *string
}

func (t *PoolFile) ListPoolFilesForPool(p *ListPoolFilesForPoolParameters) (r *http.Response, err error) {
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
	if p.PoolFileCategoryFilter != nil {
		for i := range *p.PoolFileCategoryFilter {
			queryParameters.Add(`poolFileCategoryFilter[]`, (*p.PoolFileCategoryFilter)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}

	return t.restClient.Get(
		`/v2/PoolFile/UseCase/ListPoolFilesForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreatePoolAssetParameters struct {
	PoolId               string
	Image                string
	PoolFileCategoryType string
	FileName             string
	PoolAssetId          *string
}

func (t *PoolFile) CreatePoolAsset(p *CreatePoolAssetParameters) (r *http.Response, err error) {
	// TODO
	return
}

func (t *PoolFile) CreatePoolAssetWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolFile/UseCase/CreatePoolAsset`,
		data,
		nil,
		nil,
	)
}

type CreatePoolDocumentParameters struct {
	PoolId               string
	Document             string
	PoolFileCategoryType string
	FileName             string
	PoolDocumentId       *string
}

func (t *PoolFile) CreatePoolDocument(p *CreatePoolDocumentParameters) (r *http.Response, err error) {
	// TODO
	return
}

func (t *PoolFile) CreatePoolDocumentWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolFile/UseCase/CreatePoolDocument`,
		data,
		nil,
		nil,
	)
}
