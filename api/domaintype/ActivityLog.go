/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type ActivityLog struct {
}

func NewActivityLog() *ActivityLog {
	return &ActivityLog{}
}

type ActivityLogActionType struct {
	Slug                    string
	Name                    string
	Description             string
	IsAdded                 bool
	IsCreated               bool
	IsAddedToTransaction    bool
	IsAddedToWebConference  bool
	IsAddedToWorld          bool
	IsUpdatedForWorld       bool
	IsChangedRoleForWorld   bool
	IsChangedTeamIdForWorld bool
	IsChangedTitleForWorld  bool
	IsCheckIn               bool
	IsInteracted            bool
	IsEmail                 bool
	IsRemoved               bool
	IsResponded             bool
	IsUpdated               bool
	IsTransferred           bool
}

func (f *ActivityLog) ListActivityLogActionTypes() []ActivityLogActionType {
	return []ActivityLogActionType{
		{
			Slug:                    `added`,
			Name:                    `Added`,
			Description:             ``,
			IsAdded:                 true,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `created`,
			Name:                    `Created`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               true,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             true,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `added-to-transaction`,
			Name:                    `Added To Transaction`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    true,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `added-to-web-conference`,
			Name:                    `Added To Web Conference`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  true,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `added-to-world`,
			Name:                    `Added To World`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          true,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `updated-for-world`,
			Name:                    `Updated For World`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       true,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `changed-role-for-world`,
			Name:                    `Changed Role For World`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   true,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `changed-team-id-for-world`,
			Name:                    `Changed Team Id For World`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: true,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `changed-title-for-world`,
			Name:                    `Changed Title For World`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  true,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `check-in`,
			Name:                    `Check-In`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               true,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `interacted`,
			Name:                    `Interacted`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            true,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `email`,
			Name:                    `Email`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 true,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `removed`,
			Name:                    `Removed`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               true,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `responded`,
			Name:                    `Responded`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           false,
		},
		{
			Slug:                    `updated`,
			Name:                    `Updated`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               true,
			IsTransferred:           false,
		},
		{
			Slug:                    `transferred`,
			Name:                    `Transferred`,
			Description:             ``,
			IsAdded:                 false,
			IsCreated:               false,
			IsAddedToTransaction:    false,
			IsAddedToWebConference:  false,
			IsAddedToWorld:          false,
			IsUpdatedForWorld:       false,
			IsChangedRoleForWorld:   false,
			IsChangedTeamIdForWorld: false,
			IsChangedTitleForWorld:  false,
			IsCheckIn:               false,
			IsInteracted:            false,
			IsEmail:                 false,
			IsRemoved:               false,
			IsResponded:             false,
			IsUpdated:               false,
			IsTransferred:           true,
		},
	}
}
