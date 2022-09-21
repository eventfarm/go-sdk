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

type Link struct {
	restClient rest.RestClientInterface
}

func NewLink(restClient rest.RestClientInterface) *Link {
	return &Link{restClient}
}

// GET: Queries

type GetLinkParameters struct {
	LinkId string
}

func (t *Link) GetLink(p *GetLinkParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)

	return t.restClient.Get(
		`/v2/Link/UseCase/GetLink`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListLinksForEventParameters struct {
	EventId           string
	Page              *int64 // >= 1
	ItemsPerPage      *int64 // 1-500
	SortBy            *string
	SortDirection     *string
	Query             *string
	ShouldHideDeleted *bool
}

func (t *Link) ListLinksForEvent(p *ListLinksForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
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
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*p.ShouldHideDeleted))
	}

	return t.restClient.Get(
		`/v2/Link/UseCase/ListLinksForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddLinkToEventParameters struct {
	LinkId  string
	EventId string
}

func (t *Link) AddLinkToEvent(p *AddLinkToEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Link/UseCase/AddLinkToEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) AddLinkToEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/AddLinkToEvent`,
		data,
		nil,
		nil,
	)
}

type AddLinkToProfileParameters struct {
	LinkId    string
	ProfileId string
}

func (t *Link) AddLinkToProfile(p *AddLinkToProfileParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)
	queryParameters.Add(`profileId`, p.ProfileId)

	return t.restClient.Post(
		`/v2/Link/UseCase/AddLinkToProfile`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) AddLinkToProfileWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/AddLinkToProfile`,
		data,
		nil,
		nil,
	)
}

type AddLinksToProfileParameters struct {
	LinkIds   []string
	ProfileId string
}

func (t *Link) AddLinksToProfile(p *AddLinksToProfileParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range p.LinkIds {
		queryParameters.Add(`linkIds[]`, p.LinkIds[i])
	}
	queryParameters.Add(`profileId`, p.ProfileId)

	return t.restClient.Post(
		`/v2/Link/UseCase/AddLinksToProfile`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) AddLinksToProfileWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/AddLinksToProfile`,
		data,
		nil,
		nil,
	)
}

type CreateLinkParameters struct {
	PoolId    string
	Url       string
	ShownText *string
	LinkType  *string
	LinkId    *string
}

func (t *Link) CreateLink(p *CreateLinkParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`url`, p.Url)
	if p.ShownText != nil {
		queryParameters.Add(`shownText`, *p.ShownText)
	}
	if p.LinkType != nil {
		queryParameters.Add(`linkType`, *p.LinkType)
	}
	if p.LinkId != nil {
		queryParameters.Add(`linkId`, *p.LinkId)
	}

	return t.restClient.Post(
		`/v2/Link/UseCase/CreateLink`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) CreateLinkWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/CreateLink`,
		data,
		nil,
		nil,
	)
}

type DeleteLinkParameters struct {
	LinkId string
}

func (t *Link) DeleteLink(p *DeleteLinkParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)

	return t.restClient.Post(
		`/v2/Link/UseCase/DeleteLink`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) DeleteLinkWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/DeleteLink`,
		data,
		nil,
		nil,
	)
}

type RemoveLinkParameters struct {
	LinkId string
}

func (t *Link) RemoveLink(p *RemoveLinkParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)

	return t.restClient.Post(
		`/v2/Link/UseCase/RemoveLink`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) RemoveLinkWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/RemoveLink`,
		data,
		nil,
		nil,
	)
}

type RemoveLinkFromEventParameters struct {
	LinkId  string
	EventId string
}

func (t *Link) RemoveLinkFromEvent(p *RemoveLinkFromEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Link/UseCase/RemoveLinkFromEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) RemoveLinkFromEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/RemoveLinkFromEvent`,
		data,
		nil,
		nil,
	)
}

type RemoveLinkFromProfileParameters struct {
	LinkId    string
	ProfileId string
}

func (t *Link) RemoveLinkFromProfile(p *RemoveLinkFromProfileParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)
	queryParameters.Add(`profileId`, p.ProfileId)

	return t.restClient.Post(
		`/v2/Link/UseCase/RemoveLinkFromProfile`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) RemoveLinkFromProfileWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/RemoveLinkFromProfile`,
		data,
		nil,
		nil,
	)
}

type RemoveLinksFromProfileParameters struct {
	LinkIds   []string
	ProfileId string
}

func (t *Link) RemoveLinksFromProfile(p *RemoveLinksFromProfileParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range p.LinkIds {
		queryParameters.Add(`linkIds[]`, p.LinkIds[i])
	}
	queryParameters.Add(`profileId`, p.ProfileId)

	return t.restClient.Post(
		`/v2/Link/UseCase/RemoveLinksFromProfile`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) RemoveLinksFromProfileWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/RemoveLinksFromProfile`,
		data,
		nil,
		nil,
	)
}

type SetLinkTypeParameters struct {
	LinkId   string
	LinkType string
}

func (t *Link) SetLinkType(p *SetLinkTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)
	queryParameters.Add(`linkType`, p.LinkType)

	return t.restClient.Post(
		`/v2/Link/UseCase/SetLinkType`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) SetLinkTypeWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/SetLinkType`,
		data,
		nil,
		nil,
	)
}

type SetShownTextParameters struct {
	LinkId    string
	ShownText string
}

func (t *Link) SetShownText(p *SetShownTextParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)
	queryParameters.Add(`shownText`, p.ShownText)

	return t.restClient.Post(
		`/v2/Link/UseCase/SetShownText`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) SetShownTextWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/SetShownText`,
		data,
		nil,
		nil,
	)
}

type SetUrlParameters struct {
	LinkId string
	Url    string
}

func (t *Link) SetUrl(p *SetUrlParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)
	queryParameters.Add(`url`, p.Url)

	return t.restClient.Post(
		`/v2/Link/UseCase/SetUrl`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) SetUrlWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/SetUrl`,
		data,
		nil,
		nil,
	)
}

type UpdateLinkParameters struct {
	LinkId    string
	PoolId    string
	Url       *string
	ShownText *string
	LinkType  *string
}

func (t *Link) UpdateLink(p *UpdateLinkParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`linkId`, p.LinkId)
	queryParameters.Add(`poolId`, p.PoolId)
	if p.Url != nil {
		queryParameters.Add(`url`, *p.Url)
	}
	if p.ShownText != nil {
		queryParameters.Add(`shownText`, *p.ShownText)
	}
	if p.LinkType != nil {
		queryParameters.Add(`linkType`, *p.LinkType)
	}

	return t.restClient.Post(
		`/v2/Link/UseCase/UpdateLink`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Link) UpdateLinkWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Link/UseCase/UpdateLink`,
		data,
		nil,
		nil,
	)
}
