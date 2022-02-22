/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/eventfarm/go-sdk/rest"
)

type Lead struct {
	restClient rest.RestClientInterface
}

func NewLead(restClient rest.RestClientInterface) *Lead {
	return &Lead{restClient}
}

// GET: Queries

type GetLeadParameters struct {
	LeadId string
}

func (t *Lead) GetLead(p *GetLeadParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`leadId`, p.LeadId)

	return t.restClient.Get(
		`/v2/Lead/UseCase/GetLead`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListLeadsForExhibitorParameters struct {
	ExhibitorId   string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
	SortBy        *string
	SortDirection *string
}

func (t *Lead) ListLeadsForExhibitor(p *ListLeadsForExhibitorParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)
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
		`/v2/Lead/UseCase/ListLeadsForExhibitor`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateLeadParameters struct {
	ExhibitorId  string
	FirstName    string
	LastName     string
	EmailAddress *string
	PhoneNumber  *string
	Score        *float64
	Notes        *string
	SourceUserId *string
	LeadId       *string
}

func (t *Lead) CreateLead(p *CreateLeadParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`exhibitorId`, p.ExhibitorId)
	queryParameters.Add(`firstName`, p.FirstName)
	queryParameters.Add(`lastName`, p.LastName)
	if p.EmailAddress != nil {
		queryParameters.Add(`emailAddress`, *p.EmailAddress)
	}
	if p.PhoneNumber != nil {
		queryParameters.Add(`phoneNumber`, *p.PhoneNumber)
	}
	if p.Score != nil {
		queryParameters.Add(`score`, fmt.Sprintf("%f", *p.Score))
	}
	if p.Notes != nil {
		queryParameters.Add(`notes`, *p.Notes)
	}
	if p.SourceUserId != nil {
		queryParameters.Add(`sourceUserId`, *p.SourceUserId)
	}
	if p.LeadId != nil {
		queryParameters.Add(`leadId`, *p.LeadId)
	}

	return t.restClient.Post(
		`/v2/Lead/UseCase/CreateLead`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Lead) CreateLeadWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Lead/UseCase/CreateLead`,
		data,
		nil,
		nil,
	)
}

type DeleteLeadParameters struct {
	LeadId string
}

func (t *Lead) DeleteLead(p *DeleteLeadParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`leadId`, p.LeadId)

	return t.restClient.Post(
		`/v2/Lead/UseCase/DeleteLead`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Lead) DeleteLeadWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Lead/UseCase/DeleteLead`,
		data,
		nil,
		nil,
	)
}

type RemoveLeadParameters struct {
	LeadId string
}

func (t *Lead) RemoveLead(p *RemoveLeadParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`leadId`, p.LeadId)

	return t.restClient.Post(
		`/v2/Lead/UseCase/RemoveLead`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Lead) RemoveLeadWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Lead/UseCase/RemoveLead`,
		data,
		nil,
		nil,
	)
}

type SetLeadEmailAddressParameters struct {
	LeadId           string
	LeadEmailAddress string
}

func (t *Lead) SetLeadEmailAddress(p *SetLeadEmailAddressParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`leadId`, p.LeadId)
	queryParameters.Add(`leadEmailAddress`, p.LeadEmailAddress)

	return t.restClient.Post(
		`/v2/Lead/UseCase/SetLeadEmailAddress`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Lead) SetLeadEmailAddressWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Lead/UseCase/SetLeadEmailAddress`,
		data,
		nil,
		nil,
	)
}

type SetLeadFirstNameParameters struct {
	LeadId        string
	LeadFirstName string
}

func (t *Lead) SetLeadFirstName(p *SetLeadFirstNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`leadId`, p.LeadId)
	queryParameters.Add(`leadFirstName`, p.LeadFirstName)

	return t.restClient.Post(
		`/v2/Lead/UseCase/SetLeadFirstName`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Lead) SetLeadFirstNameWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Lead/UseCase/SetLeadFirstName`,
		data,
		nil,
		nil,
	)
}

type SetLeadLastNameParameters struct {
	LeadId       string
	LeadLastName string
}

func (t *Lead) SetLeadLastName(p *SetLeadLastNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`leadId`, p.LeadId)
	queryParameters.Add(`leadLastName`, p.LeadLastName)

	return t.restClient.Post(
		`/v2/Lead/UseCase/SetLeadLastName`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Lead) SetLeadLastNameWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Lead/UseCase/SetLeadLastName`,
		data,
		nil,
		nil,
	)
}

type SetLeadNotesParameters struct {
	LeadId    string
	LeadNotes string
}

func (t *Lead) SetLeadNotes(p *SetLeadNotesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`leadId`, p.LeadId)
	queryParameters.Add(`leadNotes`, p.LeadNotes)

	return t.restClient.Post(
		`/v2/Lead/UseCase/SetLeadNotes`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Lead) SetLeadNotesWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Lead/UseCase/SetLeadNotes`,
		data,
		nil,
		nil,
	)
}

type SetLeadPhoneNumberParameters struct {
	LeadId          string
	LeadPhoneNumber string
}

func (t *Lead) SetLeadPhoneNumber(p *SetLeadPhoneNumberParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`leadId`, p.LeadId)
	queryParameters.Add(`leadPhoneNumber`, p.LeadPhoneNumber)

	return t.restClient.Post(
		`/v2/Lead/UseCase/SetLeadPhoneNumber`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Lead) SetLeadPhoneNumberWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Lead/UseCase/SetLeadPhoneNumber`,
		data,
		nil,
		nil,
	)
}

type SetLeadScoreParameters struct {
	LeadId    string
	LeadScore float64
}

func (t *Lead) SetLeadScore(p *SetLeadScoreParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`leadId`, p.LeadId)
	queryParameters.Add(`leadScore`, fmt.Sprintf("%f", p.LeadScore))

	return t.restClient.Post(
		`/v2/Lead/UseCase/SetLeadScore`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Lead) SetLeadScoreWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Lead/UseCase/SetLeadScore`,
		data,
		nil,
		nil,
	)
}
