/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type WebConference struct {
	restClient rest.RestClientInterface
}

func NewWebConference(restClient rest.RestClientInterface) *WebConference {
	return &WebConference{restClient}
}

// GET: Queries

type GetWebConferenceRedirectURIByTypeParameters struct {
	PoolId string
	Type   string
}

func (t *WebConference) GetWebConferenceRedirectURIByType(p *GetWebConferenceRedirectURIByTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`type`, p.Type)

	return t.restClient.Get(
		`/v2/WebConference/UseCase/GetWebConferenceRedirectURIByType`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetWebConferenceSettingsForUserParameters struct {
	UserId  string
	EventId string
	Type    string
}

func (t *WebConference) GetWebConferenceSettingsForUser(p *GetWebConferenceSettingsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`type`, p.Type)

	return t.restClient.Get(
		`/v2/WebConference/UseCase/GetWebConferenceSettingsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListWebConferenceConnectionsForPoolParameters struct {
	PoolId string
	Type   *string
}

func (t *WebConference) ListWebConferenceConnectionsForPool(p *ListWebConferenceConnectionsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.Type != nil {
		queryParameters.Add(`type`, *p.Type)
	}

	return t.restClient.Get(
		`/v2/WebConference/UseCase/ListWebConferenceConnectionsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListWebConferencesForAuthUserParameters struct {
	UserId  string
	EventId string
	Type    string
	Format  string
}

func (t *WebConference) ListWebConferencesForAuthUser(p *ListWebConferencesForAuthUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`type`, p.Type)
	queryParameters.Add(`format`, p.Format)

	return t.restClient.Get(
		`/v2/WebConference/UseCase/ListWebConferencesForAuthUser`,
		&queryParameters,
		nil,
		nil,
	)
}

type VerifyWebConferenceSettingsAndCapacityForEventParameters struct {
	EventId string
}

func (t *WebConference) VerifyWebConferenceSettingsAndCapacityForEvent(p *VerifyWebConferenceSettingsAndCapacityForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/WebConference/UseCase/VerifyWebConferenceSettingsAndCapacityForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CompleteWebConferenceDeAuthorizeFlowParameters struct {
	Payload []string
	Type    string
}

func (t *WebConference) CompleteWebConferenceDeAuthorizeFlow(p *CompleteWebConferenceDeAuthorizeFlowParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range p.Payload {
		queryParameters.Add(`payload[]`, p.Payload[i])
	}
	queryParameters.Add(`type`, p.Type)

	return t.restClient.Post(
		`/v2/WebConference/UseCase/CompleteWebConferenceDeAuthorizeFlow`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *WebConference) CompleteWebConferenceDeAuthorizeFlowWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/WebConference/UseCase/CompleteWebConferenceDeAuthorizeFlow`,
		data,
		nil,
		nil,
	)
}

type CompleteWebConferenceOAuthFlowParameters struct {
	Code  string
	State []string
	Type  string
}

func (t *WebConference) CompleteWebConferenceOAuthFlow(p *CompleteWebConferenceOAuthFlowParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`code`, p.Code)
	for i := range p.State {
		queryParameters.Add(`state[]`, p.State[i])
	}
	queryParameters.Add(`type`, p.Type)

	return t.restClient.Post(
		`/v2/WebConference/UseCase/CompleteWebConferenceOAuthFlow`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *WebConference) CompleteWebConferenceOAuthFlowWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/WebConference/UseCase/CompleteWebConferenceOAuthFlow`,
		data,
		nil,
		nil,
	)
}

type FixWebConferenceSettingsForEventParameters struct {
	EventId string
}

func (t *WebConference) FixWebConferenceSettingsForEvent(p *FixWebConferenceSettingsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/WebConference/UseCase/FixWebConferenceSettingsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *WebConference) FixWebConferenceSettingsForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/WebConference/UseCase/FixWebConferenceSettingsForEvent`,
		data,
		nil,
		nil,
	)
}

type QuickCreateWebConferenceForEventParameters struct {
	EventId string
	UserId  string
	Type    string
	Format  string
}

func (t *WebConference) QuickCreateWebConferenceForEvent(p *QuickCreateWebConferenceForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`type`, p.Type)
	queryParameters.Add(`format`, p.Format)

	return t.restClient.Post(
		`/v2/WebConference/UseCase/QuickCreateWebConferenceForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *WebConference) QuickCreateWebConferenceForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/WebConference/UseCase/QuickCreateWebConferenceForEvent`,
		data,
		nil,
		nil,
	)
}

type RemoveWebConferenceConnectionFromPoolParameters struct {
	UserId string
	PoolId string
	Type   string
}

func (t *WebConference) RemoveWebConferenceConnectionFromPool(p *RemoveWebConferenceConnectionFromPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`type`, p.Type)

	return t.restClient.Post(
		`/v2/WebConference/UseCase/RemoveWebConferenceConnectionFromPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *WebConference) RemoveWebConferenceConnectionFromPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/WebConference/UseCase/RemoveWebConferenceConnectionFromPool`,
		data,
		nil,
		nil,
	)
}
