/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Question struct {
}

func NewQuestion() *Question {
	return &Question{}
}

type AnswerBindingType struct {
	Slug         string
	Name         string
	Description  string
	IsEqualTo    bool
	IsNotEqualTo bool
}

type QuestionContextType struct {
	Slug           string
	Name           string
	Description    string
	IsRegistration bool
	IsLeadCapture  bool
}

type QuestionType struct {
	Slug        string
	Name        string
	Description string
	IsCheckbox  bool
	IsRadio     bool
	IsMulti     bool
	IsText      bool
	IsSelect    bool
	IsDate      bool
	IsWaiver    bool
	IsAddress   bool
}

func (f *Question) ListAnswerBindingTypes() []AnswerBindingType {
	return []AnswerBindingType{
		{
			Slug:         `equal_to`,
			Name:         `Equal To`,
			Description:  ``,
			IsEqualTo:    true,
			IsNotEqualTo: false,
		},
		{
			Slug:         `not_equal_to`,
			Name:         `Not Equal To`,
			Description:  ``,
			IsEqualTo:    false,
			IsNotEqualTo: true,
		},
	}
}

func (f *Question) ListQuestionContextTypes() []QuestionContextType {
	return []QuestionContextType{
		{
			Slug:           `registration`,
			Name:           `Registration`,
			Description:    ``,
			IsRegistration: true,
			IsLeadCapture:  false,
		},
		{
			Slug:           `lead`,
			Name:           `Lead Capture`,
			Description:    ``,
			IsRegistration: false,
			IsLeadCapture:  true,
		},
	}
}

func (f *Question) ListQuestionTypes() []QuestionType {
	return []QuestionType{
		{
			Slug:        `checkbox`,
			Name:        `Checkboxes`,
			Description: ``,
			IsCheckbox:  true,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    false,
			IsAddress:   false,
		},
		{
			Slug:        `radio`,
			Name:        `Radio Buttons`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     true,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    false,
			IsAddress:   false,
		},
		{
			Slug:        `multi`,
			Name:        `Paragraph`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     true,
			IsText:      false,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    false,
			IsAddress:   false,
		},
		{
			Slug:        `text`,
			Name:        `Short Answer`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      true,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    false,
			IsAddress:   false,
		},
		{
			Slug:        `select`,
			Name:        `Dropdown Select`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    true,
			IsDate:      false,
			IsWaiver:    false,
			IsAddress:   false,
		},
		{
			Slug:        `date`,
			Name:        `Select Date`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    false,
			IsDate:      true,
			IsWaiver:    false,
			IsAddress:   false,
		},
		{
			Slug:        `waiver`,
			Name:        `Sign Waiver`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    true,
			IsAddress:   false,
		},
		{
			Slug:        `address`,
			Name:        `Address`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    false,
			IsAddress:   true,
		},
	}
}
