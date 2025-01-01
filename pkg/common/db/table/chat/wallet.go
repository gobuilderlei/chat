package chat

import (
	"context"
	"github.com/openimsdk/tools/db/pagination"
)

//钱包系统,包含 消费者,商家,公司账户及商家

// 存在问题:应该可以将这些写入一个表格
type Wallet struct {
	UserID        string  `json:"useId" bson:"user_id"`
	UserType      int     `json:"userType" bson:"user_type"`            //用户类型,1:消费者,2:商家,3:公司账户,4:商家 用户类别 1:个人,10,个人推荐人一级,11,个人推荐人二级,20,商家推荐人1级,21,商家推荐人二级,30,商户,40,股东县区,41,股东....,50,平台
	AllPoints     float32 `json:"allPoints" bson:"all_points"`          //总积分记账积分,数额不减少
	PointsKeeping float32 `json:"points_keeping" bson:"points_keeping"` //会减少的积分,变成抵扣券后玖就删除,达到365天后也减少积分
	Voucher       float32 `json:"voucher" bson:"voucher"`               //抵扣券
}

func (Wallet) TableName() string {
	return "wallet"
}

type WalletInterface interface {
	Create(ctx context.Context, wallet ...*Wallet) error
	Update(ctx context.Context, userId string, data map[string]any) error
	GetByUserID(ctx context.Context, userId string) (*Wallet, error)
	GetAllPointsKeepingBySystem(ctx context.Context, userType int) ([]*Wallet, float32, error)
	GetPaginationlPointsKeepingBySystem(ctx context.Context, userType int, pagination pagination.Pagination) (int64, []*Wallet, error)
}
