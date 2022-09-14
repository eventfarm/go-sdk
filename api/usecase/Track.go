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

type Track struct {
	restClient rest.RestClientInterface
}

func NewTrack(restClient rest.RestClientInterface) *Track {
	return &Track{restClient}
}

// GET: Queries

type GetTrackParameters struct {
	TrackId string
}

func (t *Track) GetTrack(p *GetTrackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`trackId`, p.TrackId)

	return t.restClient.Get(
		`/v2/Track/UseCase/GetTrack`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListTracksForEventParameters struct {
	EventId       string
	WithData      *[]string // Event
	SortBy        *string
	SortDirection *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
}

func (t *Track) ListTracksForEvent(p *ListTracksForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
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
		`/v2/Track/UseCase/ListTracksForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateTrackParameters struct {
	EventId     string
	Name        string
	Description *string
	TrackId     *string
}

func (t *Track) CreateTrack(p *CreateTrackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`name`, p.Name)
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}
	if p.TrackId != nil {
		queryParameters.Add(`trackId`, *p.TrackId)
	}

	return t.restClient.Post(
		`/v2/Track/UseCase/CreateTrack`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Track) CreateTrackWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Track/UseCase/CreateTrack`,
		data,
		nil,
		nil,
	)
}

type DeleteTrackParameters struct {
	TrackId string
}

func (t *Track) DeleteTrack(p *DeleteTrackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`trackId`, p.TrackId)

	return t.restClient.Post(
		`/v2/Track/UseCase/DeleteTrack`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Track) DeleteTrackWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Track/UseCase/DeleteTrack`,
		data,
		nil,
		nil,
	)
}

type RemoveTrackParameters struct {
	TrackId string
}

func (t *Track) RemoveTrack(p *RemoveTrackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`trackId`, p.TrackId)

	return t.restClient.Post(
		`/v2/Track/UseCase/RemoveTrack`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Track) RemoveTrackWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Track/UseCase/RemoveTrack`,
		data,
		nil,
		nil,
	)
}

type SetTrackDescriptionParameters struct {
	TrackId          string
	TrackDescription string
}

func (t *Track) SetTrackDescription(p *SetTrackDescriptionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`trackId`, p.TrackId)
	queryParameters.Add(`trackDescription`, p.TrackDescription)

	return t.restClient.Post(
		`/v2/Track/UseCase/SetTrackDescription`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Track) SetTrackDescriptionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Track/UseCase/SetTrackDescription`,
		data,
		nil,
		nil,
	)
}

type SetTrackNameParameters struct {
	TrackId   string
	TrackName string
}

func (t *Track) SetTrackName(p *SetTrackNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`trackId`, p.TrackId)
	queryParameters.Add(`trackName`, p.TrackName)

	return t.restClient.Post(
		`/v2/Track/UseCase/SetTrackName`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Track) SetTrackNameWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Track/UseCase/SetTrackName`,
		data,
		nil,
		nil,
	)
}

type UpdateTrackParameters struct {
	TrackId     string
	Name        *string
	Description *string
}

func (t *Track) UpdateTrack(p *UpdateTrackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`trackId`, p.TrackId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}

	return t.restClient.Post(
		`/v2/Track/UseCase/UpdateTrack`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Track) UpdateTrackWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Track/UseCase/UpdateTrack`,
		data,
		nil,
		nil,
	)
}