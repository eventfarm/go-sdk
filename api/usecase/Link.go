/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

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

// POST: Commands

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
	Url       *string
	ShownText *string
	LinkId    *string
}

func (t *Link) CreateLink(p *CreateLinkParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.Url != nil {
		queryParameters.Add(`url`, *p.Url)
	}
	if p.ShownText != nil {
		queryParameters.Add(`shownText`, *p.ShownText)
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
