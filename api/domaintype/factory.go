/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Factory struct {
}

func NewDomainTypeFactory() *Factory {
	return &Factory{}
}

func (f *Factory) ActivityLog() *ActivityLog {
	return NewActivityLog()
}

func (f *Factory) AppVersion() *AppVersion {
	return NewAppVersion()
}

func (f *Factory) EFx() *EFx {
	return NewEFx()
}

func (f *Factory) EmailDesign() *EmailDesign {
	return NewEmailDesign()
}

func (f *Factory) EmailMessage() *EmailMessage {
	return NewEmailMessage()
}

func (f *Factory) EmailTemplate() *EmailTemplate {
	return NewEmailTemplate()
}

func (f *Factory) Event() *Event {
	return NewEvent()
}

func (f *Factory) Group() *Group {
	return NewGroup()
}

func (f *Factory) HealthPass() *HealthPass {
	return NewHealthPass()
}

func (f *Factory) Import() *Import {
	return NewImport()
}

func (f *Factory) Integration() *Integration {
	return NewIntegration()
}

func (f *Factory) IntegrationFieldMapping() *IntegrationFieldMapping {
	return NewIntegrationFieldMapping()
}

func (f *Factory) IntegrationStatusMapping() *IntegrationStatusMapping {
	return NewIntegrationStatusMapping()
}

func (f *Factory) Invitation() *Invitation {
	return NewInvitation()
}

func (f *Factory) Lead() *Lead {
	return NewLead()
}

func (f *Factory) Link() *Link {
	return NewLink()
}

func (f *Factory) Mobile() *Mobile {
	return NewMobile()
}

func (f *Factory) Payment() *Payment {
	return NewPayment()
}

func (f *Factory) Pool() *Pool {
	return NewPool()
}

func (f *Factory) Profile() *Profile {
	return NewProfile()
}

func (f *Factory) Promotion() *Promotion {
	return NewPromotion()
}

func (f *Factory) Question() *Question {
	return NewQuestion()
}

func (f *Factory) Queue() *Queue {
	return NewQueue()
}

func (f *Factory) Region() *Region {
	return NewRegion()
}

func (f *Factory) Report() *Report {
	return NewReport()
}

func (f *Factory) Salesforce() *Salesforce {
	return NewSalesforce()
}

func (f *Factory) Salutation() *Salutation {
	return NewSalutation()
}

func (f *Factory) SitePage() *SitePage {
	return NewSitePage()
}

func (f *Factory) Stack() *Stack {
	return NewStack()
}

func (f *Factory) TicketBlock() *TicketBlock {
	return NewTicketBlock()
}

func (f *Factory) User() *User {
	return NewUser()
}

func (f *Factory) Venue() *Venue {
	return NewVenue()
}

func (f *Factory) Virbela() *Virbela {
	return NewVirbela()
}

func (f *Factory) WebConference() *WebConference {
	return NewWebConference()
}
