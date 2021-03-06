/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Salutation struct {
}

func NewSalutation() *Salutation {
	return &Salutation{}
}

type SalutationLocaleType struct {
	Slug                  string
	Name                  string
	Description           string
	IsEnglish             bool
	IsGerman              bool
	IsFrench              bool
	IsHebrew              bool
	IsPolish              bool
	IsPortuguese          bool
	IsSpanish             bool
	IsSpanishSouthAmerica bool
	IsThai                bool
	IsItalian             bool
	IsChineseTraditional  bool
	IsChineseMandarin     bool
	IsJapanese            bool
	IsKorean              bool
}

func (f *Salutation) ListSalutationLocaleTypes() []SalutationLocaleType {
	return []SalutationLocaleType{
		{
			Slug:                  `english`,
			Name:                  `English`,
			Description:           ``,
			IsEnglish:             true,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `german`,
			Name:                  `German`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              true,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `french`,
			Name:                  `French`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              true,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `hebrew`,
			Name:                  `Hebrew`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              true,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `polish`,
			Name:                  `Polish`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              true,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `portuguese`,
			Name:                  `Portuguese`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          true,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `spanish`,
			Name:                  `Spanish`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             true,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `spanish-south-america`,
			Name:                  `Spanish (South America)`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: true,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `thai`,
			Name:                  `Thai`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                true,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `italian`,
			Name:                  `Italian`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             true,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `chinese-traditional`,
			Name:                  `Chinese (Tradition)`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  true,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `chinese-mandarin`,
			Name:                  `Chinese (Mandarin)`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     true,
			IsJapanese:            false,
			IsKorean:              false,
		},
		{
			Slug:                  `japanese`,
			Name:                  `Japanese`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            true,
			IsKorean:              false,
		},
		{
			Slug:                  `korean`,
			Name:                  `Korean`,
			Description:           ``,
			IsEnglish:             false,
			IsGerman:              false,
			IsFrench:              false,
			IsHebrew:              false,
			IsPolish:              false,
			IsPortuguese:          false,
			IsSpanish:             false,
			IsSpanishSouthAmerica: false,
			IsThai:                false,
			IsItalian:             false,
			IsChineseTraditional:  false,
			IsChineseMandarin:     false,
			IsJapanese:            false,
			IsKorean:              true,
		},
	}
}
