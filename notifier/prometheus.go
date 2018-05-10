package notifier

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/3Golds/prom-webhook-wechat/models"
	"github.com/3Golds/prom-webhook-wechat/tpl"
	"github.com/pkg/errors"
)

func BuildWechatNotification(promMessage *models.WebhookMessage, chatid string) (*models.WechatNotification, error) {
	content, err := tpl.ExecuteTextString(`{{ template "ding.link.content" . }}`, promMessage)
	if err != nil {
		return nil, err
	}
	var buttons []models.WechatNotificationButton
	for _, alert := range promMessage.Alerts.Firing() {
		buttons = append(buttons, models.WechatNotificationButton{
			ActionURL: alert.GeneratorURL,
		})
	}

	notification := &models.WechatNotification{
		Chatid:  chatid,
		Msgtype: "text",
		Text: &models.WechatNotificationText{
			Content: content,
		},
	}
	return notification, nil
}

func SendWechatNotification(httpClient *http.Client, ApiURL string, notification *models.WechatNotification) (*models.WechatNotificationResponse, error) {
	body, err := json.Marshal(&notification)
	if err != nil {
		return nil, errors.Wrap(err, "error encoding Wechat request")
	}

	httpReq, err := http.NewRequest("POST", ApiURL, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "error building Wechat request")
	}
	httpReq.Header.Set("Content-Type", "application/json")

	req, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "error sending notification to Wechat")
	}
	defer req.Body.Close()

	if req.StatusCode != 200 {
		return nil, errors.Errorf("unacceptable response code %d", req.StatusCode)
	}

	var robotResp models.WechatNotificationResponse
	enc := json.NewDecoder(req.Body)
	if err := enc.Decode(&robotResp); err != nil {
		return nil, errors.Wrap(err, "error decoding response from DingTalk")
	}

	return &robotResp, nil
}
