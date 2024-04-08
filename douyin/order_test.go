package douyin

import (
	"strconv"
	"testing"
	"time"

	"github.com/go-pay/gopay"
)

func TestClient_RequestOrder(t *testing.T) {
	requestData := &RequestOrderData{
		SkuList: []*RequestOrderSkuList{
			{
				SkuId:      "1",
				Price:      9900,
				Quantity:   1,
				Title:      "title",
				ImageList:  []string{"http://xxx"},
				Type:       401,
				TagGroupId: "tag_group_7272625659888058380",
			},
		},
		OutOrderNo:  "aa",
		TotalAmount: 1,
		OrderEntrySchema: &RequestOrderSchema{
			Path: "page/path/index",
		},
	}
	privateKeyStr := `
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCfDYrroPj5oLBR
/xtb+I+ooMoiUpsL/bBDCLYUwjOu7s0petoUo68S7J78E88X+poO3/SYrBQZDsCV
4b5eUL2M/qTZB42THVPLG42DxdRZa+wHoL/2DCX/mntfMjFccSyymne1yPMJdSvY
VVmmGqxysrE5IOi/KmplArIf4qPCbQGhCg5hDceNs/EVfeDGYe+9Vc+BDBciYbo8
kkBvPi2hiA+v7t2anGPXWD6rHYOw1QLSsXhLQhc3ujgQij/tz3wKcc4nsIf1L2DY
UGnMhsp6P+iXVrgOFe8CWicoS354jM57ucxo6xleTnl/WKXZZsRpTP/01edwop/x
vzZVykGbAgMBAAECggEAdzmJlxM1TrnbMbvO4GQ0G61bl9rgCl8CD12qT2k7oLe0
5pEGhE4mYEMq8b5PkKPzc48BJKho8FeUmUV8k+pmDBidrWnyAAMN2sQukkZq5RT2
+sjO3DtDCJTeQf+37JmRbdhIP+X/+Gjyktl3uAFqnxe/rRk0HohG9KzOpQQYyRYK
ix0HP1ze3D5XKoYnAxABdHyxNUuNQSqtSCPta/JT6QGkMzk+WwCJTvckEqo9piQH
8jCphQjl4zQWFQK+WWPX54Z8176oUXGMLCnqYx/HYGGJAoZlQNoadbvXARrDKCNI
Jzv8GYjddBNU68ZCzU/W1PbhElbDVKZ9pRu1D1WXUQKBgQDKoDGQ332pDiMj/6uA
IdkelZTN22tsTUCHxUiJMJJeUryavGbngxYBNfDbtzaUZuN0chMZ+WOnZZA8wH18
8lTpVus8YT9wR9EErRTfkQcvelxAP5eRGWqSPrIec+YhO9AW1G3g8EYzm2KSz3JR
d9XcFXH/o06bKJXPYYyUZiZupQKBgQDI8xHcy0dHAOYMFTMztNRVW6qBrFCPQG/R
2R1Z0jAev8YFRevcLlr0VcyMmcz7WxGimxAjckzwOL/okWHXW2SjVp7MkMJ9SFOF
4UQSp0CkePWA56jEOE1g9HDoQWjFWaK2aRtQcZ/+clfxUogcL9v8+D77JIltApj7
Jn/ZgHE7PwKBgEq7OGyxMNxn/WfqhOs2EKjqDD68XWtNNq6cgXsvsdwd8be0ItPi
EfySU2oFsZicNemdpRPgWfPETqVJbT7m9ZA14X6cc1RK+HkcCOXHzHmjTGDZ02HN
mPOlMfZvUoIRWFcDNB9RHuMrvPhekFAhvXt1YV0icvxgxJf+52VYI3aNAoGAKBcu
giFgKA9K1jBIldFG53yxSMurCtltNa2eSHRBu3DBPf4UL0pkRQj1FTv+BEvH+ev7
zUaZiPeZefm/Tmriah/28JMU6k/KshQeM6aApA+p0zzkk7kz3tsFx6B9GZpndwMD
OwpgRDBl0TSJCS6XMiII1qwUatYw1TJ6IVeZv5sCgYA9hyyQ/4Pv/dPbmu579W6A
2n/fBpw2SRhqc/TaiQEww3RoiQR2W91Z/1r4PXAANJ4cCVr1mwFrtzUUv3FI7pyK
k/jkGMJBV5F/iwr9e3ikf/8PdxiDTOQTeDY2bUxlz1RhajDSVekRb8Mt8N3tRZfS
puXwZN026m5SgNgswvI1Zw==
`
	byteAuth, err := client.GetByteAuthorization(requestData, privateKeyStr, strconv.FormatInt(time.Now().Unix(), 10), "1")
	t.Log(requestData, byteAuth, err)
}

func TestClient_PushOrder(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("access_token", "a").
		Set("app_name", "douyin").
		Set("open_id", "douyin").
		SetBodyMap("order_detail", func(b gopay.BodyMap) {
			b.Set("order_id", "20240227170553916"). // 开发者侧业务单号
								Set("create_time", 1). // 订单创建的时间,13 位毫秒时间戳
								Set("status", "已支付").
								Set("amount", 1).                             // 订单商品总数
								Set("total_price", 100).                      // 订单总价，单位为分
								Set("detail_url", "pages/order/orderDetail"). // 小程序订单详情页 path
								SetBodyMap("item_list", func(b gopay.BodyMap) {
					b.Set("item_code", ""). // 开发者侧商品 ID，长度 <= 64 byte
								Set("img", "").   // 子订单商品图片 URL
								Set("title", ""). // 子订单商品介绍标题
								Set("amount", 1). // 单类商品的数目
								Set("price", 1)   // 单类商品的总价
				}) // 子订单商品列表，不可为空
		}).
		Set("order_status", 2).
		Set("order_type", 0).
		Set("update_time", 123) //订单信息变更时间，10 位秒级时间戳
	t.Logf("bm:%+v", bm)
	resp, err := client.PushOrder(ctx, bm)
	if err != nil {
		t.Logf("支付失败，%v", err)
	}
	t.Logf("resp:%+v, \nerr:%s", resp, err)
}
