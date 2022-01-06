/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type HealthPass struct {
	restClient rest.RestClientInterface
}

func NewHealthPass(restClient rest.RestClientInterface) *HealthPass {
	return &HealthPass{restClient}
}

// GET: Queries

// POST: Commands

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
