/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Venue struct {
}

func NewVenue() *Venue {
	return &Venue{}
}

type VenueType struct {
	Slug        string
	Name        string
	Description string
	IsPhysical  bool
	IsVirbela   bool
	IsVirtual   bool
	IsOther     bool
}

func (f *Venue) ListVenueTypes() []VenueType {
	return []VenueType{
		{
			Slug:        `physical`,
			Name:        `Physical`,
			Description: ``,
			IsPhysical:  true,
			IsVirbela:   false,
			IsVirtual:   false,
			IsOther:     false,
		},
		{
			Slug:        `virbela`,
			Name:        `Virbela`,
			Description: ``,
			IsPhysical:  false,
			IsVirbela:   true,
			IsVirtual:   false,
			IsOther:     false,
		},
		{
			Slug:        `virtual`,
			Name:        `Virtual`,
			Description: ``,
			IsPhysical:  false,
			IsVirbela:   false,
			IsVirtual:   true,
			IsOther:     false,
		},
		{
			Slug:        `other`,
			Name:        `Other`,
			Description: ``,
			IsPhysical:  false,
			IsVirbela:   false,
			IsVirtual:   false,
			IsOther:     true,
		},
	}
}
