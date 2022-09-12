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

type Venue struct {
	restClient rest.RestClientInterface
}

func NewVenue(restClient rest.RestClientInterface) *Venue {
	return &Venue{restClient}
}

// GET: Queries

type GetVenueParameters struct {
	VenueId string
}

func (t *Venue) GetVenue(p *GetVenueParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`venueId`, p.VenueId)

	return t.restClient.Get(
		`/v2/Venue/UseCase/GetVenue`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetVenueLocationParameters struct {
	VenueLocationId string
}

func (t *Venue) GetVenueLocation(p *GetVenueLocationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`venueLocationId`, p.VenueLocationId)

	return t.restClient.Get(
		`/v2/Venue/UseCase/GetVenueLocation`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListVenuesForPoolParameters struct {
	PoolId        string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
	SortBy        *string
	SortDirection *string
}

func (t *Venue) ListVenuesForPool(p *ListVenuesForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}

	return t.restClient.Get(
		`/v2/Venue/UseCase/ListVenuesForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateVenueParameters struct {
	PoolId      string
	Name        string
	VenueType   *string
	Street1     *string
	Street2     *string
	City        *string
	County      *string
	Province    *string
	CountryId   *int64
	PostalCode  *string
	Description *string
	Url         *string
	Capacity    *int64
	VenueId     *string
}

func (t *Venue) CreateVenue(p *CreateVenueParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`name`, p.Name)
	if p.VenueType != nil {
		queryParameters.Add(`venueType`, *p.VenueType)
	}
	if p.Street1 != nil {
		queryParameters.Add(`street1`, *p.Street1)
	}
	if p.Street2 != nil {
		queryParameters.Add(`street2`, *p.Street2)
	}
	if p.City != nil {
		queryParameters.Add(`city`, *p.City)
	}
	if p.County != nil {
		queryParameters.Add(`county`, *p.County)
	}
	if p.Province != nil {
		queryParameters.Add(`province`, *p.Province)
	}
	if p.CountryId != nil {
		queryParameters.Add(`countryId`, strconv.FormatInt(*p.CountryId, 10))
	}
	if p.PostalCode != nil {
		queryParameters.Add(`postalCode`, *p.PostalCode)
	}
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}
	if p.Url != nil {
		queryParameters.Add(`url`, *p.Url)
	}
	if p.Capacity != nil {
		queryParameters.Add(`capacity`, strconv.FormatInt(*p.Capacity, 10))
	}
	if p.VenueId != nil {
		queryParameters.Add(`venueId`, *p.VenueId)
	}

	return t.restClient.Post(
		`/v2/Venue/UseCase/CreateVenue`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Venue) CreateVenueWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Venue/UseCase/CreateVenue`,
		data,
		nil,
		nil,
	)
}

type CreateVenueLocationParameters struct {
	VenueId         string
	Name            string
	GpsCoordinates  *string
	VenueLocationId *string
}

func (t *Venue) CreateVenueLocation(p *CreateVenueLocationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`venueId`, p.VenueId)
	queryParameters.Add(`name`, p.Name)
	if p.GpsCoordinates != nil {
		queryParameters.Add(`gpsCoordinates`, *p.GpsCoordinates)
	}
	if p.VenueLocationId != nil {
		queryParameters.Add(`venueLocationId`, *p.VenueLocationId)
	}

	return t.restClient.Post(
		`/v2/Venue/UseCase/CreateVenueLocation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Venue) CreateVenueLocationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Venue/UseCase/CreateVenueLocation`,
		data,
		nil,
		nil,
	)
}

type DeleteVenueParameters struct {
	VenueId string
}

func (t *Venue) DeleteVenue(p *DeleteVenueParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`venueId`, p.VenueId)

	return t.restClient.Post(
		`/v2/Venue/UseCase/DeleteVenue`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Venue) DeleteVenueWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Venue/UseCase/DeleteVenue`,
		data,
		nil,
		nil,
	)
}

type DeleteVenueLocationParameters struct {
	VenueLocationId string
}

func (t *Venue) DeleteVenueLocation(p *DeleteVenueLocationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`venueLocationId`, p.VenueLocationId)

	return t.restClient.Post(
		`/v2/Venue/UseCase/DeleteVenueLocation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Venue) DeleteVenueLocationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Venue/UseCase/DeleteVenueLocation`,
		data,
		nil,
		nil,
	)
}

type RemoveVenueParameters struct {
	VenueId string
}

func (t *Venue) RemoveVenue(p *RemoveVenueParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`venueId`, p.VenueId)

	return t.restClient.Post(
		`/v2/Venue/UseCase/RemoveVenue`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Venue) RemoveVenueWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Venue/UseCase/RemoveVenue`,
		data,
		nil,
		nil,
	)
}

type RemoveVenueLocationParameters struct {
	VenueLocationId string
}

func (t *Venue) RemoveVenueLocation(p *RemoveVenueLocationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`venueLocationId`, p.VenueLocationId)

	return t.restClient.Post(
		`/v2/Venue/UseCase/RemoveVenueLocation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Venue) RemoveVenueLocationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Venue/UseCase/RemoveVenueLocation`,
		data,
		nil,
		nil,
	)
}

type SetVenueLocationParameters struct {
	VenueId         string
	VenueLocationId string
}

func (t *Venue) SetVenueLocation(p *SetVenueLocationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`venueId`, p.VenueId)
	queryParameters.Add(`venueLocationId`, p.VenueLocationId)

	return t.restClient.Post(
		`/v2/Venue/UseCase/SetVenueLocation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Venue) SetVenueLocationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Venue/UseCase/SetVenueLocation`,
		data,
		nil,
		nil,
	)
}

type UpdateVenueParameters struct {
	VenueId     string
	PoolId      string
	Name        string
	VenueType   *string
	Street1     *string
	Street2     *string
	City        *string
	County      *string
	Province    *string
	CountryId   *int64
	PostalCode  *string
	Description *string
	Url         *string
	Capacity    *int64
}

func (t *Venue) UpdateVenue(p *UpdateVenueParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`venueId`, p.VenueId)
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`name`, p.Name)
	if p.VenueType != nil {
		queryParameters.Add(`venueType`, *p.VenueType)
	}
	if p.Street1 != nil {
		queryParameters.Add(`street1`, *p.Street1)
	}
	if p.Street2 != nil {
		queryParameters.Add(`street2`, *p.Street2)
	}
	if p.City != nil {
		queryParameters.Add(`city`, *p.City)
	}
	if p.County != nil {
		queryParameters.Add(`county`, *p.County)
	}
	if p.Province != nil {
		queryParameters.Add(`province`, *p.Province)
	}
	if p.CountryId != nil {
		queryParameters.Add(`countryId`, strconv.FormatInt(*p.CountryId, 10))
	}
	if p.PostalCode != nil {
		queryParameters.Add(`postalCode`, *p.PostalCode)
	}
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}
	if p.Url != nil {
		queryParameters.Add(`url`, *p.Url)
	}
	if p.Capacity != nil {
		queryParameters.Add(`capacity`, strconv.FormatInt(*p.Capacity, 10))
	}

	return t.restClient.Post(
		`/v2/Venue/UseCase/UpdateVenue`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Venue) UpdateVenueWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Venue/UseCase/UpdateVenue`,
		data,
		nil,
		nil,
	)
}

type UpdateVenueLocationParameters struct {
	VenueLocationId string
	VenueId         *string
	Name            *string
	GpsCoordinates  *string
}

func (t *Venue) UpdateVenueLocation(p *UpdateVenueLocationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`venueLocationId`, p.VenueLocationId)
	if p.VenueId != nil {
		queryParameters.Add(`venueId`, *p.VenueId)
	}
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.GpsCoordinates != nil {
		queryParameters.Add(`gpsCoordinates`, *p.GpsCoordinates)
	}

	return t.restClient.Post(
		`/v2/Venue/UseCase/UpdateVenueLocation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Venue) UpdateVenueLocationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Venue/UseCase/UpdateVenueLocation`,
		data,
		nil,
		nil,
	)
}
