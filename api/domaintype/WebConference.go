/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type WebConference struct {
}

func NewWebConference() *WebConference {
	return &WebConference{}
}

type WebConferenceFormatType struct {
	Slug        string
	Name        string
	Description string
	IsMeeting   bool
	IsWebinar   bool
	IsNone      bool
}

type WebConferenceSourceType struct {
	Slug        string
	Name        string
	Description string
	IsZoom      bool
	IsNone      bool
}

func (f *WebConference) ListWebConferenceFormatTypes() []WebConferenceFormatType {
	return []WebConferenceFormatType{
		{
			Slug:        `webinar`,
			Name:        `Webinar`,
			Description: ``,
			IsMeeting:   false,
			IsWebinar:   true,
			IsNone:      false,
		},
		{
			Slug:        `meeting`,
			Name:        `Meeting`,
			Description: ``,
			IsMeeting:   true,
			IsWebinar:   false,
			IsNone:      false,
		},
		{
			Slug:        `none`,
			Name:        `None`,
			Description: ``,
			IsMeeting:   false,
			IsWebinar:   false,
			IsNone:      true,
		},
	}
}

func (f *WebConference) ListWebConferenceSourceTypes() []WebConferenceSourceType {
	return []WebConferenceSourceType{
		{
			Slug:        `zoom`,
			Name:        `Zoom`,
			Description: ``,
			IsZoom:      true,
			IsNone:      false,
		},
		{
			Slug:        `none`,
			Name:        `None`,
			Description: ``,
			IsZoom:      false,
			IsNone:      true,
		},
	}
}
