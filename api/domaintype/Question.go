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
