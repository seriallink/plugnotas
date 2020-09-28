package plugnotas

type ComercialTributavelString struct {
	Comercial  *string `json:"comercial,omitempty"`
	Tributavel *string `json:"tributavel,omitempty"`
}

type ComercialTributavelNumber struct {
	Comercial  *float64 `json:"comercial,omitempty"`
	Tributavel *float64 `json:"tributavel,omitempty"`
}
