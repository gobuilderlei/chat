package chat

import (
	"context"
	shopdb "github.com/openimsdk/chat/pkg/common/db/table/chat"
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

func (w *Wallet) Create(ctx context.Context, wallet *shopdb.Wallet) error {
	return nil
}
func (w *Wallet) Update(ctx context.Context, userId string, data map[string]any) (bool, error) {
	return false, nil
}
func (w *Wallet) GetByUserID(ctx context.Context, userId string) (*shopdb.Wallet, error) {
	return nil, nil
}
