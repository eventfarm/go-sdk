/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type PoolPaymentGateway struct {
	restClient rest.RestClientInterface
}

func NewPoolPaymentGateway(restClient rest.RestClientInterface) *PoolPaymentGateway {
	return &PoolPaymentGateway{restClient}
}

// GET: Queries

type ListPaymentGatewaysForPoolParameters struct {
	PoolId string
}

func (t *PoolPaymentGateway) ListPaymentGatewaysForPool(p *ListPaymentGatewaysForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/PoolPaymentGateway/UseCase/ListPaymentGatewaysForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListPaymentGatewaysForSpreedlyParameters struct {
	PoolId string
}

func (t *PoolPaymentGateway) ListPaymentGatewaysForSpreedly(p *ListPaymentGatewaysForSpreedlyParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/PoolPaymentGateway/UseCase/ListPaymentGatewaysForSpreedly`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreatePaymentGatewayForPoolParameters struct {
	PoolId                 string
	Name                   string
	PaymentGatewayStringId string
	CreateParameters       []string
	PoolPaymentGatewayId   *string
}

func (t *PoolPaymentGateway) CreatePaymentGatewayForPool(p *CreatePaymentGatewayForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`paymentGatewayStringId`, p.PaymentGatewayStringId)
	for i := range p.CreateParameters {
		queryParameters.Add(`createParameters[]`, p.CreateParameters[i])
	}
	if p.PoolPaymentGatewayId != nil {
		queryParameters.Add(`poolPaymentGatewayId`, *p.PoolPaymentGatewayId)
	}

	return t.restClient.Post(
		`/v2/PoolPaymentGateway/UseCase/CreatePaymentGatewayForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolPaymentGateway) CreatePaymentGatewayForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolPaymentGateway/UseCase/CreatePaymentGatewayForPool`,
		data,
		nil,
		nil,
	)
}

type DeletePaymentGatewayForPoolParameters struct {
	PoolPaymentGatewayId string
}

func (t *PoolPaymentGateway) DeletePaymentGatewayForPool(p *DeletePaymentGatewayForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolPaymentGatewayId`, p.PoolPaymentGatewayId)

	return t.restClient.Post(
		`/v2/PoolPaymentGateway/UseCase/DeletePaymentGatewayForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolPaymentGateway) DeletePaymentGatewayForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolPaymentGateway/UseCase/DeletePaymentGatewayForPool`,
		data,
		nil,
		nil,
	)
}

type RemovePaymentGatewayForPoolParameters struct {
	PoolPaymentGatewayId string
}

func (t *PoolPaymentGateway) RemovePaymentGatewayForPool(p *RemovePaymentGatewayForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolPaymentGatewayId`, p.PoolPaymentGatewayId)

	return t.restClient.Post(
		`/v2/PoolPaymentGateway/UseCase/RemovePaymentGatewayForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolPaymentGateway) RemovePaymentGatewayForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolPaymentGateway/UseCase/RemovePaymentGatewayForPool`,
		data,
		nil,
		nil,
	)
}

type SetDefaultPaymentGatewayForPoolParameters struct {
	PoolPaymentGatewayId string
}

func (t *PoolPaymentGateway) SetDefaultPaymentGatewayForPool(p *SetDefaultPaymentGatewayForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolPaymentGatewayId`, p.PoolPaymentGatewayId)

	return t.restClient.Post(
		`/v2/PoolPaymentGateway/UseCase/SetDefaultPaymentGatewayForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolPaymentGateway) SetDefaultPaymentGatewayForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolPaymentGateway/UseCase/SetDefaultPaymentGatewayForPool`,
		data,
		nil,
		nil,
	)
}

type SetPoolPaymentGatewayForEventParameters struct {
	PoolPaymentGatewayId string
	EventId              string
}

func (t *PoolPaymentGateway) SetPoolPaymentGatewayForEvent(p *SetPoolPaymentGatewayForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolPaymentGatewayId`, p.PoolPaymentGatewayId)
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/PoolPaymentGateway/UseCase/SetPoolPaymentGatewayForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolPaymentGateway) SetPoolPaymentGatewayForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolPaymentGateway/UseCase/SetPoolPaymentGatewayForEvent`,
		data,
		nil,
		nil,
	)
}

type UpdatePaymentGatewayForPoolParameters struct {
	PoolPaymentGatewayId   string
	Name                   string
	PaymentGatewayStringId string
	UpdateParameters       []string
}

func (t *PoolPaymentGateway) UpdatePaymentGatewayForPool(p *UpdatePaymentGatewayForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolPaymentGatewayId`, p.PoolPaymentGatewayId)
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`paymentGatewayStringId`, p.PaymentGatewayStringId)
	for i := range p.UpdateParameters {
		queryParameters.Add(`updateParameters[]`, p.UpdateParameters[i])
	}

	return t.restClient.Post(
		`/v2/PoolPaymentGateway/UseCase/UpdatePaymentGatewayForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolPaymentGateway) UpdatePaymentGatewayForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolPaymentGateway/UseCase/UpdatePaymentGatewayForPool`,
		data,
		nil,
		nil,
	)
}
