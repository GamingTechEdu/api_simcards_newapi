package models

type GetListSimuc struct {
	Id             string  `json:"id"`
	Nserlum        string  `json:"nserlum"`
	Doc            string  `json:"doc"`
	NumberDoc      string  `json:"numberDoc"`
	Datedoc        string  `json:"datedoc"`
	User           string  `json:"user"`
	Violation      *string `json:"violation"`
	Obs            *string `json:"obs"`
	EvaluatorDate  *string `json:"evaluatorDate"`
	DefectRelated  *string `json:"defectrelated"`
	ApparentDefect *string `json:"apparentdefect"`
	Defect         *string `json:"defect"`
	Evaluator      *string `json:"evaluator"`
	Components     *string `json:"components"`
	Value          *string `json:"value"`
	Garantee       *string `json:"garantee"`
	Nivel          *string `json:"nivel"`
	Exist          *string `json:"exist"`
}
