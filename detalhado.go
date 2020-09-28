package plugnotas

type Detalhado struct {
	Municipal *float64 `json:"municipal,omitempty"`
	Estadual  *float64 `json:"estadual,omitempty"`
	Federal   *float64 `json:"federal,omitempty"`
}
