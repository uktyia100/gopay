package douyin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-pay/gopay"
)

// *********** 担保支付相关 *************

// 预下单接口
// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/pay-list/pay
func (u *Client) CreateOrder(ctx context.Context, bm gopay.BodyMap) (createOrderData *CreateOrderData, err error) {
	if err = bm.CheckEmptyError("out_order_no", "total_amount", "subject", "body"); err != nil {
		return nil, err
	}
	var bs []byte
	bm.Set("notify_url", u.NotifyUrl)
	if bs, err = u.doPost(ctx, "/api/apps/ecpay/v1/create_order", bm); err != nil {
		return nil, err
	}
	ulRsp := new(CreateOrderResponse)
	if err = json.Unmarshal(bs, ulRsp); err != nil || ulRsp.ErrNo != 0 {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}

	return ulRsp.Data, nil
}

// 支付结果查询
// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/pay-list/query
func (u *Client) QueryOrder(ctx context.Context, bm gopay.BodyMap) (createOrderData *CreateOrderData, err error) {
	if err = bm.CheckEmptyError("out_order_no"); err != nil {
		return nil, err
	}
	var bs []byte
	bm.Set("notify_url", u.NotifyUrl)
	if bs, err = u.doPost(ctx, "/api/apps/ecpay/v1/create_order", bm); err != nil {
		return nil, err
	}
	ulRsp := new(CreateOrderResponse)
	if err = json.Unmarshal(bs, ulRsp); err != nil || ulRsp.ErrNo != 0 {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}

	return ulRsp.Data, nil
}

// 订单同步
// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/ecpay/order/order-sync/
func (u *Client) PushOrder(ctx context.Context, bm gopay.BodyMap) (resp *PushOrderResponse, err error) {
	var bs []byte
	if bs, err = u.doPost(ctx, "/api/apps/order/v2/push", bm); err != nil {
		return nil, err
	}
	ulRsp := new(PushOrderResponse)
	if err = json.Unmarshal(bs, ulRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}

	return ulRsp, nil
}
