package chat

import (
	"context"
	"github.com/openimsdk/tools/db/pagination"
)

//交易系统
//订单系统

type ShopOrder struct {
	UUId   string `json:"uuid" bson:"uuid"`      //订单编号
	userId string `json:"userId" bson:"user_id"` //用户编号
	//订单类型
	OrderType int    `json:"orderType" bson:"order_type"` //订单类型 0为线下订单,线下订单不提供退款服务,可以无商品编号,状态直接为已完成,1为线上订单
	goodsId   string `json:"goodsId" bson:"goods_id"`     //商品编号
	//商家编号
	MerchantId string `json:"merchantId" bson:"merchant_id"` //商家编号
	//支付类型
	PayType string `json:"payType" bson:"pay_type"` //支付类型
	//订单状态
	Status int `json:"status" bson:"status"` //订单状态 0为==待支付 1为==待发货 2为==待收货 3为==已完成 4为==已取消 5为==退款中 6为==退款完成
	//订单金额
	Amount float32 `json:"amount" bson:"amount"` //订单金额
	//支付金额情况
	PayAmount PayAmount `json:"payAmount" bson:"pay_amount"` //支付金额情况
	//订单创建时间
	CreateTime int64 `json:"createTime" bson:"create_time"` //订单创建时间
	//订单支付时间
	PayTime int64 `json:"payTime" bson:"pay_time"` //订单支付时间
	//订单完成时间
	FinishTime int64  `json:"finishTime" bson:"finish_time"` //订单完成时间
	Encryption string `json:"encryption" bson:"encryption"`  //hash加密 每个订单都有专属的加密串
}

type PayAmount struct {
	AlipayAmount   float32 `json:"alipayAmount" bson:"alipay_amount"`     //支付宝支付金额
	WechatAmount   float32 `json:"wechatAmount" bson:"wechat_amount"`     //微信支付金额
	UnionpayAmount float32 `json:"unionpayAmount" bson:"unionpay_amount"` //银行卡支付金额
	VoucherAmount  float32 `json:"voucherAmount" bson:"voucher_amount"`   //代金券支付金额
}

func (ShopOrder) TableName() string {
	return "shop_order"
}

type ShopOrderInterface interface {
	Create(ctx context.Context, order ...*ShopOrder) error
	GetByUUId(ctx context.Context, UUId string) (*ShopOrder, error)
	GetByUserId(ctx context.Context, userId string, pagination pagination.Pagination) (int64, []*ShopOrder, error)
	GetByMerchantId(ctx context.Context, merchantId string, pagination pagination.Pagination) (int64, []*ShopOrder, error)
	GetByStatus(ctx context.Context, ordertype, status int, pagination pagination.Pagination) (int64, []*ShopOrder, error)
	GetByGoodsId(ctx context.Context, goodsId string, pagination pagination.Pagination) (int64, []*ShopOrder, error)
	GetByAmount(ctx context.Context, minAmount, maxAmount float32, pagination pagination.Pagination) (int64, []*ShopOrder, error)
}
