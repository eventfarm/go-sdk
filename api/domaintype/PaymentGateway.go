/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type PaymentGateway struct {
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}

type PaymentGatewayType struct {
	Slug           string
	Name           string
	Description    string
	IsStripe       bool
	IsPaypal       bool
	IsAuthorizeNet bool
}

func (f *PaymentGateway) ListPaymentGatewayTypes() []PaymentGatewayType {
	return []PaymentGatewayType{
		{
			Slug:           `STRIPE`,
			Name:           `STRIPE`,
			Description:    `STRIPE`,
			IsStripe:       true,
			IsPaypal:       false,
			IsAuthorizeNet: false,
		},
		{
			Slug:           `PAYPAL`,
			Name:           `PAYPAL`,
			Description:    `PAYPAL`,
			IsStripe:       false,
			IsPaypal:       true,
			IsAuthorizeNet: false,
		},
		{
			Slug:           `AUTHORIZENET`,
			Name:           `AUTHORIZENET`,
			Description:    `AUTHORIZENET`,
			IsStripe:       false,
			IsPaypal:       false,
			IsAuthorizeNet: true,
		},
	}
}
