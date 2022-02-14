package form3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(accountAPIBaseURL string) *Client {
	return &Client{
		BaseURL: accountAPIBaseURL,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) GetBufferedStreamForAccount(accountMesage *AccountMessage) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(accountMesage)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func (c *Client) GenerateUniqueID() (string, error) {

	id := uuid.New()

	fmt.Println(id.String())

	uuidString := id.String()

	return uuidString, nil
}

func (c *Client) PopulateAccountMesage(bankId string, bankIdCode string, baseCurrency string, bic string, country string, name []string, organisationId string, id string) (*AccountMessage, error) {

	accountAttributes := AccountAttributes{
		BankID:       bankId,
		BankIDCode:   bankIdCode,
		BaseCurrency: baseCurrency,
		Bic:          bic,
		Country:      country,
		Name:         name,
	}
	accountData := AccountData{
		Attributes:     accountAttributes,
		ID:             id,
		OrganisationID: organisationId,
		Type:           "accounts",
	}

	acccountMessage := AccountMessage{
		Data: accountData,
	}

	return &acccountMessage, nil
}
