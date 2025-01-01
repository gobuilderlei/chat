package chat

import (
	"context"
	"github.com/openimsdk/tools/db/pagination"
)

// 每日系统积分刷新记录表
// 前端自动根据当前订单情况后,将积分数据写入此表,每日系统将根据此表刷新用户的积分数据
type PointsRefreshRecord struct {
	UserID         string  `json:"user_id" bson:"user_id"`
	TotalPoints    float32 `json:"totalPoints" bson:"total_points"`       //总积分,不删除的
	Operator       int     `json:"operator" bson:"operator"`              //操作人  0:系统 1:用户
	RefreshTime    int64   `json:"refreshTime" bson:"refresh_time"`       //刷新时间戳
	Points         float32 `json:"points" bson:"points"`                  //保留小数点后2位数//数值直接截断,不四舍五入
	RefreshVoucher float32 `json:"refreshVoucher" bson:"refresh_voucher"` //刷新的抵扣券
	Note           string  `json:"note" bson:"note"`                      //备注
	Encryption     string  `json:"encryption" bson:"encryption"`          //hash加密//每人加密数值都是上次的加密后的数值
}

func (PointsRefreshRecord) TableName() string {
	return "points_refresh_record"
}

// 只有创建与查询,不提供更新修改与删除
type PointsRefreshRecordInterface interface {
	Create(ctx context.Context, record ...*PointsRefreshRecord) error
	Take(ctx context.Context, userID string, pagination pagination.Pagination) (int64, []*PointsRefreshRecord, error)
	TakeLast(ctx context.Context, userID string) (*PointsRefreshRecord, error)
}
