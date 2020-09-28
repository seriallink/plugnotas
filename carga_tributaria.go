package plugnotas

type CargaTributaria struct {
	Valor      *float64 `json:"valor,omitempty"`
	Percentual *float64 `json:"percentual,omitempty"`
	Fonte      *string  `json:"fonte,omitempty"`
}
