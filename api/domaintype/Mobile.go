/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Mobile struct {
}

func NewMobile() *Mobile {
	return &Mobile{}
}

type MobileDTOType struct {
	Slug          string
	Name          string
	Description   string
	IsInvitations bool
	IsEvents      bool
}

func (f *Mobile) ListMobileDTOTypes() []MobileDTOType {
	return []MobileDTOType{
		{
			Slug:          `invitations`,
			Name:          `Invitations`,
			Description:   ``,
			IsInvitations: true,
			IsEvents:      false,
		},
		{
			Slug:          `events`,
			Name:          `Events`,
			Description:   ``,
			IsInvitations: false,
			IsEvents:      true,
		},
	}
}
