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

type Campaign struct {
	restClient rest.RestClientInterface
}

func NewCampaign(restClient rest.RestClientInterface) *Campaign {
	return &Campaign{restClient}
}

// GET: Queries

type GetCampaignParameters struct {
	CampaignId string
}

func (t *Campaign) GetCampaign(p *GetCampaignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`campaignId`, p.CampaignId)

	return t.restClient.Get(
		`/v2/Campaign/UseCase/GetCampaign`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListCampaignsForPoolParameters struct {
	PoolId        string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
	SortBy        *string
	SortDirection *string
}

func (t *Campaign) ListCampaignsForPool(p *ListCampaignsForPoolParameters) (r *http.Response, err error) {
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
		`/v2/Campaign/UseCase/ListCampaignsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddInvitationToCampaignParameters struct {
	CampaignId   string
	InvitationId string
}

func (t *Campaign) AddInvitationToCampaign(p *AddInvitationToCampaignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`campaignId`, p.CampaignId)
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Campaign/UseCase/AddInvitationToCampaign`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Campaign) AddInvitationToCampaignWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Campaign/UseCase/AddInvitationToCampaign`,
		data,
		nil,
		nil,
	)
}

type ArchiveCampaignParameters struct {
	CampaignId string
}

func (t *Campaign) ArchiveCampaign(p *ArchiveCampaignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`campaignId`, p.CampaignId)

	return t.restClient.Post(
		`/v2/Campaign/UseCase/ArchiveCampaign`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Campaign) ArchiveCampaignWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Campaign/UseCase/ArchiveCampaign`,
		data,
		nil,
		nil,
	)
}

type CreateCampaignParameters struct {
	Name       string
	EventId    string
	CampaignId *string
}

func (t *Campaign) CreateCampaign(p *CreateCampaignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`eventId`, p.EventId)
	if p.CampaignId != nil {
		queryParameters.Add(`campaignId`, *p.CampaignId)
	}

	return t.restClient.Post(
		`/v2/Campaign/UseCase/CreateCampaign`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Campaign) CreateCampaignWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Campaign/UseCase/CreateCampaign`,
		data,
		nil,
		nil,
	)
}

type DeleteCampaignParameters struct {
	CampaignId string
}

func (t *Campaign) DeleteCampaign(p *DeleteCampaignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`campaignId`, p.CampaignId)

	return t.restClient.Post(
		`/v2/Campaign/UseCase/DeleteCampaign`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Campaign) DeleteCampaignWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Campaign/UseCase/DeleteCampaign`,
		data,
		nil,
		nil,
	)
}

type RemoveInvitationFromCampaignParameters struct {
	CampaignId   string
	InvitationId string
}

func (t *Campaign) RemoveInvitationFromCampaign(p *RemoveInvitationFromCampaignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`campaignId`, p.CampaignId)
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Campaign/UseCase/RemoveInvitationFromCampaign`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Campaign) RemoveInvitationFromCampaignWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Campaign/UseCase/RemoveInvitationFromCampaign`,
		data,
		nil,
		nil,
	)
}

type UpdateCampaignParameters struct {
	CampaignId string
	Name       string
	EventId    string
}

func (t *Campaign) UpdateCampaign(p *UpdateCampaignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`campaignId`, p.CampaignId)
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Campaign/UseCase/UpdateCampaign`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Campaign) UpdateCampaignWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Campaign/UseCase/UpdateCampaign`,
		data,
		nil,
		nil,
	)
}
