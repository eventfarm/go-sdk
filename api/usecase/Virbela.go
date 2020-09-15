/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type Virbela struct {
	restClient rest.RestClientInterface
}

func NewVirbela(restClient rest.RestClientInterface) *Virbela {
	return &Virbela{restClient}
}

// GET: Queries

type CreateUserForVirbelaParameters struct {
	InvitationId string
	Password     string
}

func (t *Virbela) CreateUserForVirbela(p *CreateUserForVirbelaParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`password`, p.Password)

	return t.restClient.Get(
		`/v2/Virbela/UseCase/CreateUserForVirbela`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetUserStatusForVirbelaParameters struct {
	InvitationId string
}

func (t *Virbela) GetUserStatusForVirbela(p *GetUserStatusForVirbelaParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Get(
		`/v2/Virbela/UseCase/GetUserStatusForVirbela`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListVirbelaWorldsForEventFarmParameters struct {
	PoolId    string
	WorldName *string
	WorldId   *string
}

func (t *Virbela) ListVirbelaWorldsForEventFarm(p *ListVirbelaWorldsForEventFarmParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.WorldName != nil {
		queryParameters.Add(`worldName`, *p.WorldName)
	}
	if p.WorldId != nil {
		queryParameters.Add(`worldId`, *p.WorldId)
	}

	return t.restClient.Get(
		`/v2/Virbela/UseCase/ListVirbelaWorldsForEventFarm`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type SuspendAllUsersForVirbelaParameters struct {
	EventId string
}

func (t *Virbela) SuspendAllUsersForVirbela(p *SuspendAllUsersForVirbelaParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Virbela/UseCase/SuspendAllUsersForVirbela`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Virbela) SuspendAllUsersForVirbelaWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Virbela/UseCase/SuspendAllUsersForVirbela`,
		data,
		nil,
		nil,
	)
}

type SuspendUserForVirbelaParameters struct {
	InvitationId string
}

func (t *Virbela) SuspendUserForVirbela(p *SuspendUserForVirbelaParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Virbela/UseCase/SuspendUserForVirbela`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Virbela) SuspendUserForVirbelaWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Virbela/UseCase/SuspendUserForVirbela`,
		data,
		nil,
		nil,
	)
}

type UpdateUserRoleForVirbelaParameters struct {
	InvitationId string
	Role         string
}

func (t *Virbela) UpdateUserRoleForVirbela(p *UpdateUserRoleForVirbelaParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`role`, p.Role)

	return t.restClient.Post(
		`/v2/Virbela/UseCase/UpdateUserRoleForVirbela`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Virbela) UpdateUserRoleForVirbelaWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Virbela/UseCase/UpdateUserRoleForVirbela`,
		data,
		nil,
		nil,
	)
}

type UpdateUserTeamIdForVirbelaParameters struct {
	InvitationId string
	TeamId       string
}

func (t *Virbela) UpdateUserTeamIdForVirbela(p *UpdateUserTeamIdForVirbelaParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`teamId`, p.TeamId)

	return t.restClient.Post(
		`/v2/Virbela/UseCase/UpdateUserTeamIdForVirbela`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Virbela) UpdateUserTeamIdForVirbelaWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Virbela/UseCase/UpdateUserTeamIdForVirbela`,
		data,
		nil,
		nil,
	)
}

type UpdateUserTitleForVirbelaParameters struct {
	InvitationId string
	Title        string
}

func (t *Virbela) UpdateUserTitleForVirbela(p *UpdateUserTitleForVirbelaParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`title`, p.Title)

	return t.restClient.Post(
		`/v2/Virbela/UseCase/UpdateUserTitleForVirbela`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Virbela) UpdateUserTitleForVirbelaWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Virbela/UseCase/UpdateUserTitleForVirbela`,
		data,
		nil,
		nil,
	)
}
