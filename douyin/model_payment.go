package douyin

type CreateOrderData struct {
	OrderId    string `json:"order_id,omitempty"` // 抖音侧的订单号
	OrderToken string `json:"order_token,omitempty"`
}

type CreateOrderResponse struct {
	ErrNo   int              `json:"err_no,omitempty"`   // 执行结果
	ErrTips string           `json:"err_tips,omitempty"` // 返回错误信息
	Data    *CreateOrderData `json:"data"`
}

type QueryOrderResponse struct {
	ErrNo       int              `json:"err_no,omitempty"`       // 执行结果
	ErrTips     string           `json:"err_tips,omitempty"`     // 返回错误信息
	OutOrderNo  string           `json:"out_order_no,omitempty"` // 开发者侧的订单号
	OrderId     string           `json:"order_id,omitempty"`     // 抖音侧的订单号
	PaymentInfo *PaymentInfoData `json:"payment_info,omitempty"` // 支付信息
	CpsInfo     *CpsInfoData     `json:"cps_info,omitempty"`     // cps信息
}

type PaymentInfoData struct {
	TotalFee    int    `json:"total_fee,omitempty"`    // 支付金额
	OrderStatus string `json:"order_status,omitempty"` // 支付状态枚举值
	PayTime     string `json:"pay_time,omitempty"`     // 支付完成时间
	Way         int    `json:"way,omitempty"`          // 支付渠道
	ChannelNo   string `json:"channel_no,omitempty"`   // 支付渠道侧的支付单号
	SellerUid   string `json:"seller_uid,omitempty"`   // 该笔交易卖家商户号
	ItemId      string `json:"item_id,omitempty"`      // 订单来源视频对应视频 id
	CpExtra     string `json:"cp_extra,omitempty"`     // 预下单时开发者传入字段
}

type CpsInfoData struct {
	TotalFee string `json:"share_amount,omitempty"` // 达人分佣金额
	DouyinId string `json:"douyin_id,omitempty"`    // 达人抖音号
	Nickname string `json:"nickname,omitempty"`     // 达人昵称
}
