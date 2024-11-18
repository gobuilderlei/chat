package chat

import "context"

//钱包系统,包含 消费者,商家,公司账户及商家

// 存在问题:应该可以将这些写入一个表格
type Wallet struct {
	UserID        string  `json:"useId" bson:"user_id"`
	PointsKeeping float32 `json:"points_keeping" bson:"points_keeping"` //记账积分,数额不减少
	Voucher       float32 `json:"voucher" bson:"voucher"`               //抵扣券
}

func (Wallet) TableName() string {
	return "wallet"
}

type WalletInterface interface {
	Create(ctx context.Context, wallet ...*Wallet) error
	Update(ctx context.Context, userId string, data map[string]any) error
	GetByUserID(ctx context.Context, userId string) (*Wallet, error)
}
