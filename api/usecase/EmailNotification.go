package usecase

import (
	"bitbucket.ef.network/go/sdk"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = fmt.Sprint("Foo")
var _ = http.NoBody

type EmailNotification struct {
	restClient sdk.RestClientInterface
}

func NewEmailNotification(restClient sdk.RestClientInterface) *EmailNotification {
	return &EmailNotification{restClient}
}

// GET: Queries
// @param string EventId

type GetOpenActionsForEventOverLastMonthParameters struct {
	EventId string
}

func (t *EmailNotification) GetOpenActionsForEventOverLastMonth(p *GetOpenActionsForEventOverLastMonthParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/EmailNotification/UseCase/GetOpenActionsForEventOverLastMonth`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string EmailMessageId
// @param string Type
// @param int CreatedAt
// @param string|null EventId
// @param string|null EmailNotificationId

type CreateSparkpostNotificationParameters struct {
	EmailMessageId      string
	Type                string
	CreatedAt           int
	EventId             *string
	EmailNotificationId *string
}

func (t *EmailNotification) CreateSparkpostNotification(p *CreateSparkpostNotificationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailMessageId`, p.EmailMessageId)
	queryParameters.Add(`type`, p.Type)
	queryParameters.Add(`createdAt`, strconv.Itoa(p.CreatedAt))
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}
	if p.EmailNotificationId != nil {
		queryParameters.Add(`emailNotificationId`, *p.EmailNotificationId)
	}

	return t.restClient.Post(
		`/v2/EmailNotification/UseCase/CreateSparkpostNotification`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param int|null TotalRecords

type SimulateEmailNotificationsForEventParameters struct {
	EventId      string
	TotalRecords *int
}

func (t *EmailNotification) SimulateEmailNotificationsForEvent(p *SimulateEmailNotificationsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.TotalRecords != nil {
		queryParameters.Add(`totalRecords`, strconv.Itoa(*p.TotalRecords))
	}

	return t.restClient.Post(
		`/v2/EmailNotification/UseCase/SimulateEmailNotificationsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}
