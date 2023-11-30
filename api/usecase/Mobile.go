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

type Mobile struct {
	restClient rest.RestClientInterface
}

func NewMobile(restClient rest.RestClientInterface) *Mobile {
	return &Mobile{restClient}
}

// GET: Queries

type GetEventDetailsForUserParameters struct {
	EventId string
	UserId  string
}

func (t *Mobile) GetEventDetailsForUser(p *GetEventDetailsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`userId`, p.UserId)

	return t.restClient.Get(
		`/v2/Mobile/UseCase/GetEventDetailsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetInvitationDetailsParameters struct {
	InvitationId string
}

func (t *Mobile) GetInvitationDetails(p *GetInvitationDetailsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Get(
		`/v2/Mobile/UseCase/GetInvitationDetails`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListEventsParameters struct {
	UserId        string
	Query         *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
	SortBy        *string
	SortDirection *string
}

func (t *Mobile) ListEvents(p *ListEventsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
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
		`/v2/Mobile/UseCase/ListEvents`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsParameters struct {
	EventId               string
	Query                 *string
	LastModifiedTimestamp *int64
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-1000
}

func (t *Mobile) ListInvitations(p *ListInvitationsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
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
		`/v2/Mobile/UseCase/ListInvitations`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
