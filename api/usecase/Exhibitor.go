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

type Exhibitor struct {
	restClient rest.RestClientInterface
}

func NewExhibitor(restClient rest.RestClientInterface) *Exhibitor {
	return &Exhibitor{restClient}
}

// GET: Queries

type GetExhibitorParameters struct {
	ExhibitorId string
}

func (t *Exhibitor) GetExhibitor(p *GetExhibitorParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)

	return t.restClient.Get(
		`/v2/Exhibitor/UseCase/GetExhibitor`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListExhibitorsForEventParameters struct {
	EventId       string
	Query         *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
	SortBy        *string
	SortDirection *string
}

func (t *Exhibitor) ListExhibitorsForEvent(p *ListExhibitorsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
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
		`/v2/Exhibitor/UseCase/ListExhibitorsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListExhibitorsForUserParameters struct {
	PoolId           string
	UserId           string
	GroupOwnerUserId *string
	Page             *int64 // >= 1
	ItemsPerPage     *int64 // 1-500
	SortBy           *string
	SortDirection    *string
}

func (t *Exhibitor) ListExhibitorsForUser(p *ListExhibitorsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	if p.GroupOwnerUserId != nil {
		queryParameters.Add(`groupOwnerUserId`, *p.GroupOwnerUserId)
	}
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
		`/v2/Exhibitor/UseCase/ListExhibitorsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddUsersToExhibitorParameters struct {
	ExhibitorId string
	UserIds     []string
}

func (t *Exhibitor) AddUsersToExhibitor(p *AddUsersToExhibitorParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)
	for i := range p.UserIds {
		queryParameters.Add(`userIds[]`, p.UserIds[i])
	}

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/AddUsersToExhibitor`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) AddUsersToExhibitorWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/AddUsersToExhibitor`,
		data,
		nil,
		nil,
	)
}

type CreateExhibitorForEventParameters struct {
	EventId       string
	ExhibitorName string
	Description   *string
	LogoUrl       *string
	ExhibitorId   *string
}

func (t *Exhibitor) CreateExhibitorForEvent(p *CreateExhibitorForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`exhibitorName`, p.ExhibitorName)
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}
	if p.LogoUrl != nil {
		queryParameters.Add(`logoUrl`, *p.LogoUrl)
	}
	if p.ExhibitorId != nil {
		queryParameters.Add(`exhibitorId`, *p.ExhibitorId)
	}

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/CreateExhibitorForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) CreateExhibitorForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/CreateExhibitorForEvent`,
		data,
		nil,
		nil,
	)
}

type DeleteExhibitorParameters struct {
	ExhibitorId string
}

func (t *Exhibitor) DeleteExhibitor(p *DeleteExhibitorParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/DeleteExhibitor`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) DeleteExhibitorWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/DeleteExhibitor`,
		data,
		nil,
		nil,
	)
}

type RemoveExhibitorParameters struct {
	ExhibitorId string
}

func (t *Exhibitor) RemoveExhibitor(p *RemoveExhibitorParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/RemoveExhibitor`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) RemoveExhibitorWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/RemoveExhibitor`,
		data,
		nil,
		nil,
	)
}

type RemoveUsersFromExhibitorParameters struct {
	ExhibitorId string
	UserIds     []string
}

func (t *Exhibitor) RemoveUsersFromExhibitor(p *RemoveUsersFromExhibitorParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)
	for i := range p.UserIds {
		queryParameters.Add(`userIds[]`, p.UserIds[i])
	}

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/RemoveUsersFromExhibitor`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) RemoveUsersFromExhibitorWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/RemoveUsersFromExhibitor`,
		data,
		nil,
		nil,
	)
}

type SetExhibitorDescriptionParameters struct {
	ExhibitorId          string
	ExhibitorDescription string
}

func (t *Exhibitor) SetExhibitorDescription(p *SetExhibitorDescriptionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)
	queryParameters.Add(`exhibitorDescription`, p.ExhibitorDescription)

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/SetExhibitorDescription`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) SetExhibitorDescriptionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/SetExhibitorDescription`,
		data,
		nil,
		nil,
	)
}

type SetExhibitorLogoUrlParameters struct {
	ExhibitorId      string
	ExhibitorLogoURL string
}

func (t *Exhibitor) SetExhibitorLogoUrl(p *SetExhibitorLogoUrlParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)
	queryParameters.Add(`exhibitorLogoURL`, p.ExhibitorLogoURL)

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/SetExhibitorLogoUrl`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) SetExhibitorLogoUrlWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/SetExhibitorLogoUrl`,
		data,
		nil,
		nil,
	)
}

type SetExhibitorNameParameters struct {
	ExhibitorId   string
	ExhibitorName string
}

func (t *Exhibitor) SetExhibitorName(p *SetExhibitorNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)
	queryParameters.Add(`exhibitorName`, p.ExhibitorName)

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/SetExhibitorName`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) SetExhibitorNameWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/SetExhibitorName`,
		data,
		nil,
		nil,
	)
}
