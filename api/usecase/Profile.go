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

type Profile struct {
	restClient rest.RestClientInterface
}

func NewProfile(restClient rest.RestClientInterface) *Profile {
	return &Profile{restClient}
}

// GET: Queries

type GetProfileParameters struct {
	ProfileId string
}

func (t *Profile) GetProfile(p *GetProfileParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`profileId`, p.ProfileId)

	return t.restClient.Get(
		`/v2/Profile/UseCase/GetProfile`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListProfilesForEventParameters struct {
	EventId           string
	Page              *int64 // >= 1
	ItemsPerPage      *int64 // 1-500
	SortBy            *string
	SortDirection     *string
	Query             *string
	ShouldHideDeleted *bool
}

func (t *Profile) ListProfilesForEvent(p *ListProfilesForEventParameters) (r *http.Response, err error) {
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
		`/v2/Profile/UseCase/ListProfilesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListProfilesForPoolParameters struct {
	PoolId        string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
	SortBy        *string
	SortDirection *string
}

func (t *Profile) ListProfilesForPool(p *ListProfilesForPoolParameters) (r *http.Response, err error) {
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
		`/v2/Profile/UseCase/ListProfilesForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddProfileToEventParameters struct {
	ProfileId        string
	EventId          string
	EventProfileType string
}

func (t *Profile) AddProfileToEvent(p *AddProfileToEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`profileId`, p.ProfileId)
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`eventProfileType`, p.EventProfileType)

	return t.restClient.Post(
		`/v2/Profile/UseCase/AddProfileToEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) AddProfileToEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/AddProfileToEvent`,
		data,
		nil,
		nil,
	)
}

type AddProfilesToEventParameters struct {
	ProfileIds       []string
	EventId          string
	EventProfileType string
}

func (t *Profile) AddProfilesToEvent(p *AddProfilesToEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range p.ProfileIds {
		queryParameters.Add(`profileIds[]`, p.ProfileIds[i])
	}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`eventProfileType`, p.EventProfileType)

	return t.restClient.Post(
		`/v2/Profile/UseCase/AddProfilesToEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) AddProfilesToEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/AddProfilesToEvent`,
		data,
		nil,
		nil,
	)
}

type CreateProfileParameters struct {
	PoolId           string
	FirstName        *string
	LastName         *string
	Company          *string
	EmailAddress     *string
	Title            *string
	Description      *string
	ImageUrl         *string
	ProfileType      *string
	ProfileId        *string
	EventId          *string
	EventProfileType *string
}

func (t *Profile) CreateProfile(p *CreateProfileParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.FirstName != nil {
		queryParameters.Add(`firstName`, *p.FirstName)
	}
	if p.LastName != nil {
		queryParameters.Add(`lastName`, *p.LastName)
	}
	if p.Company != nil {
		queryParameters.Add(`company`, *p.Company)
	}
	if p.EmailAddress != nil {
		queryParameters.Add(`emailAddress`, *p.EmailAddress)
	}
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}
	if p.ImageUrl != nil {
		queryParameters.Add(`imageUrl`, *p.ImageUrl)
	}
	if p.ProfileType != nil {
		queryParameters.Add(`profileType`, *p.ProfileType)
	}
	if p.ProfileId != nil {
		queryParameters.Add(`profileId`, *p.ProfileId)
	}
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}
	if p.EventProfileType != nil {
		queryParameters.Add(`eventProfileType`, *p.EventProfileType)
	}

	return t.restClient.Post(
		`/v2/Profile/UseCase/CreateProfile`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) CreateProfileWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/CreateProfile`,
		data,
		nil,
		nil,
	)
}

type DeleteProfileParameters struct {
	ProfileId string
}

func (t *Profile) DeleteProfile(p *DeleteProfileParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`profileId`, p.ProfileId)

	return t.restClient.Post(
		`/v2/Profile/UseCase/DeleteProfile`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) DeleteProfileWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/DeleteProfile`,
		data,
		nil,
		nil,
	)
}

type RemoveProfileParameters struct {
	ProfileId string
}

func (t *Profile) RemoveProfile(p *RemoveProfileParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`profileId`, p.ProfileId)

	return t.restClient.Post(
		`/v2/Profile/UseCase/RemoveProfile`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) RemoveProfileWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/RemoveProfile`,
		data,
		nil,
		nil,
	)
}

type RemoveProfileFromEventParameters struct {
	ProfileId string
	EventId   string
}

func (t *Profile) RemoveProfileFromEvent(p *RemoveProfileFromEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`profileId`, p.ProfileId)
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Profile/UseCase/RemoveProfileFromEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) RemoveProfileFromEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/RemoveProfileFromEvent`,
		data,
		nil,
		nil,
	)
}

type SetCompanyParameters struct {
	ProfileId string
	Company   string
}

func (t *Profile) SetCompany(p *SetCompanyParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`profileId`, p.ProfileId)
	queryParameters.Add(`company`, p.Company)

	return t.restClient.Post(
		`/v2/Profile/UseCase/SetCompany`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) SetCompanyWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/SetCompany`,
		data,
		nil,
		nil,
	)
}

type SetDescriptionParameters struct {
	ProfileId   string
	Description string
}

func (t *Profile) SetDescription(p *SetDescriptionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`profileId`, p.ProfileId)
	queryParameters.Add(`description`, p.Description)

	return t.restClient.Post(
		`/v2/Profile/UseCase/SetDescription`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) SetDescriptionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/SetDescription`,
		data,
		nil,
		nil,
	)
}

type SetNameParameters struct {
	ProfileId string
	FirstName *string
	LastName  *string
}

func (t *Profile) SetName(p *SetNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`profileId`, p.ProfileId)
	if p.FirstName != nil {
		queryParameters.Add(`firstName`, *p.FirstName)
	}
	if p.LastName != nil {
		queryParameters.Add(`lastName`, *p.LastName)
	}

	return t.restClient.Post(
		`/v2/Profile/UseCase/SetName`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) SetNameWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/SetName`,
		data,
		nil,
		nil,
	)
}

type SetProfileTypeParameters struct {
	ProfileId   string
	ProfileType string
}

func (t *Profile) SetProfileType(p *SetProfileTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`profileId`, p.ProfileId)
	queryParameters.Add(`profileType`, p.ProfileType)

	return t.restClient.Post(
		`/v2/Profile/UseCase/SetProfileType`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) SetProfileTypeWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/SetProfileType`,
		data,
		nil,
		nil,
	)
}

type UpdateEventProfileTypeParameters struct {
	EventProfileId   string
	EventProfileType string
}

func (t *Profile) UpdateEventProfileType(p *UpdateEventProfileTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventProfileId`, p.EventProfileId)
	queryParameters.Add(`eventProfileType`, p.EventProfileType)

	return t.restClient.Post(
		`/v2/Profile/UseCase/UpdateEventProfileType`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) UpdateEventProfileTypeWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/UpdateEventProfileType`,
		data,
		nil,
		nil,
	)
}

type UpdateProfileParameters struct {
	ProfileId    string
	PoolId       *string
	FirstName    *string
	LastName     *string
	Company      *string
	EmailAddress *string
	Title        *string
	Description  *string
	ImageUrl     *string
	ProfileType  *string
}

func (t *Profile) UpdateProfile(p *UpdateProfileParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`profileId`, p.ProfileId)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.FirstName != nil {
		queryParameters.Add(`firstName`, *p.FirstName)
	}
	if p.LastName != nil {
		queryParameters.Add(`lastName`, *p.LastName)
	}
	if p.Company != nil {
		queryParameters.Add(`company`, *p.Company)
	}
	if p.EmailAddress != nil {
		queryParameters.Add(`emailAddress`, *p.EmailAddress)
	}
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}
	if p.ImageUrl != nil {
		queryParameters.Add(`imageUrl`, *p.ImageUrl)
	}
	if p.ProfileType != nil {
		queryParameters.Add(`profileType`, *p.ProfileType)
	}

	return t.restClient.Post(
		`/v2/Profile/UseCase/UpdateProfile`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Profile) UpdateProfileWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Profile/UseCase/UpdateProfile`,
		data,
		nil,
		nil,
	)
}
