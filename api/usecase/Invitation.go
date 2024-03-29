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

type Invitation struct {
	restClient rest.RestClientInterface
}

func NewInvitation(restClient rest.RestClientInterface) *Invitation {
	return &Invitation{restClient}
}

// GET: Queries

type CheckIfInvitationExistsForEventFromClearJSONParameters struct {
	EventId     string
	FirstName   string
	LastName    string
	ClearUserId string
	Score       string
}

func (t *Invitation) CheckIfInvitationExistsForEventFromClearJSON(p *CheckIfInvitationExistsForEventFromClearJSONParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`firstName`, p.FirstName)
	queryParameters.Add(`lastName`, p.LastName)
	queryParameters.Add(`clearUserId`, p.ClearUserId)
	queryParameters.Add(`score`, p.Score)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/CheckIfInvitationExistsForEventFromClearJSON`,
		&queryParameters,
		nil,
		nil,
	)
}

type CheckIfInvitationExistsForEventFromClearQRCodeParameters struct {
	EventId string
	ClearId string
}

func (t *Invitation) CheckIfInvitationExistsForEventFromClearQRCode(p *CheckIfInvitationExistsForEventFromClearQRCodeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`clearId`, p.ClearId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/CheckIfInvitationExistsForEventFromClearQRCode`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetAllInvitationStatusTypesForStackMethodTypeParameters struct {
	StackMethodType string
}

func (t *Invitation) GetAllInvitationStatusTypesForStackMethodType(p *GetAllInvitationStatusTypesForStackMethodTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackMethodType`, p.StackMethodType)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetAllInvitationStatusTypesForStackMethodType`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetCheckInCountsForEventParameters struct {
	EventId string
}

func (t *Invitation) GetCheckInCountsForEvent(p *GetCheckInCountsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetCheckInCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetCheckInCountsForTicketBlockParameters struct {
	TicketBlockId string
}

func (t *Invitation) GetCheckInCountsForTicketBlock(p *GetCheckInCountsForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetCheckInCountsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetInvitationParameters struct {
	InvitationId       string
	WithData           *[]string // UserHealthPass | Event | UserName | User | UserIdentifier | Stack | TicketType | QuestionResponse | Answer | Purchase | DayPassAvailabilityCounts | EFxActivationStatus | RelatedInvitation
	WithUserAttributes *[]string
}

func (t *Invitation) GetInvitation(p *GetInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetInvitationCountsForEventParameters struct {
	EventId string
}

func (t *Invitation) GetInvitationCountsForEvent(p *GetInvitationCountsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetInvitationCountsForTicketBlockParameters struct {
	TicketBlockId string
}

func (t *Invitation) GetInvitationCountsForTicketBlock(p *GetInvitationCountsForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationCountsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetInvitationCountsForUserParameters struct {
	UserId string
	PoolId *string
}

func (t *Invitation) GetInvitationCountsForUser(p *GetInvitationCountsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationCountsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetInvitationLastActionCountsForEventParameters struct {
	EventId string
}

func (t *Invitation) GetInvitationLastActionCountsForEvent(p *GetInvitationLastActionCountsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationLastActionCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetInvitationStatusTypeCountsForEventParameters struct {
	EventId string
}

func (t *Invitation) GetInvitationStatusTypeCountsForEvent(p *GetInvitationStatusTypeCountsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationStatusTypeCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetInvitationStatusTypeCountsForTicketBlockParameters struct {
	TicketBlockId string
}

func (t *Invitation) GetInvitationStatusTypeCountsForTicketBlock(p *GetInvitationStatusTypeCountsForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationStatusTypeCountsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForEventParameters struct {
	EventId                      string
	WithData                     *[]string // UserHealthPasses | UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt | DayPassAvailabilityCounts | RelatedInvitation
	WithUserAttributes           *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom | virbela | healthpass
	Query                        *string
	StatusFilter                 *[]string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	LastModifiedTimestamp        *int64
	IsCheckedIn                  *bool
	SortBy                       *string
	SortDirection                *string
	Page                         *int64    // >= 1
	ItemsPerPage                 *int64    // 1-250
	HealthPassScoreFilter        *[]string // green | red | amber | unknown
	ExcludeHealthPassScoreFilter *[]string // green | red | amber | unknown
}

func (t *Invitation) ListInvitationsForEvent(p *ListInvitationsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
	if p.HealthPassScoreFilter != nil {
		for i := range *p.HealthPassScoreFilter {
			queryParameters.Add(`healthPassScoreFilter[]`, (*p.HealthPassScoreFilter)[i])
		}
	}
	if p.ExcludeHealthPassScoreFilter != nil {
		for i := range *p.ExcludeHealthPassScoreFilter {
			queryParameters.Add(`excludeHealthPassScoreFilter[]`, (*p.ExcludeHealthPassScoreFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForStackParameters struct {
	EventId               string
	StackId               string
	WithData              *[]string // UserHealthPasses | UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt | DayPassAvailabilityCounts | RelatedInvitation
	WithUserAttributes    *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom | virbela | healthpass
	Query                 *string
	StatusFilter          *[]string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	LastModifiedTimestamp *int64
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-250
}

func (t *Invitation) ListInvitationsForStack(p *ListInvitationsForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`stackId`, p.StackId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
		`/v2/Invitation/UseCase/ListInvitationsForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForTicketBlockParameters struct {
	TicketBlockId         string
	WithData              *[]string // UserHealthPasses | UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt | DayPassAvailabilityCounts | RelatedInvitation
	WithUserAttributes    *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom | virbela | healthpass
	Query                 *string
	StatusFilter          *[]string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	LastModifiedTimestamp *int64
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-250
}

func (t *Invitation) ListInvitationsForTicketBlock(p *ListInvitationsForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
		`/v2/Invitation/UseCase/ListInvitationsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForTransactionParameters struct {
	TransactionId         string
	WithData              *[]string // UserHealthPasses | UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt | DayPassAvailabilityCounts | RelatedInvitation
	WithUserAttributes    *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom | virbela | healthpass
	Query                 *string
	StatusFilter          *[]string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	LastModifiedTimestamp *int64
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-250
}

func (t *Invitation) ListInvitationsForTransaction(p *ListInvitationsForTransactionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`transactionId`, p.TransactionId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
		`/v2/Invitation/UseCase/ListInvitationsForTransaction`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForUserParameters struct {
	UserId                string
	PoolId                *string
	EventId               *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-250
	EventDateFilterType   *string
	SortDirection         *string
	WithData              *[]string // Event | EventWithTags | Stack | StackAndTicketType
	StatusFilter          *[]string
	EventAttributesFilter *[]string
}

func (t *Invitation) ListInvitationsForUser(p *ListInvitationsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.EventAttributesFilter != nil {
		for i := range *p.EventAttributesFilter {
			queryParameters.Add(`eventAttributesFilter[]`, (*p.EventAttributesFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForUserByEmailParameters struct {
	Email               string
	PoolId              *string
	EventId             *string
	Page                *int64 // >= 1
	ItemsPerPage        *int64 // 1-250
	EventDateFilterType *string
	SortDirection       *string
	WithData            *[]string // Event | Stack | StackAndTicketType
	StatusFilter        *[]string
}

func (t *Invitation) ListInvitationsForUserByEmail(p *ListInvitationsForUserByEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, p.Email)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForUserByEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForUserByEmailForEventsParameters struct {
	Email               string
	EventIds            []string
	PoolId              string
	Page                *int64 // >= 1
	ItemsPerPage        *int64 // 1-250
	EventDateFilterType *string
	SortDirection       *string
	WithData            *[]string // Event | Stack | StackAndTicketType
	StatusFilter        *[]string
}

func (t *Invitation) ListInvitationsForUserByEmailForEvents(p *ListInvitationsForUserByEmailForEventsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, p.Email)
	for i := range p.EventIds {
		queryParameters.Add(`eventIds[]`, p.EventIds[i])
	}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForUserByEmailForEvents`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForUserForParentParameters struct {
	UserId              string
	ParentEventId       string
	PoolId              *string
	Page                *int64 // >= 1
	ItemsPerPage        *int64 // 1-500
	EventDateFilterType *string
	SortDirection       *string
	WithData            *[]string // Event | Stack | Venue | SessionTracks
	StatusFilter        *[]string
	ExcludeParent       *bool
}

func (t *Invitation) ListInvitationsForUserForParent(p *ListInvitationsForUserForParentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`parentEventId`, p.ParentEventId)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.ExcludeParent != nil {
		queryParameters.Add(`excludeParent`, strconv.FormatBool(*p.ExcludeParent))
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForUserForParent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListRelatedInvitationsParameters struct {
	RelatedInvitationId          string
	WithData                     *[]string // UserHealthPasses | UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt | DayPassAvailabilityCounts | RelatedInvitation
	WithUserAttributes           *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom | virbela | healthpass
	Query                        *string
	StatusFilter                 *[]string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	LastModifiedTimestamp        *int64
	IsCheckedIn                  *bool
	SortBy                       *string
	SortDirection                *string
	Page                         *int64    // >= 1
	ItemsPerPage                 *int64    // 1-250
	HealthPassScoreFilter        *[]string // green | red | amber | unknown
	ExcludeHealthPassScoreFilter *[]string // green | red | amber | unknown
}

func (t *Invitation) ListRelatedInvitations(p *ListRelatedInvitationsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`relatedInvitationId`, p.RelatedInvitationId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
	if p.HealthPassScoreFilter != nil {
		for i := range *p.HealthPassScoreFilter {
			queryParameters.Add(`healthPassScoreFilter[]`, (*p.HealthPassScoreFilter)[i])
		}
	}
	if p.ExcludeHealthPassScoreFilter != nil {
		for i := range *p.ExcludeHealthPassScoreFilter {
			queryParameters.Add(`excludeHealthPassScoreFilter[]`, (*p.ExcludeHealthPassScoreFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListRelatedInvitations`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListWaitlistForEventParameters struct {
	EventId               string
	WithData              *[]string // UserHealthPasses | UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt | DayPassAvailabilityCounts | RelatedInvitation
	WithUserAttributes    *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom | virbela | healthpass
	Query                 *string
	LastModifiedTimestamp *int64
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-100
}

func (t *Invitation) ListWaitlistForEvent(p *ListWaitlistForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
		`/v2/Invitation/UseCase/ListWaitlistForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListWaitlistForStackParameters struct {
	EventId               string
	StackId               string
	WithData              *[]string // UserHealthPasses | UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt | DayPassAvailabilityCounts | RelatedInvitation
	WithUserAttributes    *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom | virbela | healthpass
	Query                 *string
	StatusFilter          *[]string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	LastModifiedTimestamp *int64
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-100
}

func (t *Invitation) ListWaitlistForStack(p *ListWaitlistForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`stackId`, p.StackId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
		`/v2/Invitation/UseCase/ListWaitlistForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddInvitationToWaitlistParameters struct {
	InvitationId string
}

func (t *Invitation) AddInvitationToWaitlist(p *AddInvitationToWaitlistParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/AddInvitationToWaitlist`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) AddInvitationToWaitlistWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/AddInvitationToWaitlist`,
		data,
		nil,
		nil,
	)
}

type ChangeInvitationStatusParameters struct {
	InvitationId     string
	InvitationStatus string
}

func (t *Invitation) ChangeInvitationStatus(p *ChangeInvitationStatusParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`invitationStatus`, p.InvitationStatus)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/ChangeInvitationStatus`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) ChangeInvitationStatusWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/ChangeInvitationStatus`,
		data,
		nil,
		nil,
	)
}

type ChangeInviteCountParameters struct {
	InvitationId string
	InviteCount  *int64
}

func (t *Invitation) ChangeInviteCount(p *ChangeInviteCountParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.InviteCount != nil {
		queryParameters.Add(`inviteCount`, strconv.FormatInt(*p.InviteCount, 10))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/ChangeInviteCount`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) ChangeInviteCountWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/ChangeInviteCount`,
		data,
		nil,
		nil,
	)
}

type CheckInParameters struct {
	InvitationId string
	CheckInAt    *int64
	IsWebCheckIn *bool
}

func (t *Invitation) CheckIn(p *CheckInParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.CheckInAt != nil {
		queryParameters.Add(`checkInAt`, strconv.FormatInt(*p.CheckInAt, 10))
	}
	if p.IsWebCheckIn != nil {
		queryParameters.Add(`isWebCheckIn`, strconv.FormatBool(*p.IsWebCheckIn))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CheckIn`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) CheckInWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/CheckIn`,
		data,
		nil,
		nil,
	)
}

type CheckInWebConferenceRegistrantParameters struct {
	RegistrantId string
}

func (t *Invitation) CheckInWebConferenceRegistrant(p *CheckInWebConferenceRegistrantParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`registrantId`, p.RegistrantId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CheckInWebConferenceRegistrant`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) CheckInWebConferenceRegistrantWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/CheckInWebConferenceRegistrant`,
		data,
		nil,
		nil,
	)
}

type CreateInvitationParameters struct {
	EventId                   string
	StackId                   string
	InvitationStatus          string
	InviteSource              string
	IsCheckedIn               bool
	InviteCount               int64 // >= 1
	Email                     *string
	FirstName                 *string
	LastName                  *string
	Company                   *string
	Position                  *string
	CheckInNotes              *string
	InvitationId              *string
	ShouldSendInvitation      *bool
	InvitationNotes           *string
	Title                     *string
	Telephone                 *string
	Other                     *string
	CreatedTime               *int64
	ForceDuplicateInvitations *bool
	RelatedInvitationId       *string
	ExternalId                *string
}

func (t *Invitation) CreateInvitation(p *CreateInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`invitationStatus`, p.InvitationStatus)
	queryParameters.Add(`inviteSource`, p.InviteSource)
	queryParameters.Add(`isCheckedIn`, strconv.FormatBool(p.IsCheckedIn))
	queryParameters.Add(`inviteCount`, strconv.FormatInt(p.InviteCount, 10))
	if p.Email != nil {
		queryParameters.Add(`email`, *p.Email)
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
	if p.Position != nil {
		queryParameters.Add(`position`, *p.Position)
	}
	if p.CheckInNotes != nil {
		queryParameters.Add(`checkInNotes`, *p.CheckInNotes)
	}
	if p.InvitationId != nil {
		queryParameters.Add(`invitationId`, *p.InvitationId)
	}
	if p.ShouldSendInvitation != nil {
		queryParameters.Add(`shouldSendInvitation`, strconv.FormatBool(*p.ShouldSendInvitation))
	}
	if p.InvitationNotes != nil {
		queryParameters.Add(`invitationNotes`, *p.InvitationNotes)
	}
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.Telephone != nil {
		queryParameters.Add(`telephone`, *p.Telephone)
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
	}
	if p.CreatedTime != nil {
		queryParameters.Add(`createdTime`, strconv.FormatInt(*p.CreatedTime, 10))
	}
	if p.ForceDuplicateInvitations != nil {
		queryParameters.Add(`forceDuplicateInvitations`, strconv.FormatBool(*p.ForceDuplicateInvitations))
	}
	if p.RelatedInvitationId != nil {
		queryParameters.Add(`relatedInvitationId`, *p.RelatedInvitationId)
	}
	if p.ExternalId != nil {
		queryParameters.Add(`externalId`, *p.ExternalId)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) CreateInvitationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/CreateInvitation`,
		data,
		nil,
		nil,
	)
}

type CreateInvitationForTicketBlockParameters struct {
	EventId                      string
	StackId                      string
	TicketBlockId                string
	InvitationStatus             string
	InviteSource                 string
	IsCheckedIn                  bool
	InviteCount                  int64 // >= 1
	Email                        *string
	FirstName                    *string
	LastName                     *string
	Company                      *string
	Position                     *string
	CheckInNotes                 *string
	InvitationId                 *string
	ShouldSendInvitation         *bool
	InvitationNotes              *string
	Title                        *string
	Telephone                    *string
	Other                        *string
	ShouldCreateParentInvitation *bool
}

func (t *Invitation) CreateInvitationForTicketBlock(p *CreateInvitationForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`invitationStatus`, p.InvitationStatus)
	queryParameters.Add(`inviteSource`, p.InviteSource)
	queryParameters.Add(`isCheckedIn`, strconv.FormatBool(p.IsCheckedIn))
	queryParameters.Add(`inviteCount`, strconv.FormatInt(p.InviteCount, 10))
	if p.Email != nil {
		queryParameters.Add(`email`, *p.Email)
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
	if p.Position != nil {
		queryParameters.Add(`position`, *p.Position)
	}
	if p.CheckInNotes != nil {
		queryParameters.Add(`checkInNotes`, *p.CheckInNotes)
	}
	if p.InvitationId != nil {
		queryParameters.Add(`invitationId`, *p.InvitationId)
	}
	if p.ShouldSendInvitation != nil {
		queryParameters.Add(`shouldSendInvitation`, strconv.FormatBool(*p.ShouldSendInvitation))
	}
	if p.InvitationNotes != nil {
		queryParameters.Add(`invitationNotes`, *p.InvitationNotes)
	}
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.Telephone != nil {
		queryParameters.Add(`telephone`, *p.Telephone)
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
	}
	if p.ShouldCreateParentInvitation != nil {
		queryParameters.Add(`shouldCreateParentInvitation`, strconv.FormatBool(*p.ShouldCreateParentInvitation))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) CreateInvitationForTicketBlockWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/CreateInvitationForTicketBlock`,
		data,
		nil,
		nil,
	)
}

type CreateInvitationForUserParameters struct {
	UserId               string
	GuestsPerInvitation  int64 // >= 1
	StackId              string
	InvitationStatusType string
	ShouldSendEmail      *bool
	TicketBlockId        *string
}

func (t *Invitation) CreateInvitationForUser(p *CreateInvitationForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`guestsPerInvitation`, strconv.FormatInt(p.GuestsPerInvitation, 10))
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`invitationStatusType`, p.InvitationStatusType)
	if p.ShouldSendEmail != nil {
		queryParameters.Add(`shouldSendEmail`, strconv.FormatBool(*p.ShouldSendEmail))
	}
	if p.TicketBlockId != nil {
		queryParameters.Add(`ticketBlockId`, *p.TicketBlockId)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) CreateInvitationForUserWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/CreateInvitationForUser`,
		data,
		nil,
		nil,
	)
}

type CreateInvitationsFromGroupParameters struct {
	GroupId                string
	StackId                string
	GuestsPerInvitation    int64 // >= 1
	InvitationCreationType string
}

func (t *Invitation) CreateInvitationsFromGroup(p *CreateInvitationsFromGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`guestsPerInvitation`, strconv.FormatInt(p.GuestsPerInvitation, 10))
	queryParameters.Add(`invitationCreationType`, p.InvitationCreationType)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) CreateInvitationsFromGroupWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroup`,
		data,
		nil,
		nil,
	)
}

type CreateInvitationsFromGroupForCIOEventParameters struct {
	EventId             string
	GroupId             string
	GuestsPerInvitation int64 // >= 1
	StackId             *string
}

func (t *Invitation) CreateInvitationsFromGroupForCIOEvent(p *CreateInvitationsFromGroupForCIOEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`guestsPerInvitation`, strconv.FormatInt(p.GuestsPerInvitation, 10))
	if p.StackId != nil {
		queryParameters.Add(`stackId`, *p.StackId)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroupForCIOEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) CreateInvitationsFromGroupForCIOEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroupForCIOEvent`,
		data,
		nil,
		nil,
	)
}

type CreateInvitationsFromGroupForTicketBlockParameters struct {
	GroupId                string
	TicketBlockId          string
	StackId                string
	GuestsPerInvitation    int64 // >= 1
	InvitationCreationType string
}

func (t *Invitation) CreateInvitationsFromGroupForTicketBlock(p *CreateInvitationsFromGroupForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`guestsPerInvitation`, strconv.FormatInt(p.GuestsPerInvitation, 10))
	queryParameters.Add(`invitationCreationType`, p.InvitationCreationType)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroupForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) CreateInvitationsFromGroupForTicketBlockWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroupForTicketBlock`,
		data,
		nil,
		nil,
	)
}

type CreateSessionInvitationParameters struct {
	ParentEventInvitationId string
	SessionId               string
	InvitationId            *string
}

func (t *Invitation) CreateSessionInvitation(p *CreateSessionInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`parentEventInvitationId`, p.ParentEventInvitationId)
	queryParameters.Add(`sessionId`, p.SessionId)
	if p.InvitationId != nil {
		queryParameters.Add(`invitationId`, *p.InvitationId)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateSessionInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) CreateSessionInvitationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/CreateSessionInvitation`,
		data,
		nil,
		nil,
	)
}

type CreateWebhookParameters struct {
	EventId       string
	WebhookType   string
	WebhookMethod string
	Url           string
}

func (t *Invitation) CreateWebhook(p *CreateWebhookParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`webhookType`, p.WebhookType)
	queryParameters.Add(`webhookMethod`, p.WebhookMethod)
	queryParameters.Add(`url`, p.Url)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateWebhook`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) CreateWebhookWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/CreateWebhook`,
		data,
		nil,
		nil,
	)
}

type DeleteWebhookParameters struct {
	WebhookId string
}

func (t *Invitation) DeleteWebhook(p *DeleteWebhookParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`webhookId`, p.WebhookId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/DeleteWebhook`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) DeleteWebhookWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/DeleteWebhook`,
		data,
		nil,
		nil,
	)
}

type DisableArrivalAlertParameters struct {
	InvitationId string
}

func (t *Invitation) DisableArrivalAlert(p *DisableArrivalAlertParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/DisableArrivalAlert`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) DisableArrivalAlertWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/DisableArrivalAlert`,
		data,
		nil,
		nil,
	)
}

type PromoteInvitationsFromWaitlistParameters struct {
	InvitationIds   []string
	IsConfirmed     bool
	ShouldSendEmail *bool
}

func (t *Invitation) PromoteInvitationsFromWaitlist(p *PromoteInvitationsFromWaitlistParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range p.InvitationIds {
		queryParameters.Add(`invitationIds[]`, p.InvitationIds[i])
	}
	queryParameters.Add(`isConfirmed`, strconv.FormatBool(p.IsConfirmed))
	if p.ShouldSendEmail != nil {
		queryParameters.Add(`shouldSendEmail`, strconv.FormatBool(*p.ShouldSendEmail))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/PromoteInvitationsFromWaitlist`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) PromoteInvitationsFromWaitlistWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/PromoteInvitationsFromWaitlist`,
		data,
		nil,
		nil,
	)
}

type RescindAllInvitationsParameters struct {
	EventId string
}

func (t *Invitation) RescindAllInvitations(p *RescindAllInvitationsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/RescindAllInvitations`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) RescindAllInvitationsWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/RescindAllInvitations`,
		data,
		nil,
		nil,
	)
}

type RescindInvitationParameters struct {
	InvitationId string
}

func (t *Invitation) RescindInvitation(p *RescindInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/RescindInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) RescindInvitationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/RescindInvitation`,
		data,
		nil,
		nil,
	)
}

type RescindInvitationForSessionParameters struct {
	InvitationId string
}

func (t *Invitation) RescindInvitationForSession(p *RescindInvitationForSessionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/RescindInvitationForSession`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) RescindInvitationForSessionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/RescindInvitationForSession`,
		data,
		nil,
		nil,
	)
}

type ResendAllInvitationEmailsParameters struct {
	EventId  string
	DayCount int64 // 0-90
}

func (t *Invitation) ResendAllInvitationEmails(p *ResendAllInvitationEmailsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`dayCount`, strconv.FormatInt(p.DayCount, 10))

	return t.restClient.Post(
		`/v2/Invitation/UseCase/ResendAllInvitationEmails`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) ResendAllInvitationEmailsWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/ResendAllInvitationEmails`,
		data,
		nil,
		nil,
	)
}

type ResendAllTicketBlockInvitationEmailsParameters struct {
	TicketBlockId string
}

func (t *Invitation) ResendAllTicketBlockInvitationEmails(p *ResendAllTicketBlockInvitationEmailsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/ResendAllTicketBlockInvitationEmails`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) ResendAllTicketBlockInvitationEmailsWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/ResendAllTicketBlockInvitationEmails`,
		data,
		nil,
		nil,
	)
}

type ResendConfirmationEmailParameters struct {
	InvitationId string
}

func (t *Invitation) ResendConfirmationEmail(p *ResendConfirmationEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/ResendConfirmationEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) ResendConfirmationEmailWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/ResendConfirmationEmail`,
		data,
		nil,
		nil,
	)
}

type ResendInvitationEmailParameters struct {
	InvitationId string
}

func (t *Invitation) ResendInvitationEmail(p *ResendInvitationEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/ResendInvitationEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) ResendInvitationEmailWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/ResendInvitationEmail`,
		data,
		nil,
		nil,
	)
}

type SetAllQuestionResponsesParameters struct {
	InvitationId                       string
	QuestionId                         string
	QuestionResponseIdsWithAnswersJson string
}

func (t *Invitation) SetAllQuestionResponses(p *SetAllQuestionResponsesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`questionResponseIdsWithAnswersJson`, p.QuestionResponseIdsWithAnswersJson)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetAllQuestionResponses`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) SetAllQuestionResponsesWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/SetAllQuestionResponses`,
		data,
		nil,
		nil,
	)
}

type SetArrivalAlertEmailParameters struct {
	InvitationId           string
	ToEmail                string
	CcEmails               *[]string
	ShouldSendArrivalAlert *bool
}

func (t *Invitation) SetArrivalAlertEmail(p *SetArrivalAlertEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`toEmail`, p.ToEmail)
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.ShouldSendArrivalAlert != nil {
		queryParameters.Add(`shouldSendArrivalAlert`, strconv.FormatBool(*p.ShouldSendArrivalAlert))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetArrivalAlertEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) SetArrivalAlertEmailWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/SetArrivalAlertEmail`,
		data,
		nil,
		nil,
	)
}

type SetCheckInNotesParameters struct {
	InvitationId string
	CheckInNotes *string
}

func (t *Invitation) SetCheckInNotes(p *SetCheckInNotesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.CheckInNotes != nil {
		queryParameters.Add(`checkInNotes`, *p.CheckInNotes)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetCheckInNotes`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) SetCheckInNotesWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/SetCheckInNotes`,
		data,
		nil,
		nil,
	)
}

type SetDayPassCountParameters struct {
	InvitationId string
	DayPassCount *int64 // >= -1
}

func (t *Invitation) SetDayPassCount(p *SetDayPassCountParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.DayPassCount != nil {
		queryParameters.Add(`dayPassCount`, strconv.FormatInt(*p.DayPassCount, 10))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetDayPassCount`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) SetDayPassCountWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/SetDayPassCount`,
		data,
		nil,
		nil,
	)
}

type SetInvitationNotesParameters struct {
	InvitationId    string
	InvitationNotes *string
}

func (t *Invitation) SetInvitationNotes(p *SetInvitationNotesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.InvitationNotes != nil {
		queryParameters.Add(`invitationNotes`, *p.InvitationNotes)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetInvitationNotes`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) SetInvitationNotesWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/SetInvitationNotes`,
		data,
		nil,
		nil,
	)
}

type SetQuestionResponseParameters struct {
	InvitationId string
	QuestionId   string
	AnswerIds    *[]string
	Text         *string
}

func (t *Invitation) SetQuestionResponse(p *SetQuestionResponseParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`questionId`, p.QuestionId)
	if p.AnswerIds != nil {
		for i := range *p.AnswerIds {
			queryParameters.Add(`answerIds[]`, (*p.AnswerIds)[i])
		}
	}
	if p.Text != nil {
		queryParameters.Add(`text`, *p.Text)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetQuestionResponse`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) SetQuestionResponseWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/SetQuestionResponse`,
		data,
		nil,
		nil,
	)
}

type UndoCheckInParameters struct {
	InvitationId string
}

func (t *Invitation) UndoCheckIn(p *UndoCheckInParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/UndoCheckIn`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) UndoCheckInWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/UndoCheckIn`,
		data,
		nil,
		nil,
	)
}

type UpdateInvitationParameters struct {
	InvitationId              string
	StackId                   string
	InvitationStatus          string
	Company                   *string
	Position                  *string
	Email                     *string
	FirstName                 *string
	LastName                  *string
	Other                     *string
	Telephone                 *string
	UpdatedTime               *int64
	ForceDuplicateInvitations *bool
	InviteCount               *int64
	Title                     *string
	CheckInNotes              *string
	RelatedInvitationId       *string
}

func (t *Invitation) UpdateInvitation(p *UpdateInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`invitationStatus`, p.InvitationStatus)
	if p.Company != nil {
		queryParameters.Add(`company`, *p.Company)
	}
	if p.Position != nil {
		queryParameters.Add(`position`, *p.Position)
	}
	if p.Email != nil {
		queryParameters.Add(`email`, *p.Email)
	}
	if p.FirstName != nil {
		queryParameters.Add(`firstName`, *p.FirstName)
	}
	if p.LastName != nil {
		queryParameters.Add(`lastName`, *p.LastName)
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
	}
	if p.Telephone != nil {
		queryParameters.Add(`telephone`, *p.Telephone)
	}
	if p.UpdatedTime != nil {
		queryParameters.Add(`updatedTime`, strconv.FormatInt(*p.UpdatedTime, 10))
	}
	if p.ForceDuplicateInvitations != nil {
		queryParameters.Add(`forceDuplicateInvitations`, strconv.FormatBool(*p.ForceDuplicateInvitations))
	}
	if p.InviteCount != nil {
		queryParameters.Add(`inviteCount`, strconv.FormatInt(*p.InviteCount, 10))
	}
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.CheckInNotes != nil {
		queryParameters.Add(`checkInNotes`, *p.CheckInNotes)
	}
	if p.RelatedInvitationId != nil {
		queryParameters.Add(`relatedInvitationId`, *p.RelatedInvitationId)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/UpdateInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Invitation) UpdateInvitationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Invitation/UseCase/UpdateInvitation`,
		data,
		nil,
		nil,
	)
}
