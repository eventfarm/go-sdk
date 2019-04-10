/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type EFx struct {
}

func NewEFx() *EFx {
	return &EFx{}
}

type EFxModuleType struct {
	Id                    string
	Slug                  string
	Name                  string
	Description           string
	IsAccessControl       bool
	IsActivation          bool
	IsConcierge           bool
	IsAthleteBag          bool
	IsDigitalMemoryBank   bool
	IsGuestInfo           bool
	IsMessaging           bool
	IsProductPickup       bool
	IsRaffle              bool
	IsReservation         bool
	IsRoamingPhotographer bool
	IsSmartBar            bool
	IsSmsQuiz             bool
	IsTeams               bool
	IsAdminOnly           bool
	IconURL               string
	DependencyNFC         bool
	DependencySMS         bool
	CompatibilityNFC      bool
	CompatibilitySMS      bool
}

func (f *EFx) ListEFxModuleTypes() []EFxModuleType {
	return []EFxModuleType{
		{
			Id:                    `AccessControl`,
			Slug:                  `access-control`,
			Name:                  `Access Control`,
			Description:           `Scan assets to grant access to exclusive event experiences`,
			IsAccessControl:       true,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/access.png`,
			DependencyNFC:         true,
			DependencySMS:         false,
			CompatibilityNFC:      true,
			CompatibilitySMS:      true,
		},
		{
			Id:                    `AthleteBag`,
			Slug:                  `athletes-bag`,
			Name:                  `Athletes Bag`,
			Description:           `Custom module for Nike`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          true,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           true,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/nike.png`,
			DependencyNFC:         true,
			DependencySMS:         false,
			CompatibilityNFC:      true,
			CompatibilitySMS:      false,
		},
		{
			Id:                    `Concierge`,
			Slug:                  `concierge`,
			Name:                  `Concierge`,
			Description:           `Deliver key information to attendees when they scan assets at kiosk`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           true,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/concierge.png`,
			DependencyNFC:         true,
			DependencySMS:         false,
			CompatibilityNFC:      true,
			CompatibilitySMS:      false,
		},
		{
			Id:                    `DigitalMemoryBank`,
			Slug:                  `digital-memory-bank`,
			Name:                  `Digital Memory Bank`,
			Description:           `Create and share a personalized post-event content hub with your attendees`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   true,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/digital-memory-bank.png`,
			DependencyNFC:         false,
			DependencySMS:         false,
			CompatibilityNFC:      true,
			CompatibilitySMS:      true,
		},
		{
			Id:                    `GuestInfo`,
			Slug:                  `guest-info`,
			Name:                  `Guest Info`,
			Description:           `Custom diagnostic tool`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           true,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           true,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/access.png`,
			DependencyNFC:         true,
			DependencySMS:         false,
			CompatibilityNFC:      true,
			CompatibilitySMS:      false,
		},
		{
			Id:                    `Messaging`,
			Slug:                  `messaging`,
			Name:                  `Messaging`,
			Description:           `Engage attendees via SMS messaging`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           true,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/messaging.png`,
			DependencyNFC:         false,
			DependencySMS:         true,
			CompatibilityNFC:      false,
			CompatibilitySMS:      true,
		},
		{
			Id:                    `SMSQuiz`,
			Slug:                  `smsquiz`,
			Name:                  `Poll`,
			Description:           `Text guests via SMS. View responses in EFx mobile app and Apple TV app`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             true,
			IsTeams:               false,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/polling.png`,
			DependencyNFC:         false,
			DependencySMS:         true,
			CompatibilityNFC:      false,
			CompatibilitySMS:      true,
		},
		{
			Id:                    `ProductPickup`,
			Slug:                  `product-pickup`,
			Name:                  `Product Pickup`,
			Description:           `Scan asset to confirm attendees&#039; eligibility to receive or purchase products`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       true,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/product-pickup.png`,
			DependencyNFC:         true,
			DependencySMS:         false,
			CompatibilityNFC:      true,
			CompatibilitySMS:      true,
		},
		{
			Id:                    `Raffle`,
			Slug:                  `raffle`,
			Name:                  `Raffle`,
			Description:           `Select prize winners and notify them via SMS`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              true,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/raffle.png`,
			DependencyNFC:         true,
			DependencySMS:         true,
			CompatibilityNFC:      true,
			CompatibilitySMS:      true,
		},
		{
			Id:                    `Reservation`,
			Slug:                  `reservation`,
			Name:                  `Reservation`,
			Description:           `Queue attendees in a virtual line and notify them via SMS when availability opens`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         true,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/reservation.png`,
			DependencyNFC:         true,
			DependencySMS:         false,
			CompatibilityNFC:      true,
			CompatibilitySMS:      true,
		},
		{
			Id:                    `RoamingPhotographer`,
			Slug:                  `roaming-photographer`,
			Name:                  `Roaming Photographer`,
			Description:           `Take photo and apply custom event filters. Deliver images via SMS`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: true,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/photographer.png`,
			DependencyNFC:         false,
			DependencySMS:         false,
			CompatibilityNFC:      true,
			CompatibilitySMS:      true,
		},
		{
			Id:                    `Smart Bar`,
			Slug:                  `smart-bar`,
			Name:                  `Smart Bar`,
			Description:           `Attendees scan assets at bar kiosks to place orders in bartender queue`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            true,
			IsSmsQuiz:             false,
			IsTeams:               false,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/smartbar-vertical.png`,
			DependencyNFC:         true,
			DependencySMS:         false,
			CompatibilityNFC:      true,
			CompatibilitySMS:      false,
		},
		{
			Id:                    `Teams`,
			Slug:                  `teams`,
			Name:                  `Teams `,
			Description:           `Build attendee teams for competition, collaboration, and networking; track scores on the leaderboard`,
			IsAccessControl:       false,
			IsActivation:          false,
			IsConcierge:           false,
			IsAthleteBag:          false,
			IsDigitalMemoryBank:   false,
			IsGuestInfo:           false,
			IsMessaging:           false,
			IsProductPickup:       false,
			IsRaffle:              false,
			IsReservation:         false,
			IsRoamingPhotographer: false,
			IsSmartBar:            false,
			IsSmsQuiz:             false,
			IsTeams:               true,
			IsAdminOnly:           false,
			IconURL:               `https://cdn.eventfarm.com/platform/img/icons/teams.png`,
			DependencyNFC:         false,
			DependencySMS:         false,
			CompatibilityNFC:      false,
			CompatibilitySMS:      true,
		},
	}
}
