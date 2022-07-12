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

type HealthPass struct {
	restClient rest.RestClientInterface
}

func NewHealthPass(restClient rest.RestClientInterface) *HealthPass {
	return &HealthPass{restClient}
}

// GET: Queries

type GetHealthPassScoreCountsForEventParameters struct {
	EventId string
}

func (t *HealthPass) GetHealthPassScoreCountsForEvent(p *GetHealthPassScoreCountsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/HealthPass/UseCase/GetHealthPassScoreCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddUserHealthPassRecordFromClearParameters struct {
	FirstName           string
	LastName            string
	Email               string
	Score               string
	FinalScoringTime    string
	ExpiresAt           string
	HealthPassShortCode string
	ExternalUserId      string
}

func (t *HealthPass) AddUserHealthPassRecordFromClear(p *AddUserHealthPassRecordFromClearParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`firstName`, p.FirstName)
	queryParameters.Add(`lastName`, p.LastName)
	queryParameters.Add(`email`, p.Email)
	queryParameters.Add(`score`, p.Score)
	queryParameters.Add(`finalScoringTime`, p.FinalScoringTime)
	queryParameters.Add(`expiresAt`, p.ExpiresAt)
	queryParameters.Add(`healthPassShortCode`, p.HealthPassShortCode)
	queryParameters.Add(`externalUserId`, p.ExternalUserId)

	return t.restClient.Post(
		`/v2/HealthPass/UseCase/AddUserHealthPassRecordFromClear`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *HealthPass) AddUserHealthPassRecordFromClearWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/HealthPass/UseCase/AddUserHealthPassRecordFromClear`,
		data,
		nil,
		nil,
	)
}

type PostProcessAndImportHealthPassParameters struct {
	HealthPassImportId string
}

func (t *HealthPass) PostProcessAndImportHealthPass(p *PostProcessAndImportHealthPassParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`healthPassImportId`, p.HealthPassImportId)

	return t.restClient.Post(
		`/v2/HealthPass/UseCase/PostProcessAndImportHealthPass`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *HealthPass) PostProcessAndImportHealthPassWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/HealthPass/UseCase/PostProcessAndImportHealthPass`,
		data,
		nil,
		nil,
	)
}

type PostProcessHealthPassImportParameters struct {
	HealthPassImportId string
}

func (t *HealthPass) PostProcessHealthPassImport(p *PostProcessHealthPassImportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`healthPassImportId`, p.HealthPassImportId)

	return t.restClient.Post(
		`/v2/HealthPass/UseCase/PostProcessHealthPassImport`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *HealthPass) PostProcessHealthPassImportWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/HealthPass/UseCase/PostProcessHealthPassImport`,
		data,
		nil,
		nil,
	)
}

type PreProcessSpreadsheetForHealthPassImportParameters struct {
	Spreadsheet string
}

func (t *HealthPass) PreProcessSpreadsheetForHealthPassImport(p *PreProcessSpreadsheetForHealthPassImportParameters) (r *http.Response, err error) {
	// TODO
	return
}

func (t *HealthPass) PreProcessSpreadsheetForHealthPassImportWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/HealthPass/UseCase/PreProcessSpreadsheetForHealthPassImport`,
		data,
		nil,
		nil,
	)
}

type ProcessRecordsParameters struct {
	FirstResult int64
	MaxResults  int64
}

func (t *HealthPass) ProcessRecords(p *ProcessRecordsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`firstResult`, strconv.FormatInt(p.FirstResult, 10))
	queryParameters.Add(`maxResults`, strconv.FormatInt(p.MaxResults, 10))

	return t.restClient.Post(
		`/v2/HealthPass/UseCase/ProcessRecords`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *HealthPass) ProcessRecordsWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/HealthPass/UseCase/ProcessRecords`,
		data,
		nil,
		nil,
	)
}
