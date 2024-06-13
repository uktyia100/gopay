package douyin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

type Code2SessionData struct {
	SessionKey      string `json:"session_key,omitempty"`
	OpenId          string `json:"openid,omitempty"`
	AnonymousOpenId string `json:"anonymous_openid,omitempty"`
	UnionId         string `json:"unionid,omitempty"`
}

type Code2SessionResponse struct {
	ErrNo   int               `json:"err_no,omitempty"`   // 执行结果
	ErrTips string            `json:"err_tips,omitempty"` // 返回错误信息
	Data    *Code2SessionData `json:"data,omitempty"`
}

// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/log-in/code-2-session
// 小程序免登
func (u *Client) Code2Session(ctx context.Context, bm gopay.BodyMap) (resp *Code2SessionData, err error) {
	var bs []byte
	bm.Set("appid", u.AppId)
	bm.Set("secret", u.ApiKey)
	if bs, err = u.doPost(ctx, "/api/apps/v2/jscode2session", bm); err != nil {
		return nil, err
	}
	ulResp := new(Code2SessionResponse)
	if err = json.Unmarshal(bs, ulResp); err != nil || ulResp.ErrNo != 0 {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return ulResp.Data, nil
}
