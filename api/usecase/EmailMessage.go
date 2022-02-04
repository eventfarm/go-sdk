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

type EmailMessage struct {
	restClient rest.RestClientInterface
}

func NewEmailMessage(restClient rest.RestClientInterface) *EmailMessage {
	return &EmailMessage{restClient}
}

// GET: Queries

type GetEmailMessageParameters struct {
	EmailMessageId string
}

func (t *EmailMessage) GetEmailMessage(p *GetEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailMessageId`, p.EmailMessageId)

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/GetEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetEmailMessageStatsForEventParameters struct {
	EventId string
}

func (t *EmailMessage) GetEmailMessageStatsForEvent(p *GetEmailMessageStatsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/GetEmailMessageStatsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetScheduledBatchEmailMessageParameters struct {
	ScheduledBatchEmailMessageId string
}

func (t *EmailMessage) GetScheduledBatchEmailMessage(p *GetScheduledBatchEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`scheduledBatchEmailMessageId`, p.ScheduledBatchEmailMessageId)

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/GetScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetScheduledEmailMessageParameters struct {
	MongoQueueId string
}

func (t *EmailMessage) GetScheduledEmailMessage(p *GetScheduledEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`mongoQueueId`, p.MongoQueueId)

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/GetScheduledEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListOutboxEmailMessagesByEventParameters struct {
	EventId      string
	Page         *int64
	ItemsPerPage *int64
}

func (t *EmailMessage) ListOutboxEmailMessagesByEvent(p *ListOutboxEmailMessagesByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListOutboxEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListScheduledBatchEmailMessagesByEventParameters struct {
	EventId      string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-100
}

func (t *EmailMessage) ListScheduledBatchEmailMessagesByEvent(p *ListScheduledBatchEmailMessagesByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListScheduledBatchEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListScheduledEmailMessagesByEventParameters struct {
	EventId      string
	Page         *int64
	ItemsPerPage *int64
}

func (t *EmailMessage) ListScheduledEmailMessagesByEvent(p *ListScheduledEmailMessagesByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListScheduledEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListSentEmailMessagesByEventParameters struct {
	EventId      string
	Page         *int64
	ItemsPerPage *int64
}

func (t *EmailMessage) ListSentEmailMessagesByEvent(p *ListSentEmailMessagesByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListSentEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateEmailMessageParameters struct {
	EmailMessageType string
	EventId          string
	OwnerUserId      string
	EmailDesignId    string
	InvitationId     string
	Category         string
	SubCategory      string
}

func (t *EmailMessage) CreateEmailMessage(p *CreateEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailMessageType`, p.EmailMessageType)
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`category`, p.Category)
	queryParameters.Add(`subCategory`, p.SubCategory)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/CreateEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) CreateEmailMessageWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/CreateEmailMessage`,
		data,
		nil,
		nil,
	)
}

type CreatePreviewEmailMessageParameters struct {
	EventId         string
	OwnerUserId     string
	MessageContent  string
	ToEmails        []string
	FromName        string
	Subject         string
	LayoutType      string
	MessageType     string
	BackgroundColor string
	FromEmail       *string
	ReplyEmail      *string
	DomainMaskEmail *string
}

func (t *EmailMessage) CreatePreviewEmailMessage(p *CreatePreviewEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContent`, p.MessageContent)
	for i := range p.ToEmails {
		queryParameters.Add(`toEmails[]`, p.ToEmails[i])
	}
	queryParameters.Add(`fromName`, p.FromName)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`layoutType`, p.LayoutType)
	queryParameters.Add(`messageType`, p.MessageType)
	queryParameters.Add(`backgroundColor`, p.BackgroundColor)
	if p.FromEmail != nil {
		queryParameters.Add(`fromEmail`, *p.FromEmail)
	}
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/CreatePreviewEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) CreatePreviewEmailMessageWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/CreatePreviewEmailMessage`,
		data,
		nil,
		nil,
	)
}

type CreateScheduledBatchEmailMessageParameters struct {
	EventId                      string
	OwnerUserId                  string
	EmailDesignId                string
	Type                         string
	Targets                      []string // confirmed | purchased | assigned | unconfirmed | attended
	MessageSendTime              string
	Timezone                     string
	ScheduledBatchEmailMessageId *string
}

func (t *EmailMessage) CreateScheduledBatchEmailMessage(p *CreateScheduledBatchEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)
	queryParameters.Add(`type`, p.Type)
	for i := range p.Targets {
		queryParameters.Add(`targets[]`, p.Targets[i])
	}
	queryParameters.Add(`messageSendTime`, p.MessageSendTime)
	queryParameters.Add(`timezone`, p.Timezone)
	if p.ScheduledBatchEmailMessageId != nil {
		queryParameters.Add(`scheduledBatchEmailMessageId`, *p.ScheduledBatchEmailMessageId)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/CreateScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) CreateScheduledBatchEmailMessageWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/CreateScheduledBatchEmailMessage`,
		data,
		nil,
		nil,
	)
}

type RemoveScheduledBatchEmailMessageParameters struct {
	ScheduledBatchEmailMessageId string
}

func (t *EmailMessage) RemoveScheduledBatchEmailMessage(p *RemoveScheduledBatchEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`scheduledBatchEmailMessageId`, p.ScheduledBatchEmailMessageId)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/RemoveScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) RemoveScheduledBatchEmailMessageWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/RemoveScheduledBatchEmailMessage`,
		data,
		nil,
		nil,
	)
}

type ResendReceiptParameters struct {
	PaymentId string
}

func (t *EmailMessage) ResendReceipt(p *ResendReceiptParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`paymentId`, p.PaymentId)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/ResendReceipt`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) ResendReceiptWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/ResendReceipt`,
		data,
		nil,
		nil,
	)
}

type SendAMessagePreviewParameters struct {
	EventId                string
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessagePreview(p *SendAMessagePreviewParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessagePreview`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendAMessagePreviewWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendAMessagePreview`,
		data,
		nil,
		nil,
	)
}

type SendAMessageToAllCheckedInGuestsParameters struct {
	EventId                string
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	CcEmails               *[]string
	BccEmails              *[]string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessageToAllCheckedInGuests(p *SendAMessageToAllCheckedInGuestsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToAllCheckedInGuests`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendAMessageToAllCheckedInGuestsWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendAMessageToAllCheckedInGuests`,
		data,
		nil,
		nil,
	)
}

type SendAMessageToEventWaitListParameters struct {
	EventId                string
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	CcEmails               *[]string
	BccEmails              *[]string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessageToEventWaitList(p *SendAMessageToEventWaitListParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToEventWaitList`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendAMessageToEventWaitListWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendAMessageToEventWaitList`,
		data,
		nil,
		nil,
	)
}

type SendAMessageToGroupParameters struct {
	GroupId                    string
	EventId                    string
	OwnerUserId                string
	MessageContents            string
	Subject                    string
	FromName                   string
	ReplyEmail                 *string
	CcEmails                   *[]string
	BccEmails                  *[]string
	UseEventSpecificLayout     *bool
	BackgroundColor            *string
	DomainMaskEmail            *string
	EventGroupStatusFilterType *string
}

func (t *EmailMessage) SendAMessageToGroup(p *SendAMessageToGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}
	if p.EventGroupStatusFilterType != nil {
		queryParameters.Add(`eventGroupStatusFilterType`, *p.EventGroupStatusFilterType)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendAMessageToGroupWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendAMessageToGroup`,
		data,
		nil,
		nil,
	)
}

type SendAMessageToNoShowsParameters struct {
	EventId                string
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	CcEmails               *[]string
	BccEmails              *[]string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessageToNoShows(p *SendAMessageToNoShowsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToNoShows`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendAMessageToNoShowsWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendAMessageToNoShows`,
		data,
		nil,
		nil,
	)
}

type SendAMessageToTicketTypesParameters struct {
	EventId                string
	TicketTypeIds          []string
	InvitationStatusTypes  []string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	CcEmails               *[]string
	BccEmails              *[]string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessageToTicketTypes(p *SendAMessageToTicketTypesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	for i := range p.TicketTypeIds {
		queryParameters.Add(`ticketTypeIds[]`, p.TicketTypeIds[i])
	}
	for i := range p.InvitationStatusTypes {
		queryParameters.Add(`invitationStatusTypes[]`, p.InvitationStatusTypes[i])
	}
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToTicketTypes`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendAMessageToTicketTypesWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendAMessageToTicketTypes`,
		data,
		nil,
		nil,
	)
}

type SendAlertUserEmailParameters struct {
	ToEmail           string
	Subject           string
	TemplateName      string
	TemplateVariables []string
	CcEmails          *[]string
	BccEmails         *[]string
}

func (t *EmailMessage) SendAlertUserEmail(p *SendAlertUserEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`toEmail`, p.ToEmail)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`templateName`, p.TemplateName)
	for i := range p.TemplateVariables {
		queryParameters.Add(`templateVariables[]`, p.TemplateVariables[i])
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAlertUserEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendAlertUserEmailWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendAlertUserEmail`,
		data,
		nil,
		nil,
	)
}

type SendEmailMessageParameters struct {
	ToEmail             string
	Subject             string
	MessageContents     string
	EventEmail          string
	Layout              string
	Associations        *[]string
	EmailMessageDetails *[]string
	InvitationId        *string
	OwnerUserId         *string
	ContainsEchoLink    *bool
	EventId             *string
	RecipientUserId     *string
}

func (t *EmailMessage) SendEmailMessage(p *SendEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`toEmail`, p.ToEmail)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`eventEmail`, p.EventEmail)
	queryParameters.Add(`layout`, p.Layout)
	if p.Associations != nil {
		for i := range *p.Associations {
			queryParameters.Add(`associations[]`, (*p.Associations)[i])
		}
	}
	if p.EmailMessageDetails != nil {
		for i := range *p.EmailMessageDetails {
			queryParameters.Add(`emailMessageDetails[]`, (*p.EmailMessageDetails)[i])
		}
	}
	if p.InvitationId != nil {
		queryParameters.Add(`invitationId`, *p.InvitationId)
	}
	if p.OwnerUserId != nil {
		queryParameters.Add(`ownerUserId`, *p.OwnerUserId)
	}
	if p.ContainsEchoLink != nil {
		queryParameters.Add(`containsEchoLink`, strconv.FormatBool(*p.ContainsEchoLink))
	}
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}
	if p.RecipientUserId != nil {
		queryParameters.Add(`recipientUserId`, *p.RecipientUserId)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendEmailMessageWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendEmailMessage`,
		data,
		nil,
		nil,
	)
}

type SendGenericEmailParameters struct {
	ToEmail           string
	Subject           string
	TemplateName      string
	TemplateVariables []string
	FromEmail         string
	ReplyEmail        string
	CcEmails          *[]string
	BccEmails         *[]string
}

func (t *EmailMessage) SendGenericEmail(p *SendGenericEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`toEmail`, p.ToEmail)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`templateName`, p.TemplateName)
	for i := range p.TemplateVariables {
		queryParameters.Add(`templateVariables[]`, p.TemplateVariables[i])
	}
	queryParameters.Add(`fromEmail`, p.FromEmail)
	queryParameters.Add(`replyEmail`, p.ReplyEmail)
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendGenericEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendGenericEmailWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendGenericEmail`,
		data,
		nil,
		nil,
	)
}

type SendScheduledBatchEmailMessageParameters struct {
	ScheduledBatchEmailMessageId string
	UserId                       *string
}

func (t *EmailMessage) SendScheduledBatchEmailMessage(p *SendScheduledBatchEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`scheduledBatchEmailMessageId`, p.ScheduledBatchEmailMessageId)
	if p.UserId != nil {
		queryParameters.Add(`userId`, *p.UserId)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendScheduledBatchEmailMessageWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendScheduledBatchEmailMessage`,
		data,
		nil,
		nil,
	)
}

type SendSystemEmailParameters struct {
	ToEmail           string
	Subject           string
	TemplateName      string
	TemplateVariables []string
	CcEmails          *[]string
	BccEmails         *[]string
}

func (t *EmailMessage) SendSystemEmail(p *SendSystemEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`toEmail`, p.ToEmail)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`templateName`, p.TemplateName)
	for i := range p.TemplateVariables {
		queryParameters.Add(`templateVariables[]`, p.TemplateVariables[i])
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendSystemEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendSystemEmailWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendSystemEmail`,
		data,
		nil,
		nil,
	)
}

type SendVerifyEmailParameters struct {
	ToEmail           string
	Subject           string
	TemplateName      string
	TemplateVariables []string
	CcEmails          *[]string
	BccEmails         *[]string
	AppName           *string
	FromEmail         *string
	ReplyToEmail      *string
}

func (t *EmailMessage) SendVerifyEmail(p *SendVerifyEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`toEmail`, p.ToEmail)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`templateName`, p.TemplateName)
	for i := range p.TemplateVariables {
		queryParameters.Add(`templateVariables[]`, p.TemplateVariables[i])
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.AppName != nil {
		queryParameters.Add(`appName`, *p.AppName)
	}
	if p.FromEmail != nil {
		queryParameters.Add(`fromEmail`, *p.FromEmail)
	}
	if p.ReplyToEmail != nil {
		queryParameters.Add(`replyToEmail`, *p.ReplyToEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendVerifyEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SendVerifyEmailWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SendVerifyEmail`,
		data,
		nil,
		nil,
	)
}

type SimulateEmailMessageStatsForEventParameters struct {
	EventId string
}

func (t *EmailMessage) SimulateEmailMessageStatsForEvent(p *SimulateEmailMessageStatsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SimulateEmailMessageStatsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) SimulateEmailMessageStatsForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/SimulateEmailMessageStatsForEvent`,
		data,
		nil,
		nil,
	)
}

type UpdateScheduledBatchEmailMessageParameters struct {
	MongoQueueId    string
	EmailDesignId   string
	Targets         []string // confirmed | purchased | assigned | unconfirmed | attended
	MessageSendTime string
	Timezone        string
}

func (t *EmailMessage) UpdateScheduledBatchEmailMessage(p *UpdateScheduledBatchEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`mongoQueueId`, p.MongoQueueId)
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)
	for i := range p.Targets {
		queryParameters.Add(`targets[]`, p.Targets[i])
	}
	queryParameters.Add(`messageSendTime`, p.MessageSendTime)
	queryParameters.Add(`timezone`, p.Timezone)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/UpdateScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailMessage) UpdateScheduledBatchEmailMessageWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailMessage/UseCase/UpdateScheduledBatchEmailMessage`,
		data,
		nil,
		nil,
	)
}
