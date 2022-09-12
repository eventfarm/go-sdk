/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Link struct {
}

func NewLink() *Link {
	return &Link{}
}

type LinkType struct {
	Slug        string
	Name        string
	Description string
	IsInstagram bool
	IsLinkedIn  bool
	IsTwitter   bool
	IsFacebook  bool
	IsSnapchat  bool
	IsWeb       bool
}

func (f *Link) ListLinkTypes() []LinkType {
	return []LinkType{
		{
			Slug:        `instagram`,
			Name:        `Instagram`,
			Description: ``,
			IsInstagram: true,
			IsLinkedIn:  false,
			IsTwitter:   false,
			IsFacebook:  false,
			IsSnapchat:  false,
			IsWeb:       false,
		},
		{
			Slug:        `linkedin`,
			Name:        `LinkedIn`,
			Description: ``,
			IsInstagram: false,
			IsLinkedIn:  true,
			IsTwitter:   false,
			IsFacebook:  false,
			IsSnapchat:  false,
			IsWeb:       false,
		},
		{
			Slug:        `twitter`,
			Name:        `Twitter`,
			Description: ``,
			IsInstagram: false,
			IsLinkedIn:  false,
			IsTwitter:   true,
			IsFacebook:  false,
			IsSnapchat:  false,
			IsWeb:       false,
		},
		{
			Slug:        `facebook`,
			Name:        `Facebook`,
			Description: ``,
			IsInstagram: false,
			IsLinkedIn:  false,
			IsTwitter:   false,
			IsFacebook:  true,
			IsSnapchat:  false,
			IsWeb:       false,
		},
		{
			Slug:        `snapchat`,
			Name:        `Snapchat`,
			Description: ``,
			IsInstagram: false,
			IsLinkedIn:  false,
			IsTwitter:   false,
			IsFacebook:  false,
			IsSnapchat:  true,
			IsWeb:       false,
		},
		{
			Slug:        `web`,
			Name:        `Web`,
			Description: ``,
			IsInstagram: false,
			IsLinkedIn:  false,
			IsTwitter:   false,
			IsFacebook:  false,
			IsSnapchat:  false,
			IsWeb:       true,
		},
	}
}
