package models

type Simcards struct {
	Id               string  `json:"id"`
	Client           string  `json:"client"`
	Iccid            string  `json:"iccid"`
	Simcon           string  `json:"simcon"`
	Msisdn           string  `json:"msisdn"`
	Ip               string  `json:"ip"`
	Slot             string  `json:"slot"`
	Installationdate string  `json:"installationdate"`
	Activationdate   string  `json:"activationdate"`
	Supplier         string  `json:"supplier"`
	Operator         string  `json:"operator"`
	Plan             string  `json:"plan"`
	Apn              string  `json:"apn"`
	Status           *string `json:"status"`
	Stock            *string `json:"stock"`
	Substituted      *string `json:"substitued"`
	Nfsimcon         *string `json:"nfsimcon"`
	Deliverydate     *string `json:"deliverydate"`
	Obs              *string `json:"obs"`
}

type Simucs struct {
	Id             *string `json:"id"`
	Nserlum        *string `json:"nserlum"`
	Doc            *string `json:"doc"`
	NumberDoc      *string `json:"numberDoc"`
	DateDoc        *string `json:"dateDoc"`
	User           *string `json:"user"`
	Violation      *string `json:"violation"`
	Obs            *string `json:"obs"`
	EvaluatorDate  *string `json:"evaluatorDate"`
	DefectRelated  *string `json:"defectRelated"`
	ApparentDefect *string `json:"apparentDefect"`
	Defect         *string `json:"defect"`
	Evaluator      *string `json:"evaluator"`
	Components     *string `json:"components"`
	Value          *string `json:"value"`
	Garantee       *string `json:"garantee"`
	Nivel          *string `json:"nivel"`
	Exist          *string `json:"exist"`
}

var SendUser struct {
	Username  string
	Password  string
	Hierarchy string `json:"hierarchy"`
}
