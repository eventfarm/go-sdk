/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type PaymentGateway struct {
	restClient rest.RestClientInterface
}

func NewPaymentGateway(restClient rest.RestClientInterface) *PaymentGateway {
	return &PaymentGateway{restClient}
}

// GET: Queries

type ListPaymentGatewaysForPoolParameters struct {
	PoolId string
}

func (t *PaymentGateway) ListPaymentGatewaysForPool(p *ListPaymentGatewaysForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/PaymentGateway/UseCase/ListPaymentGatewaysForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreatePaymentGatewayForPoolParameters struct {
	PoolId             string
	PaymentGatewayType string
	GatewayToken       string
	PaymentGatewayId   *string
}

func (t *PaymentGateway) CreatePaymentGatewayForPool(p *CreatePaymentGatewayForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`paymentGatewayType`, p.PaymentGatewayType)
	queryParameters.Add(`gatewayToken`, p.GatewayToken)
	if p.PaymentGatewayId != nil {
		queryParameters.Add(`paymentGatewayId`, *p.PaymentGatewayId)
	}

	return t.restClient.Post(
		`/v2/PaymentGateway/UseCase/CreatePaymentGatewayForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PaymentGateway) CreatePaymentGatewayForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PaymentGateway/UseCase/CreatePaymentGatewayForPool`,
		data,
		nil,
		nil,
	)
}

type DeletePaymentGatewayForPoolParameters struct {
	PaymentGatewayId string
}

func (t *PaymentGateway) DeletePaymentGatewayForPool(p *DeletePaymentGatewayForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`paymentGatewayId`, p.PaymentGatewayId)

	return t.restClient.Post(
		`/v2/PaymentGateway/UseCase/DeletePaymentGatewayForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PaymentGateway) DeletePaymentGatewayForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PaymentGateway/UseCase/DeletePaymentGatewayForPool`,
		data,
		nil,
		nil,
	)
}

type RemovePaymentGatewayForPoolParameters struct {
	PaymentGatewayId string
}

func (t *PaymentGateway) RemovePaymentGatewayForPool(p *RemovePaymentGatewayForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`paymentGatewayId`, p.PaymentGatewayId)

	return t.restClient.Post(
		`/v2/PaymentGateway/UseCase/RemovePaymentGatewayForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PaymentGateway) RemovePaymentGatewayForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PaymentGateway/UseCase/RemovePaymentGatewayForPool`,
		data,
		nil,
		nil,
	)
}

type UpdateGatewayTokenForTypeParameters struct {
	GatewayToken     string
	PaymentGatewayId string
}

func (t *PaymentGateway) UpdateGatewayTokenForType(p *UpdateGatewayTokenForTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`gatewayToken`, p.GatewayToken)
	queryParameters.Add(`paymentGatewayId`, p.PaymentGatewayId)

	return t.restClient.Post(
		`/v2/PaymentGateway/UseCase/UpdateGatewayTokenForType`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PaymentGateway) UpdateGatewayTokenForTypeWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PaymentGateway/UseCase/UpdateGatewayTokenForType`,
		data,
		nil,
		nil,
	)
}
