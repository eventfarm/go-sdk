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

type PaymentGateway struct {
	restClient rest.RestClientInterface
}

func NewPaymentGateway(restClient rest.RestClientInterface) *PaymentGateway {
	return &PaymentGateway{restClient}
}

// GET: Queries

func (t *PaymentGateway) ListPaymentGateways() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Get(
		`/v2/PaymentGateway/UseCase/ListPaymentGateways`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreatePaymentGatewayParameters struct {
	StringId    string
	Name        string
	CompanyName string
	CompanyUrl  *string
	LogoUrl     *string
}

func (t *PaymentGateway) CreatePaymentGateway(p *CreatePaymentGatewayParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stringId`, p.StringId)
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`companyName`, p.CompanyName)
	if p.CompanyUrl != nil {
		queryParameters.Add(`companyUrl`, *p.CompanyUrl)
	}
	if p.LogoUrl != nil {
		queryParameters.Add(`logoUrl`, *p.LogoUrl)
	}

	return t.restClient.Post(
		`/v2/PaymentGateway/UseCase/CreatePaymentGateway`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PaymentGateway) CreatePaymentGatewayWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PaymentGateway/UseCase/CreatePaymentGateway`,
		data,
		nil,
		nil,
	)
}

type SetSortOrderForPaymentGatewayParameters struct {
	PaymentGatewayId string
	SortOrder        string
}

func (t *PaymentGateway) SetSortOrderForPaymentGateway(p *SetSortOrderForPaymentGatewayParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`paymentGatewayId`, p.PaymentGatewayId)
	queryParameters.Add(`sortOrder`, p.SortOrder)

	return t.restClient.Post(
		`/v2/PaymentGateway/UseCase/SetSortOrderForPaymentGateway`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PaymentGateway) SetSortOrderForPaymentGatewayWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PaymentGateway/UseCase/SetSortOrderForPaymentGateway`,
		data,
		nil,
		nil,
	)
}

type UpdatePaymentGatewayParameters struct {
	PaymentGatewayId int64
	Name             *string
	CompanyName      *string
	CompanyUrl       *string
	LogoUrl          *string
	SortOrder        *int64
}

func (t *PaymentGateway) UpdatePaymentGateway(p *UpdatePaymentGatewayParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`paymentGatewayId`, strconv.FormatInt(p.PaymentGatewayId, 10))
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.CompanyName != nil {
		queryParameters.Add(`companyName`, *p.CompanyName)
	}
	if p.CompanyUrl != nil {
		queryParameters.Add(`companyUrl`, *p.CompanyUrl)
	}
	if p.LogoUrl != nil {
		queryParameters.Add(`logoUrl`, *p.LogoUrl)
	}
	if p.SortOrder != nil {
		queryParameters.Add(`sortOrder`, strconv.FormatInt(*p.SortOrder, 10))
	}

	return t.restClient.Post(
		`/v2/PaymentGateway/UseCase/UpdatePaymentGateway`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PaymentGateway) UpdatePaymentGatewayWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PaymentGateway/UseCase/UpdatePaymentGateway`,
		data,
		nil,
		nil,
	)
}
