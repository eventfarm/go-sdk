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

type Answer struct {
	restClient rest.RestClientInterface
}

func NewAnswer(restClient rest.RestClientInterface) *Answer {
	return &Answer{restClient}
}

// GET: Queries

// POST: Commands

type CreateAnswerParameters struct {
	QuestionId string
	Text       string
	SortOrder  *int64
	IsDefault  *bool
	AnswerId   *string
}

func (t *Answer) CreateAnswer(p *CreateAnswerParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`text`, p.Text)
	if p.SortOrder != nil {
		queryParameters.Add(`sortOrder`, strconv.FormatInt(*p.SortOrder, 10))
	}
	if p.IsDefault != nil {
		queryParameters.Add(`isDefault`, strconv.FormatBool(*p.IsDefault))
	}
	if p.AnswerId != nil {
		queryParameters.Add(`answerId`, *p.AnswerId)
	}

	return t.restClient.Post(
		`/v2/Answer/UseCase/CreateAnswer`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Answer) CreateAnswerWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Answer/UseCase/CreateAnswer`,
		data,
		nil,
		nil,
	)
}

type DeleteAnswerParameters struct {
	AnswerId string
}

func (t *Answer) DeleteAnswer(p *DeleteAnswerParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, p.AnswerId)

	return t.restClient.Post(
		`/v2/Answer/UseCase/DeleteAnswer`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Answer) DeleteAnswerWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Answer/UseCase/DeleteAnswer`,
		data,
		nil,
		nil,
	)
}

type SetAnswerSortOrderParameters struct {
	AnswerId  string
	SortOrder int64
}

func (t *Answer) SetAnswerSortOrder(p *SetAnswerSortOrderParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, p.AnswerId)
	queryParameters.Add(`sortOrder`, strconv.FormatInt(p.SortOrder, 10))

	return t.restClient.Post(
		`/v2/Answer/UseCase/SetAnswerSortOrder`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Answer) SetAnswerSortOrderWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Answer/UseCase/SetAnswerSortOrder`,
		data,
		nil,
		nil,
	)
}

type UpdateAnswerParameters struct {
	AnswerId  string
	Text      string
	IsDefault *bool
}

func (t *Answer) UpdateAnswer(p *UpdateAnswerParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, p.AnswerId)
	queryParameters.Add(`text`, p.Text)
	if p.IsDefault != nil {
		queryParameters.Add(`isDefault`, strconv.FormatBool(*p.IsDefault))
	}

	return t.restClient.Post(
		`/v2/Answer/UseCase/UpdateAnswer`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Answer) UpdateAnswerWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Answer/UseCase/UpdateAnswer`,
		data,
		nil,
		nil,
	)
}
