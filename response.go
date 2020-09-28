package plugnotas

type NfResponse struct {
	Message   *string            `json:"message"`
	Protocol  *string            `json:"protocol"`
	Data      *DataResponse      `json:"data"`
	Documents []DocumentResponse `json:"documents"`
}

func (n NfResponse) Error() string {
	panic("implement me")
}

type DataResponse struct {
	Id       *string `json:"id"`
	Protocol *string `json:"protocol"`
	Response *string `json:"response"`
}

type DocumentResponse struct {
	Id *string `json:"id"`
}
