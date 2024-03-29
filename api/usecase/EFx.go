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

type EFx struct {
	restClient rest.RestClientInterface
}

func NewEFx(restClient rest.RestClientInterface) *EFx {
	return &EFx{restClient}
}

// GET: Queries

type GetAdminConfigForEventParameters struct {
	EventId string
}

func (t *EFx) GetAdminConfigForEvent(p *GetAdminConfigForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/EFx/UseCase/GetAdminConfigForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetAllModulesForEventParameters struct {
	EventId             string
	WithDisabledModules *bool
}

func (t *EFx) GetAllModulesForEvent(p *GetAllModulesForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithDisabledModules != nil {
		queryParameters.Add(`withDisabledModules`, strconv.FormatBool(*p.WithDisabledModules))
	}

	return t.restClient.Get(
		`/v2/EFx/UseCase/GetAllModulesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetEFxScreenParameters struct {
	ScreenId string
}

func (t *EFx) GetEFxScreen(p *GetEFxScreenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`screenId`, p.ScreenId)

	return t.restClient.Get(
		`/v2/EFx/UseCase/GetEFxScreen`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetEFxStationParameters struct {
	StationId string
	WithData  *[]string // StackAndTicketType | EFxScreens
}

func (t *EFx) GetEFxStation(p *GetEFxStationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stationId`, p.StationId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/EFx/UseCase/GetEFxStation`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetInvitationIdByTagIdParameters struct {
	TagId   string
	EventId string
}

func (t *EFx) GetInvitationIdByTagId(p *GetInvitationIdByTagIdParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`tagId`, p.TagId)
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/EFx/UseCase/GetInvitationIdByTagId`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListEFxScreensForEventParameters struct {
	EventId          string
	SortBy           *string
	SortDirection    *string
	Page             *int64 // >= 1
	ItemsPerPage     *int64 // 1-100
	ScreenTypeFilter *[]string
}

func (t *EFx) ListEFxScreensForEvent(p *ListEFxScreensForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
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
	if p.ScreenTypeFilter != nil {
		for i := range *p.ScreenTypeFilter {
			queryParameters.Add(`screenTypeFilter[]`, (*p.ScreenTypeFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/EFx/UseCase/ListEFxScreensForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListEFxStationsForEventParameters struct {
	EventId       string
	WithData      *[]string // StackAndTicketType | EFxScreens
	SortBy        *string
	SortDirection *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
	ModuleFilter  *[]string
	Query         *string
}

func (t *EFx) ListEFxStationsForEvent(p *ListEFxStationsForEventParameters) (r *http.Response, err error) {
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
	if p.ModuleFilter != nil {
		for i := range *p.ModuleFilter {
			queryParameters.Add(`moduleFilter[]`, (*p.ModuleFilter)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}

	return t.restClient.Get(
		`/v2/EFx/UseCase/ListEFxStationsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type ActivateInvitationParameters struct {
	InvitationId string
	TagId        string
}

func (t *EFx) ActivateInvitation(p *ActivateInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`tagId`, p.TagId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/ActivateInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) ActivateInvitationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/ActivateInvitation`,
		data,
		nil,
		nil,
	)
}

type CreateEFxScreenParameters struct {
	EventId       string
	BackgroundUrl string
	TextColor     string
	ScreenType    string
	Text          string
	ScreenId      *string
}

func (t *EFx) CreateEFxScreen(p *CreateEFxScreenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`backgroundUrl`, p.BackgroundUrl)
	queryParameters.Add(`textColor`, p.TextColor)
	queryParameters.Add(`screenType`, p.ScreenType)
	queryParameters.Add(`text`, p.Text)
	if p.ScreenId != nil {
		queryParameters.Add(`screenId`, *p.ScreenId)
	}

	return t.restClient.Post(
		`/v2/EFx/UseCase/CreateEFxScreen`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) CreateEFxScreenWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/CreateEFxScreen`,
		data,
		nil,
		nil,
	)
}

type CreateEFxStationParameters struct {
	EventId     string
	StationName string
	ModuleType  string
	StationType string
	StationId   *string
}

func (t *EFx) CreateEFxStation(p *CreateEFxStationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`stationName`, p.StationName)
	queryParameters.Add(`moduleType`, p.ModuleType)
	queryParameters.Add(`stationType`, p.StationType)
	if p.StationId != nil {
		queryParameters.Add(`stationId`, *p.StationId)
	}

	return t.restClient.Post(
		`/v2/EFx/UseCase/CreateEFxStation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) CreateEFxStationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/CreateEFxStation`,
		data,
		nil,
		nil,
	)
}

type DisableForEventParameters struct {
	EventId string
}

func (t *EFx) DisableForEvent(p *DisableForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/DisableForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) DisableForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/DisableForEvent`,
		data,
		nil,
		nil,
	)
}

type DisableModuleForEventParameters struct {
	EventId    string
	ModuleType string
}

func (t *EFx) DisableModuleForEvent(p *DisableModuleForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`moduleType`, p.ModuleType)

	return t.restClient.Post(
		`/v2/EFx/UseCase/DisableModuleForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) DisableModuleForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/DisableModuleForEvent`,
		data,
		nil,
		nil,
	)
}

type DisableNFCForEventParameters struct {
	EventId string
}

func (t *EFx) DisableNFCForEvent(p *DisableNFCForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/DisableNFCForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) DisableNFCForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/DisableNFCForEvent`,
		data,
		nil,
		nil,
	)
}

type DisableSMSForEventParameters struct {
	EventId string
}

func (t *EFx) DisableSMSForEvent(p *DisableSMSForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/DisableSMSForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) DisableSMSForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/DisableSMSForEvent`,
		data,
		nil,
		nil,
	)
}

type EnableForEventParameters struct {
	EventId string
}

func (t *EFx) EnableForEvent(p *EnableForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/EnableForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) EnableForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/EnableForEvent`,
		data,
		nil,
		nil,
	)
}

type EnableModuleForEventParameters struct {
	EventId    string
	ModuleType string
}

func (t *EFx) EnableModuleForEvent(p *EnableModuleForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`moduleType`, p.ModuleType)

	return t.restClient.Post(
		`/v2/EFx/UseCase/EnableModuleForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) EnableModuleForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/EnableModuleForEvent`,
		data,
		nil,
		nil,
	)
}

type EnableNFCForEventParameters struct {
	EventId string
}

func (t *EFx) EnableNFCForEvent(p *EnableNFCForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/EnableNFCForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) EnableNFCForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/EnableNFCForEvent`,
		data,
		nil,
		nil,
	)
}

type EnableSMSForEventParameters struct {
	EventId string
}

func (t *EFx) EnableSMSForEvent(p *EnableSMSForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/EnableSMSForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) EnableSMSForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/EnableSMSForEvent`,
		data,
		nil,
		nil,
	)
}

type RemoveEFxStationParameters struct {
	StationId string
}

func (t *EFx) RemoveEFxStation(p *RemoveEFxStationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stationId`, p.StationId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/RemoveEFxStation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) RemoveEFxStationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/RemoveEFxStation`,
		data,
		nil,
		nil,
	)
}

type RequestForEventParameters struct {
	EventId             string
	UserId              string
	RequestedEFxModules *[]string // guest-management | access-control | athletes-bag | concierge | digital-memory-bank | guest-info | messaging | smsquiz | product-pickup | raffle | reservation | roaming-photographer | smart-bar | teams | lead-retrieval
}

func (t *EFx) RequestForEvent(p *RequestForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`userId`, p.UserId)
	if p.RequestedEFxModules != nil {
		for i := range *p.RequestedEFxModules {
			queryParameters.Add(`requestedEFxModules[]`, (*p.RequestedEFxModules)[i])
		}
	}

	return t.restClient.Post(
		`/v2/EFx/UseCase/RequestForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) RequestForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/RequestForEvent`,
		data,
		nil,
		nil,
	)
}

type SendSMSWithAppLinkParameters struct {
	PhoneNumber     string
	AppCategoryType *string
}

func (t *EFx) SendSMSWithAppLink(p *SendSMSWithAppLinkParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`phoneNumber`, p.PhoneNumber)
	if p.AppCategoryType != nil {
		queryParameters.Add(`appCategoryType`, *p.AppCategoryType)
	}

	return t.restClient.Post(
		`/v2/EFx/UseCase/SendSMSWithAppLink`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) SendSMSWithAppLinkWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/SendSMSWithAppLink`,
		data,
		nil,
		nil,
	)
}

type SetAdminPinForEventParameters struct {
	EventId string
	Pin     string
}

func (t *EFx) SetAdminPinForEvent(p *SetAdminPinForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`pin`, p.Pin)

	return t.restClient.Post(
		`/v2/EFx/UseCase/SetAdminPinForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) SetAdminPinForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/SetAdminPinForEvent`,
		data,
		nil,
		nil,
	)
}

type SetContentDeliveryForEFxStationParameters struct {
	StationId     string
	SmsText       *string
	SmsMedialUrl  *string
	EmailDesignId *string
}

func (t *EFx) SetContentDeliveryForEFxStation(p *SetContentDeliveryForEFxStationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stationId`, p.StationId)
	if p.SmsText != nil {
		queryParameters.Add(`smsText`, *p.SmsText)
	}
	if p.SmsMedialUrl != nil {
		queryParameters.Add(`smsMedialUrl`, *p.SmsMedialUrl)
	}
	if p.EmailDesignId != nil {
		queryParameters.Add(`emailDesignId`, *p.EmailDesignId)
	}

	return t.restClient.Post(
		`/v2/EFx/UseCase/SetContentDeliveryForEFxStation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) SetContentDeliveryForEFxStationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/SetContentDeliveryForEFxStation`,
		data,
		nil,
		nil,
	)
}

type SetSMSForEventParameters struct {
	EventId string
	Message string
}

func (t *EFx) SetSMSForEvent(p *SetSMSForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`message`, p.Message)

	return t.restClient.Post(
		`/v2/EFx/UseCase/SetSMSForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) SetSMSForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/SetSMSForEvent`,
		data,
		nil,
		nil,
	)
}

type SetSMSMediaUrlForEventParameters struct {
	EventId  string
	MediaUrl string
}

func (t *EFx) SetSMSMediaUrlForEvent(p *SetSMSMediaUrlForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`mediaUrl`, p.MediaUrl)

	return t.restClient.Post(
		`/v2/EFx/UseCase/SetSMSMediaUrlForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) SetSMSMediaUrlForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/SetSMSMediaUrlForEvent`,
		data,
		nil,
		nil,
	)
}

type SetScreensForEFxStationParameters struct {
	StationId string
	ScreenIds *[]string
}

func (t *EFx) SetScreensForEFxStation(p *SetScreensForEFxStationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stationId`, p.StationId)
	if p.ScreenIds != nil {
		for i := range *p.ScreenIds {
			queryParameters.Add(`screenIds[]`, (*p.ScreenIds)[i])
		}
	}

	return t.restClient.Post(
		`/v2/EFx/UseCase/SetScreensForEFxStation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) SetScreensForEFxStationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/SetScreensForEFxStation`,
		data,
		nil,
		nil,
	)
}

type SetStacksForEFxStationParameters struct {
	StationId string
	StackIds  *[]string
}

func (t *EFx) SetStacksForEFxStation(p *SetStacksForEFxStationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stationId`, p.StationId)
	if p.StackIds != nil {
		for i := range *p.StackIds {
			queryParameters.Add(`stackIds[]`, (*p.StackIds)[i])
		}
	}

	return t.restClient.Post(
		`/v2/EFx/UseCase/SetStacksForEFxStation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) SetStacksForEFxStationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/SetStacksForEFxStation`,
		data,
		nil,
		nil,
	)
}

type SetWelcomeEmailDesignForEventParameters struct {
	EventId       string
	EmailDesignId *string
}

func (t *EFx) SetWelcomeEmailDesignForEvent(p *SetWelcomeEmailDesignForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.EmailDesignId != nil {
		queryParameters.Add(`emailDesignId`, *p.EmailDesignId)
	}

	return t.restClient.Post(
		`/v2/EFx/UseCase/SetWelcomeEmailDesignForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) SetWelcomeEmailDesignForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/SetWelcomeEmailDesignForEvent`,
		data,
		nil,
		nil,
	)
}

type UpdateEFxScreenParameters struct {
	ScreenId      string
	Text          string
	TextColor     string
	BackgroundUrl string
}

func (t *EFx) UpdateEFxScreen(p *UpdateEFxScreenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`screenId`, p.ScreenId)
	queryParameters.Add(`text`, p.Text)
	queryParameters.Add(`textColor`, p.TextColor)
	queryParameters.Add(`backgroundUrl`, p.BackgroundUrl)

	return t.restClient.Post(
		`/v2/EFx/UseCase/UpdateEFxScreen`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) UpdateEFxScreenWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/UpdateEFxScreen`,
		data,
		nil,
		nil,
	)
}

type UpdateEFxStationParameters struct {
	StationId   string
	StationName string
	StationType string
}

func (t *EFx) UpdateEFxStation(p *UpdateEFxStationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stationId`, p.StationId)
	queryParameters.Add(`stationName`, p.StationName)
	queryParameters.Add(`stationType`, p.StationType)

	return t.restClient.Post(
		`/v2/EFx/UseCase/UpdateEFxStation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EFx) UpdateEFxStationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EFx/UseCase/UpdateEFxStation`,
		data,
		nil,
		nil,
	)
}
