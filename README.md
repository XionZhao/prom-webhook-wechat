# Webhook adapter for Prometheus & Send Alert To Wechat Group chat

## Build

Just type and run: `make build`

Generated in the binary file The `./build` Dir

## Usage

```
usage: prom-webhook-wechat [<args>]


   -web.listen-address ":8060"
      Address to listen on for web interface.

 == WECHAT ==

   -wechat.apiurl
      Custom wechat api url

   -wechat.chatids_profile
      Custom chatid and profile (can specify multiple times,
      <profile>@<chatid>).

   -wechat.corpid
      wechat enterprise corpid.

   -wechat.corpsecret
      wechat app corpsecret.

   -wechat.timeout 5s
      Timeout for invoking wechat webhook.
```

## Exmaple

**Do not add to note that there is behind the token of the capacity(The program will get token by corpid and corpsecret)**

#### Start the single webhook and sent to a single group chat
```
./prom-webhook-wechat -wechat.corpid=CorpID -wechat.corpsecret=CorpSecret -wechat.chatids_profile=ops@CHAT_ID -wechat.apiurl=https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=
```
#### Start multiple webhook and sent to multiple group chat
```
./prom-webhook-wechat -wechat.corpid=CorpID -wechat.corpsecret=CorpSecret -wechat.chatids_profile=ops@CHAT_ID -wechat.chatids_profile=dev@CHAT_ID -wechat.apiurl=https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=
```

## Test request prom-webhook-wechat

To view `exmple/send_alert.sh`