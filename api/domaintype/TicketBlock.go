/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type TicketBlock struct {
}

func NewTicketBlock() *TicketBlock {
	return &TicketBlock{}
}

type TicketBlockType struct {
	Slug        string
	Name        string
	Description string
	IsStandard  bool
	IsSession   bool
}

func (f *TicketBlock) ListTicketBlockTypes() []TicketBlockType {
	return []TicketBlockType{
		{
			Slug:        `standard`,
			Name:        `Standard Ticket Block`,
			Description: ``,
			IsStandard:  true,
			IsSession:   false,
		},
		{
			Slug:        `session`,
			Name:        `Session Ticket Block`,
			Description: ``,
			IsStandard:  false,
			IsSession:   true,
		},
	}
}
