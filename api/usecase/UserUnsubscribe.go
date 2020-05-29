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

type UserUnsubscribe struct {
	restClient rest.RestClientInterface
}

func NewUserUnsubscribe(restClient rest.RestClientInterface) *UserUnsubscribe {
	return &UserUnsubscribe{restClient}
}

// GET: Queries

type ListSubscriptionsByEmailParameters struct {
	EmailAddress string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-100
}

func (t *UserUnsubscribe) ListSubscriptionsByEmail(p *ListSubscriptionsByEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailAddress`, p.EmailAddress)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/UserUnsubscribe/UseCase/ListSubscriptionsByEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type ResubscribeUserToEventParameters struct {
	EventId      string
	EmailAddress string
}

func (t *UserUnsubscribe) ResubscribeUserToEvent(p *ResubscribeUserToEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`emailAddress`, p.EmailAddress)

	return t.restClient.Post(
		`/v2/UserUnsubscribe/UseCase/ResubscribeUserToEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *UserUnsubscribe) ResubscribeUserToEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/UserUnsubscribe/UseCase/ResubscribeUserToEvent`,
		data,
		nil,
		nil,
	)
}

type ResubscribeUserToPoolParameters struct {
	PoolId       string
	EmailAddress string
}

func (t *UserUnsubscribe) ResubscribeUserToPool(p *ResubscribeUserToPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`emailAddress`, p.EmailAddress)

	return t.restClient.Post(
		`/v2/UserUnsubscribe/UseCase/ResubscribeUserToPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *UserUnsubscribe) ResubscribeUserToPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/UserUnsubscribe/UseCase/ResubscribeUserToPool`,
		data,
		nil,
		nil,
	)
}

type ResubscribeUserToTicketblockParameters struct {
	TicketblockId string
	EmailAddress  string
}

func (t *UserUnsubscribe) ResubscribeUserToTicketblock(p *ResubscribeUserToTicketblockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketblockId`, p.TicketblockId)
	queryParameters.Add(`emailAddress`, p.EmailAddress)

	return t.restClient.Post(
		`/v2/UserUnsubscribe/UseCase/ResubscribeUserToTicketblock`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *UserUnsubscribe) ResubscribeUserToTicketblockWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/UserUnsubscribe/UseCase/ResubscribeUserToTicketblock`,
		data,
		nil,
		nil,
	)
}

type UnsubscribeUserFromEventParameters struct {
	EventId           string
	EmailAddress      string
	UserUnsubscribeId *string
}

func (t *UserUnsubscribe) UnsubscribeUserFromEvent(p *UnsubscribeUserFromEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`emailAddress`, p.EmailAddress)
	if p.UserUnsubscribeId != nil {
		queryParameters.Add(`userUnsubscribeId`, *p.UserUnsubscribeId)
	}

	return t.restClient.Post(
		`/v2/UserUnsubscribe/UseCase/UnsubscribeUserFromEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *UserUnsubscribe) UnsubscribeUserFromEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/UserUnsubscribe/UseCase/UnsubscribeUserFromEvent`,
		data,
		nil,
		nil,
	)
}

type UnsubscribeUserFromPoolParameters struct {
	PoolId            string
	EmailAddress      string
	UserUnsubscribeId *string
}

func (t *UserUnsubscribe) UnsubscribeUserFromPool(p *UnsubscribeUserFromPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`emailAddress`, p.EmailAddress)
	if p.UserUnsubscribeId != nil {
		queryParameters.Add(`userUnsubscribeId`, *p.UserUnsubscribeId)
	}

	return t.restClient.Post(
		`/v2/UserUnsubscribe/UseCase/UnsubscribeUserFromPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *UserUnsubscribe) UnsubscribeUserFromPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/UserUnsubscribe/UseCase/UnsubscribeUserFromPool`,
		data,
		nil,
		nil,
	)
}

type UnsubscribeUserFromTicketblockParameters struct {
	TicketblockId     string
	EmailAddress      string
	UserUnsubscribeId *string
}

func (t *UserUnsubscribe) UnsubscribeUserFromTicketblock(p *UnsubscribeUserFromTicketblockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketblockId`, p.TicketblockId)
	queryParameters.Add(`emailAddress`, p.EmailAddress)
	if p.UserUnsubscribeId != nil {
		queryParameters.Add(`userUnsubscribeId`, *p.UserUnsubscribeId)
	}

	return t.restClient.Post(
		`/v2/UserUnsubscribe/UseCase/UnsubscribeUserFromTicketblock`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *UserUnsubscribe) UnsubscribeUserFromTicketblockWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/UserUnsubscribe/UseCase/UnsubscribeUserFromTicketblock`,
		data,
		nil,
		nil,
	)
}
