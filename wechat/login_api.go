package wechat

import (
	"context"
	"fmt"

	"github.com/go-pay/xhttp"
)

type JsCode2SessionResp struct {
	Openid     string `json:"openid,omitempty"`
	SessionKey string `json:"session_key,omitempty"`
	Unionid    string `json:"unionid,omitempty"`
	Errcode    int    `json:"errcode,omitempty"` // 错误码
	Errmsg     string `json:"errmsg,omitempty"`  // 错误信息
}

// JsCode2Session 微信小程序登录，code 换取 open id
// appId：小程序 appId
// appSecret：小程序 appSecret
// jsCode：小程序用户静默登陆
// 文档：https://api.weixin.qq.com/sns/jscode2session
func JsCode2Session(ctx context.Context, appId, appSecret, jsCode string) (resp *JsCode2SessionResp, err error) {
	resp = new(JsCode2SessionResp)
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, appSecret, jsCode)
	_, err = xhttp.NewClient().Req().Get(url).EndStruct(ctx, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
