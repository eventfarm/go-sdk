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

type UserIdentifier struct {
	restClient sdk.RestClientInterface
}

func NewUserIdentifier(restClient sdk.RestClientInterface) *UserIdentifier {
	return &UserIdentifier{restClient}
}

// GET: Queries

// POST: Commands
// @param string UserIdentifierId
// @param string Identifier
// @param string|null PoolId

type SetUserIdentifierParameters struct {
	UserIdentifierId string
	Identifier       string
	PoolId           *string
}

func (t *UserIdentifier) SetUserIdentifier(p *SetUserIdentifierParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userIdentifierId`, p.UserIdentifierId)
	queryParameters.Add(`identifier`, p.Identifier)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}

	return t.restClient.Post(
		`/v2/UserIdentifier/UseCase/SetUserIdentifier`,
		&queryParameters,
		nil,
		nil,
	)
}
