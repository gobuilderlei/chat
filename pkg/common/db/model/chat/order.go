package chat

import (
	"context"
	shopoder "github.com/openimsdk/chat/pkg/common/db/table/chat"
	"github.com/openimsdk/tools/db/mongoutil"
	"github.com/openimsdk/tools/db/pagination"
	"github.com/openimsdk/tools/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewOrder(db *mongo.Database) (shopoder.ShopOrderInterface, error) {
	coll := db.Collection("shop_order")
	_, err := coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "uuid", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &Shop_Order{coll: coll}, nil
}

type Shop_Order struct {
	coll *mongo.Collection
}

func (s *Shop_Order) Create(ctx context.Context, userId string, order ...*shopoder.ShopOrder) error {
	return mongoutil.InsertMany[*shopoder.ShopOrder](ctx, s.coll, order)
}
func (s *Shop_Order) GetByUUId(ctx context.Context, UUId string) (*shopoder.ShopOrder, error) {
	return mongoutil.FindOne[*shopoder.ShopOrder](ctx, s.coll, bson.M{"uuid": UUId})
}
func (s *Shop_Order) GetByUserId(ctx context.Context, userId string, pagination pagination.Pagination) (int64, []*shopoder.ShopOrder, error) {
	return mongoutil.FindPage[*shopoder.ShopOrder](ctx, s.coll, bson.M{"user_id": userId}, pagination)
}
func (s *Shop_Order) GetByUserIdForLast(ctx context.Context, userId string) (*shopoder.ShopOrder, error) {
	option := options.FindOne().SetSort(bson.D{{"_id", -1}})
	return mongoutil.FindOne[*shopoder.ShopOrder](ctx, s.coll, bson.M{"user_id": userId}, option)
}
func (s *Shop_Order) GetByMerchantId(ctx context.Context, merchantId string, pagination pagination.Pagination) (int64, []*shopoder.ShopOrder, error) {
	return mongoutil.FindPage[*shopoder.ShopOrder](ctx, s.coll, bson.M{"merchant_id": merchantId}, pagination)
}
func (s *Shop_Order) GetByStatus(ctx context.Context, ordertype, status int, pagination pagination.Pagination) (int64, []*shopoder.ShopOrder, error) {
	return mongoutil.FindPage[*shopoder.ShopOrder](ctx, s.coll, bson.M{"order_type": ordertype, "status": status}, pagination)
}
func (s *Shop_Order) GetByGoodsId(ctx context.Context, goodsId string, pagination pagination.Pagination) (int64, []*shopoder.ShopOrder, error) {
	return mongoutil.FindPage[*shopoder.ShopOrder](ctx, s.coll, bson.M{"goods_id": goodsId}, pagination)
}
func (s *Shop_Order) GetByAmount(ctx context.Context, minAmount, maxAmount float32, pagination pagination.Pagination) (int64, []*shopoder.ShopOrder, error) {
	filter := bson.M{
		"amount": bson.M{
			"$gte": minAmount, // 大于或等于 50
			"$lte": maxAmount, // 小于或等于 150
		},
	}
	return mongoutil.FindPage[*shopoder.ShopOrder](ctx, s.coll, filter, pagination)
}
