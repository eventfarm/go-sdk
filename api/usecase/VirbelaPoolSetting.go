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

type VirbelaPoolSetting struct {
	restClient rest.RestClientInterface
}

func NewVirbelaPoolSetting(restClient rest.RestClientInterface) *VirbelaPoolSetting {
	return &VirbelaPoolSetting{restClient}
}

// GET: Queries

type GetVirbelaPoolSettingByPoolParameters struct {
	PoolId string
}

func (t *VirbelaPoolSetting) GetVirbelaPoolSettingByPool(p *GetVirbelaPoolSettingByPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/VirbelaPoolSetting/UseCase/GetVirbelaPoolSettingByPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateVirbelaPoolSettingParameters struct {
	PoolId             string
	WorldId            string
	WorldName          string
	DisplayName        string
	TitleMapping       string
	TeamId             int64 // 0-51
	MacDownloadURL     string
	WindowsDownloadURL string
}

func (t *VirbelaPoolSetting) CreateVirbelaPoolSetting(p *CreateVirbelaPoolSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`worldId`, p.WorldId)
	queryParameters.Add(`worldName`, p.WorldName)
	queryParameters.Add(`displayName`, p.DisplayName)
	queryParameters.Add(`titleMapping`, p.TitleMapping)
	queryParameters.Add(`teamId`, strconv.FormatInt(p.TeamId, 10))
	queryParameters.Add(`macDownloadURL`, p.MacDownloadURL)
	queryParameters.Add(`windowsDownloadURL`, p.WindowsDownloadURL)

	return t.restClient.Post(
		`/v2/VirbelaPoolSetting/UseCase/CreateVirbelaPoolSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *VirbelaPoolSetting) CreateVirbelaPoolSettingWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/VirbelaPoolSetting/UseCase/CreateVirbelaPoolSetting`,
		data,
		nil,
		nil,
	)
}

type RemoveVirbelaPoolSettingParameters struct {
	VirbelaPoolSettingId string
}

func (t *VirbelaPoolSetting) RemoveVirbelaPoolSetting(p *RemoveVirbelaPoolSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`virbelaPoolSettingId`, p.VirbelaPoolSettingId)

	return t.restClient.Post(
		`/v2/VirbelaPoolSetting/UseCase/RemoveVirbelaPoolSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *VirbelaPoolSetting) RemoveVirbelaPoolSettingWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/VirbelaPoolSetting/UseCase/RemoveVirbelaPoolSetting`,
		data,
		nil,
		nil,
	)
}

type UpdateVirbelaPoolSettingParameters struct {
	VirbelaPoolSettingId string
	WorldId              string
	WorldName            string
	DisplayName          string
	TitleMapping         string
	TeamId               int64 // 0-51
	MacDownloadURL       string
	WindowsDownloadURL   string
}

func (t *VirbelaPoolSetting) UpdateVirbelaPoolSetting(p *UpdateVirbelaPoolSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`virbelaPoolSettingId`, p.VirbelaPoolSettingId)
	queryParameters.Add(`worldId`, p.WorldId)
	queryParameters.Add(`worldName`, p.WorldName)
	queryParameters.Add(`displayName`, p.DisplayName)
	queryParameters.Add(`titleMapping`, p.TitleMapping)
	queryParameters.Add(`teamId`, strconv.FormatInt(p.TeamId, 10))
	queryParameters.Add(`macDownloadURL`, p.MacDownloadURL)
	queryParameters.Add(`windowsDownloadURL`, p.WindowsDownloadURL)

	return t.restClient.Post(
		`/v2/VirbelaPoolSetting/UseCase/UpdateVirbelaPoolSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *VirbelaPoolSetting) UpdateVirbelaPoolSettingWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/VirbelaPoolSetting/UseCase/UpdateVirbelaPoolSetting`,
		data,
		nil,
		nil,
	)
}
