package plugnotas

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type Certificado struct {
	Id         string `json:"id"`
	Nome       string `json:"nome"`
	Hash       string `json:"hash"`
	Vencimento string `json:"vencimento"`
	Email      string `json:"email"`
}

func (c *Client) ListaCertificados() (certificados []Certificado, err error) {
	err = c.Get("certificado/", nil, nil, &certificados)
	return
}

func (c *Client) UploadCertificate(arquivo, senha, email string) (id string, err error) {

	var (
		file *os.File
		part io.Writer
		resp *http.Response
		body []byte
		gmsg GenericMessage
	)

	// init request body
	payload := &bytes.Buffer{}

	// create form writer
	writer := multipart.NewWriter(payload)

	// open file
	if file, err = os.Open(arquivo); err != nil {
		return
	}

	// close file
	defer file.Close()

	// set file
	if part, err = writer.CreateFormFile("arquivo", filepath.Base(file.Name())); err != nil {
		return
	}

	// write file
	if _, err = io.Copy(part, file); err != nil {
		return
	}

	// set password
	if err = writer.WriteField("senha", senha); err != nil {
		return
	}

	// set email
	if err = writer.WriteField("email", email); err != nil {
		return
	}

	// close writer
	if err = writer.Close(); err != nil {
		return
	}

	// set request
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%scertificado", c.GetEndpoint()), payload)
	req.Header.Add("x-api-key", c.GetAuthorization())
	req.Header.Add("content-type", "multipart/form-data")
	req.Header.Set("content-type", writer.FormDataContentType())

	// do request
	if resp, err = http.DefaultClient.Do(req); err != nil {
		return
	}

	// read response
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	// close response
	defer resp.Body.Close()

	// parse response
	if err = json.Unmarshal(body, &gmsg); err != nil {
		return
	}

	// verify status code
	if NotIn(resp.StatusCode, http.StatusOK, http.StatusCreated, http.StatusAccepted) {

		// is message an error?
		if gmsg.Error() != "" {

			// return error message
			return "", gmsg

		}

	}

	return gmsg.Data.(map[string]interface{})["id"].(string), nil

}
