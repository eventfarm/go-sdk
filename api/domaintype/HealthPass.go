/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type HealthPass struct {
}

func NewHealthPass() *HealthPass {
	return &HealthPass{}
}

type HealthPassScoreType struct {
	Slug         string
	Name         string
	Description  string
	IsCleared    bool
	IsNotCleared bool
	IsPending    bool
	IsUnknown    bool
}

func (f *HealthPass) ListHealthPassScoreTypes() []HealthPassScoreType {
	return []HealthPassScoreType{
		{
			Slug:         `green`,
			Name:         `Cleared`,
			Description:  ``,
			IsCleared:    true,
			IsNotCleared: false,
			IsPending:    false,
			IsUnknown:    false,
		},
		{
			Slug:         `red`,
			Name:         `Not Cleared`,
			Description:  ``,
			IsCleared:    false,
			IsNotCleared: true,
			IsPending:    false,
			IsUnknown:    false,
		},
		{
			Slug:         `amber`,
			Name:         `Pending`,
			Description:  ``,
			IsCleared:    false,
			IsNotCleared: false,
			IsPending:    true,
			IsUnknown:    false,
		},
		{
			Slug:         `unknown`,
			Name:         `Unknown`,
			Description:  ``,
			IsCleared:    false,
			IsNotCleared: false,
			IsPending:    false,
			IsUnknown:    true,
		},
	}
}
