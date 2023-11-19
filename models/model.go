package models

type Users struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Hierarchy string `json:"hierarchy"`
}

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
	Obs              *string `json:"obs"`
}
