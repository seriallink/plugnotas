package plugnotas

type Contribuinte struct {
	Id    string `json:"_id,omitempty"`
	Campo string `json:"campo,omitempty"`
	Texto string `json:"texto,omitempty"`
}
