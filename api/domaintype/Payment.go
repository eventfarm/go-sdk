/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Payment struct {
}

func NewPayment() *Payment {
	return &Payment{}
}

type PurchaseErrorType struct {
	Slug                string
	Name                string
	Description         string
	IsCardDeclined      bool
	IsCreditCardInvalid bool
}

func (f *Payment) ListPurchaseErrorTypes() []PurchaseErrorType {
	return []PurchaseErrorType{
		{
			Slug:                `card_declined`,
			Name:                `Your card was declined. Your request was in live mode, but used a known test card.`,
			Description:         ``,
			IsCardDeclined:      true,
			IsCreditCardInvalid: false,
		},
		{
			Slug:                `2`,
			Name:                `Credit Card Invalid`,
			Description:         `The credit card number is invalid`,
			IsCardDeclined:      false,
			IsCreditCardInvalid: true,
		},
	}
}
