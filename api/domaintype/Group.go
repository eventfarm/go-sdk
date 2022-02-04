/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Group struct {
}

func NewGroup() *Group {
	return &Group{}
}

type EventGroupStatusFilterType struct {
	Slug                    string
	Name                    string
	Description             string
	IsHasInvitation         bool
	IsDoesNotHaveInvitation bool
}

func (f *Group) ListEventGroupStatusFilterTypes() []EventGroupStatusFilterType {
	return []EventGroupStatusFilterType{
		{
			Slug:                    `has_invitation`,
			Name:                    `Contacts included on the guest list`,
			Description:             ``,
			IsHasInvitation:         true,
			IsDoesNotHaveInvitation: false,
		},
		{
			Slug:                    `does_not_have_invitation`,
			Name:                    `Contacts absent from the guest list`,
			Description:             ``,
			IsHasInvitation:         false,
			IsDoesNotHaveInvitation: true,
		},
	}
}
