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

type Pool struct {
	restClient rest.RestClientInterface
}

func NewPool(restClient rest.RestClientInterface) *Pool {
	return &Pool{restClient}
}

// GET: Queries

type GetPoolParameters struct {
	PoolId string
}

func (t *Pool) GetPool(p *GetPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Pool/UseCase/GetPool`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetPoolContractParameters struct {
	PoolId string
}

func (t *Pool) GetPoolContract(p *GetPoolContractParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Pool/UseCase/GetPoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListAccessiblePoolsForUserParameters struct {
	UserId        string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
	SortBy        *string
	SortDirection *string
}

func (t *Pool) ListAccessiblePoolsForUser(p *ListAccessiblePoolsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListAccessiblePoolsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListPoolAllotmentsForPoolParameters struct {
	PoolId        string
	SortBy        *string
	SortDirection *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
}

func (t *Pool) ListPoolAllotmentsForPool(p *ListPoolAllotmentsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPoolAllotmentsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListPoolContactsByPoolIdParameters struct {
	PoolId       string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-500
	WithData     *[]string
}

func (t *Pool) ListPoolContactsByPoolId(p *ListPoolContactsByPoolIdParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPoolContactsByPoolId`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListPoolsParameters struct {
	Name              *string
	SortBy            *string
	SortDirection     *string
	Page              *int64 // >= 1
	ItemsPerPage      *int64 // 1-100
	ShouldHideDeleted *bool
}

func (t *Pool) ListPools(p *ListPoolsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*p.ShouldHideDeleted))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPools`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListTagsForPoolParameters struct {
	PoolId        string
	SortBy        *string
	SortDirection *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
}

func (t *Pool) ListTagsForPool(p *ListTagsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListTagsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListUniqueTagNamesForPoolParameters struct {
	PoolId string
}

func (t *Pool) ListUniqueTagNamesForPool(p *ListUniqueTagNamesForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListUniqueTagNamesForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type ArchivePoolParameters struct {
	PoolId string
}

func (t *Pool) ArchivePool(p *ArchivePoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Post(
		`/v2/Pool/UseCase/ArchivePool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) ArchivePoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/ArchivePool`,
		data,
		nil,
		nil,
	)
}

type CreatePoolParameters struct {
	Name      string
	ShortName *string
	PoolId    *string
}

func (t *Pool) CreatePool(p *CreatePoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`name`, p.Name)
	if p.ShortName != nil {
		queryParameters.Add(`shortName`, *p.ShortName)
	}
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/CreatePool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) CreatePoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/CreatePool`,
		data,
		nil,
		nil,
	)
}

type CreatePoolContractParameters struct {
	PoolId           string
	PoolContractType string
	StartDate        int64
	EndDate          int64
	PoolContractId   *string
}

func (t *Pool) CreatePoolContract(p *CreatePoolContractParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`poolContractType`, p.PoolContractType)
	queryParameters.Add(`startDate`, strconv.FormatInt(p.StartDate, 10))
	queryParameters.Add(`endDate`, strconv.FormatInt(p.EndDate, 10))
	if p.PoolContractId != nil {
		queryParameters.Add(`poolContractId`, *p.PoolContractId)
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/CreatePoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) CreatePoolContractWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/CreatePoolContract`,
		data,
		nil,
		nil,
	)
}

type CreatePoolWebhookParameters struct {
	PoolId          string
	PoolWebhookType string
	Url             string
}

func (t *Pool) CreatePoolWebhook(p *CreatePoolWebhookParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`poolWebhookType`, p.PoolWebhookType)
	queryParameters.Add(`url`, p.Url)

	return t.restClient.Post(
		`/v2/Pool/UseCase/CreatePoolWebhook`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) CreatePoolWebhookWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/CreatePoolWebhook`,
		data,
		nil,
		nil,
	)
}

type DeletePoolParameters struct {
	PoolId string
}

func (t *Pool) DeletePool(p *DeletePoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Post(
		`/v2/Pool/UseCase/DeletePool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) DeletePoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/DeletePool`,
		data,
		nil,
		nil,
	)
}

type RemoveAllUserRolesFromPoolParameters struct {
	UserId string
	PoolId string
}

func (t *Pool) RemoveAllUserRolesFromPool(p *RemoveAllUserRolesFromPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Post(
		`/v2/Pool/UseCase/RemoveAllUserRolesFromPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) RemoveAllUserRolesFromPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/RemoveAllUserRolesFromPool`,
		data,
		nil,
		nil,
	)
}

type RemoveCustomerDisplayNameForPoolParameters struct {
	PoolId string
}

func (t *Pool) RemoveCustomerDisplayNameForPool(p *RemoveCustomerDisplayNameForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Post(
		`/v2/Pool/UseCase/RemoveCustomerDisplayNameForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) RemoveCustomerDisplayNameForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/RemoveCustomerDisplayNameForPool`,
		data,
		nil,
		nil,
	)
}

type RemovePrivacyPolicyLinkForPoolParameters struct {
	PoolId string
}

func (t *Pool) RemovePrivacyPolicyLinkForPool(p *RemovePrivacyPolicyLinkForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Post(
		`/v2/Pool/UseCase/RemovePrivacyPolicyLinkForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) RemovePrivacyPolicyLinkForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/RemovePrivacyPolicyLinkForPool`,
		data,
		nil,
		nil,
	)
}

type SendUpgradeRequestToCsmParameters struct {
	PoolId                string
	UserId                string
	SlackUserId           *string
	RequestedFeatureSlugs *[]string
	Other                 *string
}

func (t *Pool) SendUpgradeRequestToCsm(p *SendUpgradeRequestToCsmParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	if p.SlackUserId != nil {
		queryParameters.Add(`slackUserId`, *p.SlackUserId)
	}
	if p.RequestedFeatureSlugs != nil {
		for i := range *p.RequestedFeatureSlugs {
			queryParameters.Add(`requestedFeatureSlugs[]`, (*p.RequestedFeatureSlugs)[i])
		}
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/SendUpgradeRequestToCsm`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) SendUpgradeRequestToCsmWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/SendUpgradeRequestToCsm`,
		data,
		nil,
		nil,
	)
}

type SetCustomerDisplayNameForPoolParameters struct {
	PoolId              string
	CustomerDisplayName string
}

func (t *Pool) SetCustomerDisplayNameForPool(p *SetCustomerDisplayNameForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`customerDisplayName`, p.CustomerDisplayName)

	return t.restClient.Post(
		`/v2/Pool/UseCase/SetCustomerDisplayNameForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) SetCustomerDisplayNameForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/SetCustomerDisplayNameForPool`,
		data,
		nil,
		nil,
	)
}

type SetNameForPoolParameters struct {
	PoolId string
	Name   string
}

func (t *Pool) SetNameForPool(p *SetNameForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Pool/UseCase/SetNameForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) SetNameForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/SetNameForPool`,
		data,
		nil,
		nil,
	)
}

type SetPrivacyPolicyLinkForPoolParameters struct {
	PoolId            string
	PrivacyPolicyLink string
}

func (t *Pool) SetPrivacyPolicyLinkForPool(p *SetPrivacyPolicyLinkForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`privacyPolicyLink`, p.PrivacyPolicyLink)

	return t.restClient.Post(
		`/v2/Pool/UseCase/SetPrivacyPolicyLinkForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) SetPrivacyPolicyLinkForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/SetPrivacyPolicyLinkForPool`,
		data,
		nil,
		nil,
	)
}

type SetShortNameForPoolParameters struct {
	PoolId    string
	ShortName string
}

func (t *Pool) SetShortNameForPool(p *SetShortNameForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`shortName`, p.ShortName)

	return t.restClient.Post(
		`/v2/Pool/UseCase/SetShortNameForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) SetShortNameForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/SetShortNameForPool`,
		data,
		nil,
		nil,
	)
}

type UpdatePoolContractParameters struct {
	PoolContractId   string
	PoolContractType *string
	StartDate        *int64
	EndDate          *int64
}

func (t *Pool) UpdatePoolContract(p *UpdatePoolContractParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolContractId`, p.PoolContractId)
	if p.PoolContractType != nil {
		queryParameters.Add(`poolContractType`, *p.PoolContractType)
	}
	if p.StartDate != nil {
		queryParameters.Add(`startDate`, strconv.FormatInt(*p.StartDate, 10))
	}
	if p.EndDate != nil {
		queryParameters.Add(`endDate`, strconv.FormatInt(*p.EndDate, 10))
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/UpdatePoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Pool) UpdatePoolContractWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Pool/UseCase/UpdatePoolContract`,
		data,
		nil,
		nil,
	)
}
