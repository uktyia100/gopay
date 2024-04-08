package douyin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 支付结果回调
// 文档：https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/pay-list/callback
func ParseNotify(req *http.Request) (*NotifyMsgResp, error) {
	bs, err := io.ReadAll(io.LimitReader(req.Body, int64(5<<20))) // default 5MB change the size you want;
	defer req.Body.Close()
	if err != nil {
		return nil, err
	}

	notifyResp := new(NotifyRequest)
	if err = json.Unmarshal(bs, notifyResp); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(bs), notifyResp, err)
	}
	resp := new(NotifyMsgResp)
	if err = json.Unmarshal([]byte(notifyResp.Msg), resp); err != nil {
		return nil, fmt.Errorf("第二次解析msg失败：json.Unmarshal(%s, %+v)：%w", notifyResp.Msg, resp, err)
	}
	return resp, nil
}
