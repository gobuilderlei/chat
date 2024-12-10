package chat

import (
	"context"
	"fmt"
	shopdb "github.com/openimsdk/chat/pkg/common/db/table/chat"
	"github.com/openimsdk/tools/db/mongoutil"
	"github.com/openimsdk/tools/db/pagination"
	"github.com/openimsdk/tools/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewGoods(db *mongo.Database) (shopdb.ProductInterface, error) {
	coll := db.Collection("shop_product_abttri")
	_, err := coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{"uuid", 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &ShopProductAbttri{coll: coll}, nil
}

type ShopProductAbttri struct {
	coll *mongo.Collection
}

func (p *ShopProductAbttri) GetProduct(ctx context.Context, uuid string) (*shopdb.ProductAbttri, error) {
	if uuid == "" {
		return nil, nil
	}
	return mongoutil.FindOne[*shopdb.ProductAbttri](ctx, p.coll, bson.M{"uuid": uuid, "_id": -1}, nil)
}
func (p *ShopProductAbttri) GetProducts(ctx context.Context, uuid string, pagination pagination.Pagination) (int64, []*shopdb.ProductAbttri, error) {
	fmt.Println(uuid)
	filter := bson.M{}
	return mongoutil.FindPage[*shopdb.ProductAbttri](ctx, p.coll, filter, pagination, nil)
}

func (p *ShopProductAbttri) GetProductForuuid(ctx context.Context, uuid string) (*shopdb.ProductAbttri, error) {
	return mongoutil.FindOne[*shopdb.ProductAbttri](ctx, p.coll, bson.M{"uuid": uuid, "_id": -1}, nil)
}
func (p *ShopProductAbttri) GetProductsForuuid(ctx context.Context, uuid string, pagination pagination.Pagination) (int64, []*shopdb.ProductAbttri, error) {
	return mongoutil.FindPage[*shopdb.ProductAbttri](ctx, p.coll, bson.M{"uuid": uuid}, pagination, nil)
}

func (p *ShopProductAbttri) CreateProduct(ctx context.Context, product ...*shopdb.ProductAbttri) error {
	return mongoutil.InsertMany[*shopdb.ProductAbttri](ctx, p.coll, product)
}

func (p *ShopProductAbttri) UpdateProduct(ctx context.Context, uuid string, data map[string]any) error {
	if len(data) == 0 {
		return nil
	}
	filter := bson.M{"uuid": uuid}
	return mongoutil.UpdateOne(ctx, p.coll, filter, bson.M{"$set": data}, false)
	//return nil
}
