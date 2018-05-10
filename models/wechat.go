package models

type WechatNotificationResponse struct {
	ErrorMessage string `json:"errmsg"`
	ErrorCode    int    `json:"errcode"`
}

type WechatNotification struct {
	Text    *WechatNotificationText `json:"text,omitempty"`
	Chatid  string                  `json:"chatid"`
	Msgtype string                  `json:"msgtype"`
}

type WechatNotificationText struct {
	Content string `json:"content"`
}

type WechatNotificationButton struct {
	ActionURL string `json:"actionURL"`
}
