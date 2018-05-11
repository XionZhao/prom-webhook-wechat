package webrouter

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/3Golds/prom-webhook-wechat/models"
	"github.com/3Golds/prom-webhook-wechat/notifier"
	"github.com/3Golds/prom-webhook-wechat/request"
	"github.com/pressly/chi"
)

type WechatResource struct {
	Profileurl string
	HttpClient *http.Client
	Chatids    map[string]string
	Corpid     string
	Corpsecret string
}

func (rs *WechatResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/{profile}/send", rs.SendNotification)
	return r
}

func (rs *WechatResource) SendNotification(w http.ResponseWriter, r *http.Request) {
	profile := chi.URLParam(r, "profile")
	apiurl := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + rs.Corpid + "&corpsecret=" + rs.Corpsecret
	getTokenResp, err := request.SendGetTokenRequest(apiurl)
	if err != nil {
		log.Panicf("Failed to request: %s", err)
	}

	webhookURL := rs.Profileurl + getTokenResp.AccessToken

	var promMessage models.WebhookMessage
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&promMessage); err != nil {
		log.Printf("Cannot decode prometheus webhook JSON request: %s", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	notification, err := notifier.BuildWechatNotification(&promMessage, rs.Chatids[profile])
	if err != nil {
		log.Printf("Failed to build notification: %s", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	robotResp, err := notifier.SendWechatNotification(rs.HttpClient, webhookURL, notification)
	if err != nil {
		log.Printf("Failed to send notification: %s", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

	if robotResp.ErrorCode != 0 {
		log.Printf("Failed to send notification to wechat: [%d] %s", robotResp.ErrorCode, robotResp.ErrorMessage)
		return
	}

	log.Println("Successfully send notification to wechat")
}
