package plugnotas

type Veiculo struct {
	Placa *string `json:"placa,omitempty"`
	UF    *string `json:"uf,omitempty"`
	RNTC  *string `json:"rntc,omitempty"`
}
