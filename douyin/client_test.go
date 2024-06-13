package douyin

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay"
)

var (
	client  *Client
	appId   = "tt310b47b6ae3a956101"
	mchId   = "73297614642424076800"
	apiKey  = "c4852ecc419d2be0d4798524927686fdb2d30b95"
	salt    = "8QtBfA4WNKI3FpYcMB19FTMxNakE1k99BXg27Vco"
	baseUrl = "https://developer.toutiao.com"
	ctx     = context.Background()
	token   = "doodo.cn"
)

func TestMain(m *testing.M) {
	client = NewClient(appId, mchId, apiKey, false)
	client.DebugSwitch = gopay.DebugOn
	client.Salt = salt
	client.BaseURL = baseUrl
	client.SetToken(token)
	os.Exit(m.Run())
}
