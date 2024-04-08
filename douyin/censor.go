package douyin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/content-security/picture-detect-v2
// 图片检测V2
func (u *Client) CensorImage(ctx context.Context, bm gopay.BodyMap) (resp *CensorImgResponse, err error) {
	var bs []byte
	if bs, err = u.doPost(ctx, "/api/apps/censor/image", bm); err != nil {
		return nil, err
	}
	resp = new(CensorImgResponse)
	if err = json.Unmarshal(bs, resp); err != nil {
		fmt.Printf("err:%s\n\n\n", err)
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return
}
