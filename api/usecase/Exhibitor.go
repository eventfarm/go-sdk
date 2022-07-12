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
	WithData    *[]string
}

func (t *Exhibitor) GetExhibitor(p *GetExhibitorParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Exhibitor/UseCase/GetExhibitor`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListExhibitorsForEventParameters struct {
	EventId           string
	WithData          *[]string // totalUserRolesForExhibitor
	ShouldHideDeleted *bool
	Query             *string
	Page              *int64 // >= 1
	ItemsPerPage      *int64 // 1-500
	SortBy            *string
	SortDirection     *string
}

func (t *Exhibitor) ListExhibitorsForEvent(p *ListExhibitorsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*p.ShouldHideDeleted))
	}
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
	UserId              string
	Query               *string
	WithData            *[]string
	Page                *int64 // >= 1
	ItemsPerPage        *int64 // 1-500
	SortBy              *string
	SortDirection       *string
	EventDateFilterType *string
	PoolId              *string
}

func (t *Exhibitor) ListExhibitorsForUser(p *ListExhibitorsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
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
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}

	return t.restClient.Get(
		`/v2/Exhibitor/UseCase/ListExhibitorsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddUserToExhibitorParameters struct {
	ExhibitorId         string
	Email               string
	FirstName           string
	LastName            string
	AuthenticatedUserId *string
}

func (t *Exhibitor) AddUserToExhibitor(p *AddUserToExhibitorParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)
	queryParameters.Add(`email`, p.Email)
	queryParameters.Add(`firstName`, p.FirstName)
	queryParameters.Add(`lastName`, p.LastName)
	if p.AuthenticatedUserId != nil {
		queryParameters.Add(`authenticatedUserId`, *p.AuthenticatedUserId)
	}

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/AddUserToExhibitor`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) AddUserToExhibitorWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/AddUserToExhibitor`,
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

type DeleteAllExhibitorsAndLeadsForEventParameters struct {
	EventId string
}

func (t *Exhibitor) DeleteAllExhibitorsAndLeadsForEvent(p *DeleteAllExhibitorsAndLeadsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/DeleteAllExhibitorsAndLeadsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) DeleteAllExhibitorsAndLeadsForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/DeleteAllExhibitorsAndLeadsForEvent`,
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

type RemoveUserFromExhibitorParameters struct {
	UserId      string
	ExhibitorId string
}

func (t *Exhibitor) RemoveUserFromExhibitor(p *RemoveUserFromExhibitorParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/RemoveUserFromExhibitor`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) RemoveUserFromExhibitorWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/RemoveUserFromExhibitor`,
		data,
		nil,
		nil,
	)
}

type ResendExhibitorTeamMemberEmailParameters struct {
	ExhibitorId         string
	UserId              string
	AuthenticatedUserId *string
}

func (t *Exhibitor) ResendExhibitorTeamMemberEmail(p *ResendExhibitorTeamMemberEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)
	queryParameters.Add(`userId`, p.UserId)
	if p.AuthenticatedUserId != nil {
		queryParameters.Add(`authenticatedUserId`, *p.AuthenticatedUserId)
	}

	return t.restClient.Post(
		`/v2/Exhibitor/UseCase/ResendExhibitorTeamMemberEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Exhibitor) ResendExhibitorTeamMemberEmailWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Exhibitor/UseCase/ResendExhibitorTeamMemberEmail`,
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
