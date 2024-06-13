package douyin

import (
	"testing"

	"github.com/go-pay/gopay"
)

func TestClient_Code2Session(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("code", "hYjesl-vGcUfMbp9ZB7-mh9wD3zMHzUOnSzC6WPact3XLTSYbcbRbSeeMVwpsWHtx8dXF-JsUv2XvjBUhha5H5W3QQbUwV1FXRUfPRmKcc9jZ284gWAjy_yiIrY")
	resp, err := client.Code2Session(ctx, bm)
	if err != nil {
		t.Logf("fail：%v", err)
		return
	}
	t.Logf("resp:%+v, \nerr:%s", resp, err)
}
