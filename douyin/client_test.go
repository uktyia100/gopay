package douyin

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay"
)

var (
	client  *Client
	appId   = ""
	mchId   = ""
	apiKey  = ""
	salt    = ""
	baseUrl = "https://developer.toutiao.com"
	ctx     = context.Background()
	token   = ""
)

func TestMain(m *testing.M) {
	client = NewClient(appId, mchId, apiKey, false)
	client.DebugSwitch = gopay.DebugOn
	client.Salt = salt
	client.BaseURL = baseUrl
	client.SetToken(token)
	os.Exit(m.Run())
}
