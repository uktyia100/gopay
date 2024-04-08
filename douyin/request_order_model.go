package douyin

// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/api/industry/general_trade/create_order/requestOrder#87daf5bf

// ++++++ 下面是 tt.requestOrder 回调构造参数 ++++

type RequestOrderData struct {
	SkuList          []*RequestOrderSkuList `json:"skuList,omitempty"`          // 下单商品信息
	OutOrderNo       string                 `json:"outOrderNo,omitempty"`       // 外部订单号
	TotalAmount      int64                  `json:"totalAmount,omitempty"`      // 订单总金额
	PayExpireSeconds int                    `json:"payExpireSeconds,omitempty"` // 支付超时时间
	PayNotifyUrl     string                 `json:"payNotifyUrl,omitempty"`     // 支付结果通知地址
	MerchantUid      string                 `json:"merchantUid,omitempty"`      // 开发者自定义收款商户号
	OrderEntrySchema *RequestOrderSchema    `json:"orderEntrySchema,omitempty"` // 订单详情页
	LimitPayWayList  []int                  `json:"limitPayWayList,omitempty"`  // 屏蔽的支付方式
}

type RequestOrderSkuList struct {
	SkuId       string              `json:"skuId,omitempty"`    // 外部商品id
	Price       int64               `json:"price,omitempty"`    // 价格 单位：分
	Quantity    int                 `json:"quantity,omitempty"` // 购买数量 0 < quantity <= 100
	Title       string              `json:"title,omitempty"`
	ImageList   []string            `json:"imageList,omitempty"`
	Type        int                 `json:"type,omitempty"` // 根据接入规范，选择适合的商品类型ID传入
	TagGroupId  string              `json:"tagGroupId,omitempty"`
	EntrySchema *RequestOrderSchema `json:"entrySchema,omitempty"` // 商品详情页链接
}

// 商品详情页链接
type RequestOrderSchema struct {
	Path   string `json:"path,omitempty"`
	Params string `json:"params,omitempty"`
}

// ++++++ 下面是 tt.requestOrder 回调成功 ++++

type RequestOrderRespSuccess struct {
	OrderId       string           `json:"orderId,omitempty"` // 抖音开放平台内部的交易订单号
	ItemOrderList []*itemOrderList `json:"itemOrderList,omitempty"`
	LogId         string           `json:"logId,omitempty"`
}

type itemOrderList struct {
	ItemOrderId     string `json:"itemOrderId,omitempty"`
	SkuId           string `json:"skuId,omitempty"`
	ItemOrderAmount int    `json:"itemOrderAmount,omitempty"`
}

// ++++++ 下面是 tt.requestOrder 回调失败 ++++

type RequestOrderRespFail struct {
	ErrNo    string `json:"errNo,omitempty"`
	ErrMsg   string `json:"errMsg,omitempty"`
	ErrLogId string `json:"errLogId,omitempty"`
}
