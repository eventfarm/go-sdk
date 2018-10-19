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

type Canvas struct {
	restClient sdk.RestClientInterface
}

func NewCanvas(restClient sdk.RestClientInterface) *Canvas {
	return &Canvas{restClient}
}

// GET: Queries

// POST: Commands
// @param string EventId

type EnableCanvasForEventParameters struct {
	EventId string
}

func (t *Canvas) EnableCanvasForEvent(p *EnableCanvasForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Canvas/UseCase/EnableCanvasForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}
