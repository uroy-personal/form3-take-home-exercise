package form3

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("Form3API-Testing")
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

func (c *Client) FetchAccount(ctx context.Context, id string) (*AccountMessage, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/organisation/accounts/%s", c.BaseURL, id), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := AccountMessage{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	log.Info("%+v\n", res)
	return &res, nil
}

func (c *Client) CreateAccount(ctx context.Context, accountMesage *AccountMessage) (*AccountMessage, error) {
	buf, err := c.GetBufferedStreamForAccount(accountMesage)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/organisation/accounts/", c.BaseURL), buf)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := AccountMessage{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	log.Info("%+v\n", res)
	return &res, nil
}

func (c *Client) DeleteAccount(ctx context.Context, id string) (*string, error) {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/organisation/accounts/%s?version=0", c.BaseURL, id), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := ""

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	log.Info("Delete-Response %+v\n", res)
	return &res, nil
}
func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.Unmarshal([]byte(bodyBytes), &v); err != nil {
		return err
	}

	return nil
}
