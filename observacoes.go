package plugnotas

type Observacoes struct {
	Contribuinte []Contribuinte `json:"contribuinte,omitempty"`
	Fisco        interface{}    `json:"fisco,omitempty"`
}
