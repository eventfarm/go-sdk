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

type Country struct {
	restClient rest.RestClientInterface
}

func NewCountry(restClient rest.RestClientInterface) *Country {
	return &Country{restClient}
}

// GET: Queries

type GetCountryParameters struct {
	CountryId string
}

func (t *Country) GetCountry(p *GetCountryParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`countryId`, p.CountryId)

	return t.restClient.Get(
		`/v2/Country/UseCase/GetCountry`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListCountriesParameters struct {
	Name          *string
	SortBy        *string
	SortDirection *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
}

func (t *Country) ListCountries(p *ListCountriesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Country/UseCase/ListCountries`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
