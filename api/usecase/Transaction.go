/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type Transaction struct {
	restClient rest.RestClientInterface
}

func NewTransaction(restClient rest.RestClientInterface) *Transaction {
	return &Transaction{restClient}
}

// GET: Queries

type GetTransactionParameters struct {
	TransactionId string
	WithData      *[]string // Invitations | ActivityLogs | Payments
}

func (t *Transaction) GetTransaction(p *GetTransactionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`transactionId`, p.TransactionId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Transaction/UseCase/GetTransaction`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddInvitationToTransactionParameters struct {
	TransactionId string
	InvitationId  string
	Invitations   *[]string
}

func (t *Transaction) AddInvitationToTransaction(p *AddInvitationToTransactionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`transactionId`, p.TransactionId)
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.Invitations != nil {
		for i := range *p.Invitations {
			queryParameters.Add(`invitations[]`, (*p.Invitations)[i])
		}
	}

	return t.restClient.Post(
		`/v2/Transaction/UseCase/AddInvitationToTransaction`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Transaction) AddInvitationToTransactionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Transaction/UseCase/AddInvitationToTransaction`,
		data,
		nil,
		nil,
	)
}

type CreateTransactionParameters struct {
	PoolId                string
	TransactionId         *string
	Invitations           *[]string
	Payment               *[]string
	RequestedPromotionIds *[]string
	PrimaryInvitationId   *string
}

func (t *Transaction) CreateTransaction(p *CreateTransactionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.TransactionId != nil {
		queryParameters.Add(`transactionId`, *p.TransactionId)
	}
	if p.Invitations != nil {
		for i := range *p.Invitations {
			queryParameters.Add(`invitations[]`, (*p.Invitations)[i])
		}
	}
	if p.Payment != nil {
		for i := range *p.Payment {
			queryParameters.Add(`payment[]`, (*p.Payment)[i])
		}
	}
	if p.RequestedPromotionIds != nil {
		for i := range *p.RequestedPromotionIds {
			queryParameters.Add(`requestedPromotionIds[]`, (*p.RequestedPromotionIds)[i])
		}
	}
	if p.PrimaryInvitationId != nil {
		queryParameters.Add(`primaryInvitationId`, *p.PrimaryInvitationId)
	}

	return t.restClient.Post(
		`/v2/Transaction/UseCase/CreateTransaction`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Transaction) CreateTransactionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Transaction/UseCase/CreateTransaction`,
		data,
		nil,
		nil,
	)
}
