/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Lead struct {
}

func NewLead() *Lead {
	return &Lead{}
}

type LeadTemperatureType struct {
	Slug        string
	Name        string
	Description string
	IsCold      bool
	IsWarm      bool
	IsHot       bool
	IsNotSet    bool
}

func (f *Lead) ListLeadTemperatureTypes() []LeadTemperatureType {
	return []LeadTemperatureType{
		{
			Slug:        `not-set`,
			Name:        `Not Set`,
			Description: ``,
			IsCold:      false,
			IsWarm:      false,
			IsHot:       false,
			IsNotSet:    true,
		},
		{
			Slug:        `cold`,
			Name:        `Cold`,
			Description: ``,
			IsCold:      true,
			IsWarm:      false,
			IsHot:       false,
			IsNotSet:    false,
		},
		{
			Slug:        `warm`,
			Name:        `Warm`,
			Description: ``,
			IsCold:      false,
			IsWarm:      true,
			IsHot:       false,
			IsNotSet:    false,
		},
		{
			Slug:        `hot`,
			Name:        `Hot`,
			Description: ``,
			IsCold:      false,
			IsWarm:      false,
			IsHot:       true,
			IsNotSet:    false,
		},
	}
}
