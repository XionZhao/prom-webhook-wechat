package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/3Golds/prom-webhook-wechat/models"
	"github.com/pkg/errors"
)

// SendGetTokenRequest 获取token
func SendGetTokenRequest(ApiURL string) (*models.GetTokenResponse, error) {
	resp, err := http.Get(ApiURL)
	if err != nil {
		return nil, errors.Wrap(err, "error request Wechat request")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error IO Read")
	}

	var tokenResp models.GetTokenResponse
	err = json.Unmarshal([]byte(string(body)), &tokenResp)
	if err != nil {
		return nil, errors.Wrap(err, "error Unmarshal")
	}

	return &tokenResp, nil
}
