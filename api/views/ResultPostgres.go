package views

type ResultPostgres struct {
	Imei       *string `json:"Imei"`
	Plate      *string `json:"Plate"`
	FkCompany  *int    `json:"FkCompany"`
	Company    *string `json:"Company"`
	FkCustomer *int    `json:"FkCustomer"`
	Customer   *string `json:"Customer"`
}
