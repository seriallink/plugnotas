package plugnotas

type Volumes struct {
	Especie     *string  `json:"especie,omitempty"`
	Marca       *string  `json:"marca,omitempty"`
	Quantidade  *float64 `json:"quantidade,omitempty"`
	Numeracao   *float64 `json:"numeracao,omitempty"`
	PesoLiquido *float64 `json:"pesoLiquido,omitempty"`
	PesoBruto   *float64 `json:"pesoBruto,omitempty"`
	Lacres      *Lacres  `json:"lacres,omitempty"`
}
