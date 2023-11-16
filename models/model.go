package models

// type RecordSimcards struct {
// 	Cliente        string `json:"idCostumer"`
// 	Numerochip     string `json:"idSimcard"`
// 	Simcon         string `json:"idSimcon"`
// 	Numerolinha    string `json:"idLine"`
// 	Numeroip       string `json:"idIP"`
// 	DateInstalacao string `json:"idDateinsta"`
// 	DataAtivacao   string `json:"idDateActi"`
// 	Fornecedor     string `json:"idSupplier"`
// 	Slotsimcon     string `json:"idSlot"`
// 	Plano          string `json:"idPlan"`
// 	Obs            string `json:"observations"`
// 	Operadora      string `json:"supplierType"`
// 	Apn            string `json:"idApn"`
// }

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
	Obs              *string `json:"obs"`
}
