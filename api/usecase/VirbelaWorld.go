/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type VirbelaWorld struct {
	restClient rest.RestClientInterface
}

func NewVirbelaWorld(restClient rest.RestClientInterface) *VirbelaWorld {
	return &VirbelaWorld{restClient}
}

// GET: Queries

func (t *VirbelaWorld) ListVirbelaWorlds() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Get(
		`/v2/VirbelaWorld/UseCase/ListVirbelaWorlds`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

func (t *VirbelaWorld) PopulateVirbelaWorldsFromVirbela() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Post(
		`/v2/VirbelaWorld/UseCase/PopulateVirbelaWorldsFromVirbela`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *VirbelaWorld) PopulateVirbelaWorldsFromVirbelaWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/VirbelaWorld/UseCase/PopulateVirbelaWorldsFromVirbela`,
		data,
		nil,
		nil,
	)
}
