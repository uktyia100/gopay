package douyin

type NotifyRequest struct {
	Timestamp    string `json:"timestamp,omitempty"` // UTC时间戳
	Nonce        string `json:"nonce,omitempty"`     // 随机字符串
	Msg          string `json:"msg,omitempty"`
	MsgSignature string `json:"msg_signature,omitempty"`
	Type         string `json:"type,omitempty"`
}

type NotifyResp struct {
	ErrNo   int    `json:"err_no"`
	ErrTips string `json:"err_tips"`
}

type NotifyMsgResp struct {
	AppId          string `json:"appid,omitempty"`            // 当前交易发起的小程序id
	CpOrderno      string `json:"cp_orderno,omitempty"`       // 开发者侧的订单号
	CpExtra        string `json:"cp_extra,omitempty"`         // 预下单时开发者传入字段
	Way            string `json:"way,omitempty"`              // 支付渠道 1-微信支付，2-支付宝支付，10-抖音支付
	ChannelNo      string `json:"channel_no,omitempty"`       // 支付渠道侧单号
	PaymentOrderNo string `json:"payment_order_no,omitempty"` // 支付渠道侧PC单号
	TotalAmount    int    `json:"total_amount,omitempty"`     // 支付金额，单位为分
	Status         string `json:"status,omitempty"`           // 固定 SUCCESS
	ItemId         string `json:"item_id,omitempty"`          // 订单来源视频对应视频 id
	SellerUid      string `json:"seller_uid,omitempty"`       // 该笔交易卖家商户号
	PaidAt         int64  `json:"paid_at,omitempty"`          // 支付时间，Unix 时间戳
	OrderId        string `json:"order_id,omitempty"`         // 抖音侧订单号
}

type PushOrderResponse struct {
	ErrCode int64  `json:"err_code,omitempty"` // 执行结果
	ErrMsg  string `json:"err_msg,omitempty"`  // 返回错误信息
	Body    string `json:"data,omitempty"`     // POI 等关联业务推送结果，非 POI 订单为空，JSON 字符串
}

type AccessTokenReq struct {
	AppId     string `json:"appid"`
	Secret    string `json:"secret"`
	GrantType string `json:"grant_type"`
}

type AccessTokenResp struct {
	Data struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	} `json:"data"`
	ErrCode int    `json:"err_no"`
	ErrMsg  string `json:"err_tips"`
}
