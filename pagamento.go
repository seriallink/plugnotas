package plugnotas

type Pagamento struct {
	AVista        *bool    `json:"aVista,omitempty"`
	Meio          *string  `json:"meio,omitempty"`
	DescricaoMeio *string  `json:"descricaoMeio,omitempty"`
	Valor         *float64 `json:"valor,omitempty"`
	Cartao        *Cartao  `json:"cartao,omitempty"`
}
