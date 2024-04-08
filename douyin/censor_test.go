package douyin

import (
	"testing"

	"github.com/go-pay/gopay"
)

func TestClient_CensorImage(t *testing.T) {
	ast, err := client.GetAccessToken(ctx)
	if err != nil {
		t.Log(err)
	}
	bm := make(gopay.BodyMap)
	bm.Set("access_token", ast.Data.AccessToken)
	bm.Set("image", "http://xxx.jpg")
	resp, err := client.CensorImage(ctx, bm)
	if err != nil {
		t.Logf("fail：%v", err)
		return
	}
	t.Logf("resp:%+v, \nerr:%s", resp, err)
}
