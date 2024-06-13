package douyin

import (
	"context"
	"fmt"

	"github.com/go-pay/gopay"
)

// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/interface-request-credential/non-user-authorization/get-access-token
func (u *Client) GetAccessToken(ctx context.Context) (*AccessTokenData, error) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", u.AppId).
		Set("secret", u.ApiKey).
		Set("grant_type", "client_credential")

	resp := new(AccessTokenResp)
	if _, err := u.hc.Req().Post(u.BaseURL+"/api/apps/v2/token").SendBodyMap(bm).EndStruct(ctx, resp); err != nil {
		return nil, err
	}
	if resp.ErrCode != 0 {
		return nil, fmt.Errorf("%s[%d]", resp.ErrMsg, resp.ErrCode)
	}

	return resp.Data, nil
}
