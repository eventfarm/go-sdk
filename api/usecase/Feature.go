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

type Feature struct {
	restClient sdk.RestClientInterface
}

func NewFeature(restClient sdk.RestClientInterface) *Feature {
	return &Feature{restClient}
}

// GET: Queries

func (t *Feature) ListFeatures() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Get(
		`/v2/Feature/UseCase/ListFeatures`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
