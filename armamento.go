package plugnotas

type Armamento struct {
	Tipo            *string `json:"tipo,omitempty"`
	NumeroSerie     *string `json:"numeroSerie,omitempty"`
	NumeroSerieCano *string `json:"numeroSerieCano,omitempty"`
	Descricao       *string `json:"descricao,omitempty"`
}
