package requests

type Item struct {
	Questionnaire		int     `json:"Questionnaire"`
	QuestionnaireItem	int     `json:"QuestionnaireItem"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
