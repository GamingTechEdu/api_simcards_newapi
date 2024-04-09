package models

type Simcardstock struct {
	Id       string  `json:"id"`
	Iccid    string  `json:"iccid"`
	Supplier string  `json:"supplier"`
	Operator string  `json:"operator"`
	Plan     string  `json:"plan"`
	Apn      string  `json:"apn"`
	Status   *string `json:"status"`
	Obs      *string `json:"obs"`
}
