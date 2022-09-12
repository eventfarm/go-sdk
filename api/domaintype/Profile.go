/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Profile struct {
}

func NewProfile() *Profile {
	return &Profile{}
}

type EventProfileType struct {
	Slug          string
	Name          string
	Description   string
	IsSpeaker     bool
	IsPresenter   bool
	IsSponsor     bool
	IsPanelist    bool
	IsModerator   bool
	IsHost        bool
	IsParticipant bool
	IsOther       bool
}

type ProfileType struct {
	Slug        string
	Name        string
	Description string
	IsPerson    bool
	IsCompany   bool
	IsOther     bool
}

func (f *Profile) ListEventProfileTypes() []EventProfileType {
	return []EventProfileType{
		{
			Slug:          `speaker`,
			Name:          `Speaker`,
			Description:   ``,
			IsSpeaker:     true,
			IsPresenter:   false,
			IsSponsor:     false,
			IsPanelist:    false,
			IsModerator:   false,
			IsHost:        false,
			IsParticipant: false,
			IsOther:       false,
		},
		{
			Slug:          `presenter`,
			Name:          `Presenter`,
			Description:   ``,
			IsSpeaker:     false,
			IsPresenter:   true,
			IsSponsor:     false,
			IsPanelist:    false,
			IsModerator:   false,
			IsHost:        false,
			IsParticipant: false,
			IsOther:       false,
		},
		{
			Slug:          `sponsor`,
			Name:          `Sponsor`,
			Description:   ``,
			IsSpeaker:     false,
			IsPresenter:   false,
			IsSponsor:     true,
			IsPanelist:    false,
			IsModerator:   false,
			IsHost:        false,
			IsParticipant: false,
			IsOther:       false,
		},
		{
			Slug:          `panelist`,
			Name:          `Panelist`,
			Description:   ``,
			IsSpeaker:     false,
			IsPresenter:   false,
			IsSponsor:     false,
			IsPanelist:    true,
			IsModerator:   false,
			IsHost:        false,
			IsParticipant: false,
			IsOther:       false,
		},
		{
			Slug:          `moderator`,
			Name:          `Moderator`,
			Description:   ``,
			IsSpeaker:     false,
			IsPresenter:   false,
			IsSponsor:     false,
			IsPanelist:    false,
			IsModerator:   true,
			IsHost:        false,
			IsParticipant: false,
			IsOther:       false,
		},
		{
			Slug:          `host`,
			Name:          `Host`,
			Description:   ``,
			IsSpeaker:     false,
			IsPresenter:   false,
			IsSponsor:     false,
			IsPanelist:    false,
			IsModerator:   false,
			IsHost:        true,
			IsParticipant: false,
			IsOther:       false,
		},
		{
			Slug:          `participant`,
			Name:          `Participant`,
			Description:   ``,
			IsSpeaker:     false,
			IsPresenter:   false,
			IsSponsor:     false,
			IsPanelist:    false,
			IsModerator:   false,
			IsHost:        false,
			IsParticipant: true,
			IsOther:       false,
		},
		{
			Slug:          `other`,
			Name:          `Other`,
			Description:   ``,
			IsSpeaker:     false,
			IsPresenter:   false,
			IsSponsor:     false,
			IsPanelist:    false,
			IsModerator:   false,
			IsHost:        false,
			IsParticipant: false,
			IsOther:       true,
		},
	}
}

func (f *Profile) ListProfileTypes() []ProfileType {
	return []ProfileType{
		{
			Slug:        `person`,
			Name:        `Person`,
			Description: ``,
			IsPerson:    true,
			IsCompany:   false,
			IsOther:     false,
		},
		{
			Slug:        `company`,
			Name:        `Company`,
			Description: ``,
			IsPerson:    false,
			IsCompany:   true,
			IsOther:     false,
		},
		{
			Slug:        `other`,
			Name:        `Other`,
			Description: ``,
			IsPerson:    false,
			IsCompany:   false,
			IsOther:     true,
		},
	}
}
