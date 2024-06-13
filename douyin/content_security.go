package douyin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/content-security/content-security-detect
// 内容安全检测
func (u *Client) TextAnti(ctx context.Context, bm gopay.BodyMap) (resp *TextAntiResponse, err error) {
	var bs []byte
	if bs, err = u.doPost(ctx, "/api/v2/tags/text/antidirt", bm); err != nil {
		return nil, err
	}
	resp = new(TextAntiResponse)
	if err = json.Unmarshal(bs, resp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return
}

// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/content-security/picture-detect-v2
// 图片检测 V2
func (u *Client) CensorImage(ctx context.Context, bm gopay.BodyMap) (resp *CensorImgResponse, err error) {
	var bs []byte
	if bs, err = u.doPost(ctx, "/api/apps/censor/image", bm); err != nil {
		return nil, err
	}
	resp = new(CensorImgResponse)
	if err = json.Unmarshal(bs, resp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return
}
