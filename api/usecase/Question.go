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

type Question struct {
	restClient rest.RestClientInterface
}

func NewQuestion(restClient rest.RestClientInterface) *Question {
	return &Question{restClient}
}

// GET: Queries

type GetAllQuestionsForEventParameters struct {
	EventId  string
	WithData *[]string // Answers | TicketType | QuestionContexts
}

func (t *Question) GetAllQuestionsForEvent(p *GetAllQuestionsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Question/UseCase/GetAllQuestionsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetQuestionParameters struct {
	QuestionId string
	WithData   *[]string // Answers | TicketType | QuestionResponseCounts | AnswerQuestionResponseCounts | QuestionContexts | AnswerBindings
}

func (t *Question) GetQuestion(p *GetQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Question/UseCase/GetQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListQuestionsByEventAndContextParameters struct {
	EventId              string
	QuestionContextTypes *[]string // registration | lead
}

func (t *Question) ListQuestionsByEventAndContext(p *ListQuestionsByEventAndContextParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.QuestionContextTypes != nil {
		for i := range *p.QuestionContextTypes {
			queryParameters.Add(`questionContextTypes[]`, (*p.QuestionContextTypes)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Question/UseCase/ListQuestionsByEventAndContext`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListQuestionsForEventParameters struct {
	EventId           string
	WithData          *[]string // TicketType | Answers | AnswerBindings
	ShouldHideDeleted *bool
	Query             *string
	SortBy            *string
	SortDirection     *string
	Page              *int64
	ItemsPerPage      *int64
}

func (t *Question) ListQuestionsForEvent(p *ListQuestionsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*p.ShouldHideDeleted))
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
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
		`/v2/Question/UseCase/ListQuestionsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateAnswerBindingForQuestionParameters struct {
	QuestionId        string
	AnswerId          string
	AnswerBindingType string
}

func (t *Question) CreateAnswerBindingForQuestion(p *CreateAnswerBindingForQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`answerId`, p.AnswerId)
	queryParameters.Add(`answerBindingType`, p.AnswerBindingType)

	return t.restClient.Post(
		`/v2/Question/UseCase/CreateAnswerBindingForQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Question) CreateAnswerBindingForQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Question/UseCase/CreateAnswerBindingForQuestion`,
		data,
		nil,
		nil,
	)
}

type CreateQuestionParameters struct {
	EventId              string
	Text                 string
	QuestionType         string
	SortOrder            *int64
	IsRequired           *bool
	IsIndividual         *bool
	TicketTypeId         *string
	QuestionId           *string
	QuestionContextTypes *[]string // registration | lead
}

func (t *Question) CreateQuestion(p *CreateQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`text`, p.Text)
	queryParameters.Add(`questionType`, p.QuestionType)
	if p.SortOrder != nil {
		queryParameters.Add(`sortOrder`, strconv.FormatInt(*p.SortOrder, 10))
	}
	if p.IsRequired != nil {
		queryParameters.Add(`isRequired`, strconv.FormatBool(*p.IsRequired))
	}
	if p.IsIndividual != nil {
		queryParameters.Add(`isIndividual`, strconv.FormatBool(*p.IsIndividual))
	}
	if p.TicketTypeId != nil {
		queryParameters.Add(`ticketTypeId`, *p.TicketTypeId)
	}
	if p.QuestionId != nil {
		queryParameters.Add(`questionId`, *p.QuestionId)
	}
	if p.QuestionContextTypes != nil {
		for i := range *p.QuestionContextTypes {
			queryParameters.Add(`questionContextTypes[]`, (*p.QuestionContextTypes)[i])
		}
	}

	return t.restClient.Post(
		`/v2/Question/UseCase/CreateQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Question) CreateQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Question/UseCase/CreateQuestion`,
		data,
		nil,
		nil,
	)
}

type DeleteQuestionParameters struct {
	QuestionId string
}

func (t *Question) DeleteQuestion(p *DeleteQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)

	return t.restClient.Post(
		`/v2/Question/UseCase/DeleteQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Question) DeleteQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Question/UseCase/DeleteQuestion`,
		data,
		nil,
		nil,
	)
}

type DisableQuestionParameters struct {
	QuestionId string
}

func (t *Question) DisableQuestion(p *DisableQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)

	return t.restClient.Post(
		`/v2/Question/UseCase/DisableQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Question) DisableQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Question/UseCase/DisableQuestion`,
		data,
		nil,
		nil,
	)
}

type EnableQuestionParameters struct {
	QuestionId string
}

func (t *Question) EnableQuestion(p *EnableQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)

	return t.restClient.Post(
		`/v2/Question/UseCase/EnableQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Question) EnableQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Question/UseCase/EnableQuestion`,
		data,
		nil,
		nil,
	)
}

type RemoveAnswerBindingForQuestionParameters struct {
	AnswerBindingId string
}

func (t *Question) RemoveAnswerBindingForQuestion(p *RemoveAnswerBindingForQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerBindingId`, p.AnswerBindingId)

	return t.restClient.Post(
		`/v2/Question/UseCase/RemoveAnswerBindingForQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Question) RemoveAnswerBindingForQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Question/UseCase/RemoveAnswerBindingForQuestion`,
		data,
		nil,
		nil,
	)
}

type SetQuestionSortOrderParameters struct {
	QuestionId string
	SortOrder  int64
}

func (t *Question) SetQuestionSortOrder(p *SetQuestionSortOrderParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`sortOrder`, strconv.FormatInt(p.SortOrder, 10))

	return t.restClient.Post(
		`/v2/Question/UseCase/SetQuestionSortOrder`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Question) SetQuestionSortOrderWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Question/UseCase/SetQuestionSortOrder`,
		data,
		nil,
		nil,
	)
}

type UpdateAnswerBindingForQuestionParameters struct {
	AnswerId          string
	AnswerBindingType string
	AnswerBindingId   string
}

func (t *Question) UpdateAnswerBindingForQuestion(p *UpdateAnswerBindingForQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, p.AnswerId)
	queryParameters.Add(`answerBindingType`, p.AnswerBindingType)
	queryParameters.Add(`answerBindingId`, p.AnswerBindingId)

	return t.restClient.Post(
		`/v2/Question/UseCase/UpdateAnswerBindingForQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Question) UpdateAnswerBindingForQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Question/UseCase/UpdateAnswerBindingForQuestion`,
		data,
		nil,
		nil,
	)
}

type UpdateQuestionParameters struct {
	QuestionId           string
	Text                 string
	QuestionType         string
	QuestionContextTypes *[]string // registration | lead
	IsRequired           *bool
	IsIndividual         *bool
	TicketTypeId         *string
}

func (t *Question) UpdateQuestion(p *UpdateQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`text`, p.Text)
	queryParameters.Add(`questionType`, p.QuestionType)
	if p.QuestionContextTypes != nil {
		for i := range *p.QuestionContextTypes {
			queryParameters.Add(`questionContextTypes[]`, (*p.QuestionContextTypes)[i])
		}
	}
	if p.IsRequired != nil {
		queryParameters.Add(`isRequired`, strconv.FormatBool(*p.IsRequired))
	}
	if p.IsIndividual != nil {
		queryParameters.Add(`isIndividual`, strconv.FormatBool(*p.IsIndividual))
	}
	if p.TicketTypeId != nil {
		queryParameters.Add(`ticketTypeId`, *p.TicketTypeId)
	}

	return t.restClient.Post(
		`/v2/Question/UseCase/UpdateQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Question) UpdateQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Question/UseCase/UpdateQuestion`,
		data,
		nil,
		nil,
	)
}
