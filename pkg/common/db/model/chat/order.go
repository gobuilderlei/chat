package chat

import (
	"context"
	shopoder "github.com/openimsdk/chat/pkg/common/db/table/chat"
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

func (s *Shop_Order) Create(ctx context.Context, order *shopoder.ShopOrder) error {
	return nil
}
func (s *Shop_Order) GetByUUId(ctx context.Context, UUId string) (*shopoder.ShopOrder, error) {
	return nil, nil
}
func (s *Shop_Order) GetByUserId(ctx context.Context, userId string) ([]*shopoder.ShopOrder, error) {
	return nil, nil
}
func (s *Shop_Order) GetByMerchantId(ctx context.Context, merchantId string) ([]*shopoder.ShopOrder, error) {
	return nil, nil
}
func (s *Shop_Order) GetByStatus(ctx context.Context, ordertype, status int) ([]*shopoder.ShopOrder, error) {
	return nil, nil
}
func (s *Shop_Order) GetByGoodsId(ctx context.Context, goodsId string) ([]*shopoder.ShopOrder, error) {
	return nil, nil
}
func (s *Shop_Order) GetByAmount(ctx context.Context, minAmount, maxAmount float64) ([]*shopoder.ShopOrder, error) {
	return nil, nil
}
