package chat

import (
	"context"
	"fmt"
	shopdb "github.com/openimsdk/chat/pkg/common/db/table/chat"
	"github.com/openimsdk/tools/db/mongoutil"
	"github.com/openimsdk/tools/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewWallet(db *mongo.Database) (shopdb.WalletInterface, error) {
	coll := db.Collection("wallet")
	_, err := coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{
			Key:   "user_id",
			Value: 1,
		}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &Wallet{coll: coll}, nil
}

type Wallet struct {
	coll *mongo.Collection
}

func (w *Wallet) Create(ctx context.Context, wallet ...*shopdb.Wallet) error {
	return mongoutil.InsertMany[*shopdb.Wallet](ctx, w.coll, wallet)
}
func (w *Wallet) Update(ctx context.Context, userId string, data map[string]any) error {
	return mongoutil.UpdateOne(ctx, w.coll, bson.M{"user_id": userId}, bson.M{"$set": data}, false)
}
func (w *Wallet) GetByUserID(ctx context.Context, userId string) (*shopdb.Wallet, error) {
	return mongoutil.FindOne[*shopdb.Wallet](ctx, w.coll, bson.M{"user_id": userId})
}
func (w *Wallet) GetAllPointsKeepingBySystem(ctx context.Context, userType int) ([]shopdb.Wallet, float32, error) {
	//pipeline := mongo.Pipeline{
	//	{{"$group", bson.D{
	//		{"_id", nil}, // 不分组，计算总和
	//		{"user_type", userType},
	//		{"totalPointsKeeping", bson.D{
	//			{"$sum", "$points_keeping"},
	//		}},
	//	}}},
	//}
	//cursor, err := w.coll.Aggregate(ctx, pipeline)
	//if err != nil {
	//	w.GetAllPointsKeepingBySystem(ctx, userType)
	//}
	//defer cursor.Close(ctx)
	//var reslut struct {
	//	TotalPointsKeeping float32 `json:"totalPointsKeeping" bson:"totalPointsKeeping"`
	//}
	//if cursor.Next(ctx) {
	//	err = cursor.Decode(&reslut)
	//	if err != nil {
	//		return 0, errs.Wrap(err)
	//	}
	//	return reslut.TotalPointsKeeping, nil
	//}
	var filter bson.M
	switch userType {
	case 0: //个人及商家推荐人
		filter = bson.M{
			"user_type": bson.M{
				"$gt": 0,
				"$lt": 30,
			},
		}
		break
	case 1: //
		filter = bson.M{
			"user_type": bson.M{
				"$gt": 29,
				"$lt": 50,
			},
		}
		break
	}
	cusor, err := w.coll.Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
		w.GetAllPointsKeepingBySystem(ctx, userType)
	}
	defer cusor.Close(ctx)
	var TotalPointsKeeping float32
	for cusor.Next(ctx) {
		var wallet shopdb.Wallet
		if err := cusor.Decode(&wallet); err != nil {
			return nil, 0, errs.Wrap(err)
		}
		TotalPointsKeeping += wallet.PointsKeeping
	}
	if err := cusor.Err(); err != nil {
		return nil, 0, errs.Wrap(err)
	}
	var wallets []shopdb.Wallet
	err1 := cusor.All(ctx, &wallets)
	if err1 != nil {
		return nil, 0, errs.Wrap(err1)
	}
	return wallets, TotalPointsKeeping, nil
}

func (w *Wallet) UpdatePointsKeepingForSystem(ctx context.Context, userType int) (bool, error) {
	return false, nil
}
