package douyin

import (
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"hash"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xhttp"
	"github.com/go-pay/xlog"
)

// Client douyin
type Client struct {
	AppId       string
	MchId       string
	ApiKey      string
	BaseURL     string
	NotifyUrl   string
	Salt        string            // 盐值
	Token       string            // 支付设置 token
	IsProd      bool              // 是否生产环境
	ctx         context.Context   // 上下文
	DebugSwitch gopay.DebugSwitch // 调试开关，是否打印日志
	mu          sync.RWMutex
	sha256Hash  hash.Hash
	md5Hash     hash.Hash
	hc          *xhttp.Client
	tlsHc       *xhttp.Client
}

// NewClient 初始化抖音支付客户端
// appId：应用ID
// mchId：商户ID
// ApiKey：API秘钥值
// IsProd：是否是正式环境
func NewClient(appId, mchId, apiKey string, isProd bool) (client *Client) {
	return &Client{
		AppId:       appId,
		MchId:       mchId,
		ApiKey:      apiKey,
		IsProd:      isProd,
		ctx:         context.Background(),
		DebugSwitch: gopay.DebugOff,
		sha256Hash:  hmac.New(sha256.New, []byte(apiKey)),
		md5Hash:     md5.New(),
		hc:          xhttp.NewClient(),
		tlsHc:       xhttp.NewClient(),
	}
}

// SetBodySize 设置http response body size(MB)
func (c *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.hc.SetBodySize(sizeMB)
	}
}

func (c *Client) SetToken(token string) {
	c.Token = token
}

// 担保支付回调签名算法
// 参数："strArr" 所有字段（验证时注意不包含 sign 签名本身，不包含空字段与 type 常量字段）内容与平台上配置的 token
func (c *Client) VerifySign(notifyReq *NotifyRequest) (err error) {
	strArr := []string{c.Token, notifyReq.Timestamp, notifyReq.Nonce, notifyReq.Msg}
	sort.Strings(strArr)
	h := sha1.New()
	h.Write([]byte(strings.Join(strArr, "")))
	validSign := fmt.Sprintf("%x", h.Sum(nil))

	if notifyReq.MsgSignature != validSign {
		return fmt.Errorf("签名验证失败")
	}
	return
}

// 获取签名字符串
func (c *Client) getRsaSign(bm gopay.BodyMap) (sign string, err error) {
	var paramsArr []string
	if c.Salt == "" {
		return "", fmt.Errorf("签名缺少必要的参数")
	}
	for k, v := range bm {
		if k == "other_settle_params" || k == "app_id" || k == "thirdparty_id" || k == "sign" {
			continue
		}
		value := strings.TrimSpace(fmt.Sprintf("%v", v))
		if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") && len(value) > 1 {
			value = value[1 : len(value)-1]
		}
		value = strings.TrimSpace(value)
		if value == "" || value == "null" {
			continue
		}
		paramsArr = append(paramsArr, value)
	}
	paramsArr = append(paramsArr, c.Salt)
	sort.Strings(paramsArr)
	c.mu.Lock()
	defer func() {
		c.md5Hash.Reset()
		c.mu.Unlock()
	}()
	c.md5Hash.Write([]byte(strings.Join(paramsArr, "&")))

	return fmt.Sprintf("%x", c.md5Hash.Sum(nil)), nil
}

// POST 发起请求
func (c *Client) doPost(ctx context.Context, path string, bm gopay.BodyMap) (bs []byte, err error) {
	sign, err := c.getRsaSign(bm)
	if err != nil {
		return nil, fmt.Errorf("GetRsaSign Error: %w", err)
	}
	bm.Set("app_id", c.AppId)
	bm.Set("sign", sign)
	fmt.Println(sign)
	req := c.hc.Req()
	req.Header.Add("Accept", "application/json")
	if c.DebugSwitch == gopay.DebugOn {
		xlog.Debugf("DouYin_Req_Path: %s", path)
		xlog.Debugf("DouYin_Req_Body: %s", bm.JsonBody())
		xlog.Debugf("DouYin_Req_Headers: %#v", req.Header)
		fmt.Println(bm.JsonBody())
	}
	res, bs, err := req.Post(c.BaseURL + path).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}

// 验签示例
func DyCheckSign(timestamp, nonce, body, signature, pubKeyStr string) (bool, error) {
	pubKey, err := PemToRSAPublicKey(pubKeyStr) // 注意验签时publicKey使用平台公钥而非应用公钥
	if err != nil {
		return false, err
	}

	hashed := sha256.Sum256([]byte(timestamp + "\n" + nonce + "\n" + body + "\n"))
	signBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signBytes)
	return err == nil, nil
}

func PemToRSAPublicKey(pemKeyStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemKeyStr))
	if block == nil || len(block.Bytes) == 0 {
		return nil, fmt.Errorf("empty block in pem string")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch key := key.(type) {
	case *rsa.PublicKey:
		return key, nil
	default:
		return nil, fmt.Errorf("not rsa public key")
	}
}
