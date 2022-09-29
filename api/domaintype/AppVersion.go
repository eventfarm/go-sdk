/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type AppVersion struct {
}

func NewAppVersion() *AppVersion {
	return &AppVersion{}
}

type AppVersionType struct {
	Slug                           string
	Name                           string
	Description                    string
	IsCheckInIos                   bool
	IsCheckInAndroid               bool
	IsTicketBlockManagementIos     bool
	IsTicketBlockManagementAndroid bool
	IsEFxIos                       bool
	IsEFxAndroid                   bool
	IsEventsIos                    bool
	IsEventsAndroid                bool
	IsLeadsFlutter                 bool
}

func (f *AppVersion) ListAppVersionTypes() []AppVersionType {
	return []AppVersionType{
		{
			Slug:                           `check-in-ios`,
			Name:                           `Check In iOS`,
			Description:                    ``,
			IsCheckInIos:                   true,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: false,
			IsEFxIos:                       false,
			IsEFxAndroid:                   false,
			IsEventsIos:                    false,
			IsEventsAndroid:                false,
			IsLeadsFlutter:                 false,
		},
		{
			Slug:                           `check-in-android`,
			Name:                           `Check In Android`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               true,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: false,
			IsEFxIos:                       false,
			IsEFxAndroid:                   false,
			IsEventsIos:                    false,
			IsEventsAndroid:                false,
			IsLeadsFlutter:                 false,
		},
		{
			Slug:                           `ticket-block-mgmt-ios`,
			Name:                           `Ticket Block Management iOS`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     true,
			IsTicketBlockManagementAndroid: false,
			IsEFxIos:                       false,
			IsEFxAndroid:                   false,
			IsEventsIos:                    false,
			IsEventsAndroid:                false,
			IsLeadsFlutter:                 false,
		},
		{
			Slug:                           `ticket-block-mgmt-android`,
			Name:                           `Ticket Block Management Android`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: true,
			IsEFxIos:                       false,
			IsEFxAndroid:                   false,
			IsEventsIos:                    false,
			IsEventsAndroid:                false,
			IsLeadsFlutter:                 false,
		},
		{
			Slug:                           `efx-ios`,
			Name:                           `EFx iOS`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: false,
			IsEFxIos:                       true,
			IsEFxAndroid:                   false,
			IsEventsIos:                    false,
			IsEventsAndroid:                false,
			IsLeadsFlutter:                 false,
		},
		{
			Slug:                           `efx-android`,
			Name:                           `EFx Android`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: false,
			IsEFxIos:                       false,
			IsEFxAndroid:                   true,
			IsEventsIos:                    false,
			IsEventsAndroid:                false,
			IsLeadsFlutter:                 false,
		},
		{
			Slug:                           `events-ios`,
			Name:                           `Events iOS`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: false,
			IsEFxIos:                       false,
			IsEFxAndroid:                   false,
			IsEventsIos:                    true,
			IsEventsAndroid:                false,
			IsLeadsFlutter:                 false,
		},
		{
			Slug:                           `events-android`,
			Name:                           `Events Android`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: false,
			IsEFxIos:                       false,
			IsEFxAndroid:                   false,
			IsEventsIos:                    false,
			IsEventsAndroid:                true,
			IsLeadsFlutter:                 false,
		},
		{
			Slug:                           `leads-flutter`,
			Name:                           `Leads Flutter`,
			Description:                    ``,
			IsCheckInIos:                   false,
			IsCheckInAndroid:               false,
			IsTicketBlockManagementIos:     false,
			IsTicketBlockManagementAndroid: false,
			IsEFxIos:                       false,
			IsEFxAndroid:                   false,
			IsEventsIos:                    false,
			IsEventsAndroid:                false,
			IsLeadsFlutter:                 true,
		},
	}
}
