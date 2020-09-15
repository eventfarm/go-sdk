/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Virbela struct {
}

func NewVirbela() *Virbela {
	return &Virbela{}
}

type VirbelaRoleType struct {
	Slug        string
	Name        string
	Description string
	IsMember    bool
	IsAdmin     bool
	IsSuspended bool
	IsModerator bool
}

func (f *Virbela) ListVirbelaRoleTypes() []VirbelaRoleType {
	return []VirbelaRoleType{
		{
			Slug:        `member`,
			Name:        `Member`,
			Description: ``,
			IsMember:    true,
			IsAdmin:     false,
			IsSuspended: false,
			IsModerator: false,
		},
		{
			Slug:        `moderator`,
			Name:        `Moderator`,
			Description: ``,
			IsMember:    false,
			IsAdmin:     false,
			IsSuspended: false,
			IsModerator: true,
		},
		{
			Slug:        `admin`,
			Name:        `Admin`,
			Description: ``,
			IsMember:    false,
			IsAdmin:     true,
			IsSuspended: false,
			IsModerator: false,
		},
		{
			Slug:        `suspended`,
			Name:        `Suspended`,
			Description: ``,
			IsMember:    false,
			IsAdmin:     false,
			IsSuspended: true,
			IsModerator: false,
		},
	}
}
