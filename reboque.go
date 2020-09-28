package plugnotas

type Reboque struct {
	Placa *string `json:"placa,omitempty"`
	UF    *string `json:"uf,omitempty"`
	RNTC  *string `json:"rntc,omitempty"`
	Vagao *string `json:"vagao,omitempty"`
	Balsa *string `json:"balsa,omitempty"`
}
