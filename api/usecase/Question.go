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

type RemoveAnswerBindingForQuestionParameters struct {
	AnswerBindingId string
	QuestionId      string
}

func (t *Question) RemoveAnswerBindingForQuestion(p *RemoveAnswerBindingForQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerBindingId`, p.AnswerBindingId)
	queryParameters.Add(`questionId`, p.QuestionId)

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
