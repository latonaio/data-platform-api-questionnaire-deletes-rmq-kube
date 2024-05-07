package requests

type Header struct {
	Questionnaire		int     `json:"Questionnaire"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
