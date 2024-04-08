package douyin

import (
	"testing"

	"github.com/go-pay/gopay"
)

func TestClient_CreateOrder(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("out_order_no", "20240227170553916")
	bm.Set("total_amount", 1)
	bm.Set("subject", "测试")
	bm.Set("body", "测试抖音支付")
	bm.Set("valid_time", 180)
	bm.Set("notify_url", "https://show.world")
	t.Logf("bm:%+v", bm)
	resp, err := client.CreateOrder(ctx, bm)
	if err != nil {
		t.Logf("支付失败，%v", err)
	}
	t.Logf("resp:%+v, \nerr:%s", resp, err)
}
