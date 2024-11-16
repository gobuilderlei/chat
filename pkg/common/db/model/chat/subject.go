package chat

import (
	"context"
	shopdb "github.com/openimsdk/chat/pkg/common/db/table/chat"
	"github.com/openimsdk/tools/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewSbuject(db *mongo.Database) (shopdb.SubjectProductInterface, error) {
	coll := db.Collection("shop_subjectProduct")
	_, err := coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{
			Key:   "uuid",
			Value: 1,
		}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &subjectProduct{coll: coll}, nil
}

type subjectProduct struct {
	coll *mongo.Collection
}

func (s *subjectProduct) Create(ctx context.Context, product *shopdb.SubjectProduct) error {
	return nil
}
func (s *subjectProduct) Take(ctx context.Context, UUid string) (*shopdb.SubjectProduct, error) {
	return nil, nil
}
func (s *subjectProduct) Update(ctx context.Context, UUid string, data map[string]any) error {
	return nil
}
