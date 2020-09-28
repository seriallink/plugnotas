package plugnotas

type ImpostoImportacao struct {
	BaseCalculo               *float64 `json:"baseCalculo,omitempty"`
	Valor                     *float64 `json:"valor,omitempty"`
	ValorDespesasAduaneiras   *float64 `json:"valorDespesasAduaneiras,omitempty"`
	ValorOperacoesFinanceiras *float64 `json:"valorOperacoesFinanceiras,omitempty"`
}
