/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Salesforce struct {
}

func NewSalesforce() *Salesforce {
	return &Salesforce{}
}

type CampaignMemberImportType struct {
	Slug                     string
	Name                     string
	Description              string
	IsNewMember              bool
	IsUpdatedMember          bool
	IsSkippedMemberAsRemoved bool
	IsDuplicateInvitation    bool
	IsDuplicateEmail         bool
}

type CampaignMemberType struct {
	Slug        string
	Name        string
	Description string
	IsContact   bool
	IsLead      bool
}

type NewContactRuleType struct {
	Slug            string
	Name            string
	Description     string
	IsDoNothing     bool
	IsCreateContact bool
	IsCreateLead    bool
}

func (f *Salesforce) ListCampaignMemberImportTypes() []CampaignMemberImportType {
	return []CampaignMemberImportType{
		{
			Slug:                     `new-member`,
			Name:                     `New Member`,
			Description:              ``,
			IsNewMember:              true,
			IsUpdatedMember:          false,
			IsSkippedMemberAsRemoved: false,
			IsDuplicateInvitation:    false,
			IsDuplicateEmail:         false,
		},
		{
			Slug:                     `updated-member`,
			Name:                     `Updated Member`,
			Description:              ``,
			IsNewMember:              false,
			IsUpdatedMember:          true,
			IsSkippedMemberAsRemoved: false,
			IsDuplicateInvitation:    false,
			IsDuplicateEmail:         false,
		},
		{
			Slug:                     `skipped-member-as-removed`,
			Name:                     `Skipped Member As Removed`,
			Description:              ``,
			IsNewMember:              false,
			IsUpdatedMember:          false,
			IsSkippedMemberAsRemoved: true,
			IsDuplicateInvitation:    false,
			IsDuplicateEmail:         false,
		},
		{
			Slug:                     `duplicate-invitation`,
			Name:                     `Duplicate Invitation`,
			Description:              ``,
			IsNewMember:              false,
			IsUpdatedMember:          false,
			IsSkippedMemberAsRemoved: false,
			IsDuplicateInvitation:    true,
			IsDuplicateEmail:         false,
		},
		{
			Slug:                     `duplicate-email`,
			Name:                     `Duplicate Email`,
			Description:              ``,
			IsNewMember:              false,
			IsUpdatedMember:          false,
			IsSkippedMemberAsRemoved: false,
			IsDuplicateInvitation:    false,
			IsDuplicateEmail:         true,
		},
	}
}

func (f *Salesforce) ListCampaignMemberTypes() []CampaignMemberType {
	return []CampaignMemberType{
		{
			Slug:        `contact`,
			Name:        `Contact`,
			Description: ``,
			IsContact:   true,
			IsLead:      false,
		},
		{
			Slug:        `lead`,
			Name:        `Lead`,
			Description: ``,
			IsContact:   false,
			IsLead:      true,
		},
	}
}

func (f *Salesforce) ListNewContactRuleTypes() []NewContactRuleType {
	return []NewContactRuleType{
		{
			Slug:            `do-nothing`,
			Name:            `Do Nothing`,
			Description:     `Do nothing. They will remain only in Event Farm.`,
			IsDoNothing:     true,
			IsCreateContact: false,
			IsCreateLead:    false,
		},
		{
			Slug:            `create-contact`,
			Name:            `Create Contact`,
			Description:     `Create a new Contact in Salesforce.`,
			IsDoNothing:     false,
			IsCreateContact: true,
			IsCreateLead:    false,
		},
		{
			Slug:            `create-lead`,
			Name:            `Create Lead`,
			Description:     `Create a new Lead in Salesforce.`,
			IsDoNothing:     false,
			IsCreateContact: false,
			IsCreateLead:    true,
		},
	}
}
