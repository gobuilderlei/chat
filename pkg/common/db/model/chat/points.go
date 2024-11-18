package chat

import (
	"context"
	shopdb "github.com/openimsdk/chat/pkg/common/db/table/chat"
	"github.com/openimsdk/tools/db/mongoutil"
	"github.com/openimsdk/tools/db/pagination"
	"github.com/openimsdk/tools/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewPoints(db *mongo.Database) (shopdb.PointsRefreshRecordInterface, error) {
	coll := db.Collection("points_refresh_record")
	_, err := coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{"refresh_time", -1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return &PointsRe{coll: coll}, nil
}

type PointsRe struct {
	coll *mongo.Collection
}

func (p *PointsRe) Create(ctx context.Context, record ...*shopdb.PointsRefreshRecord) error {
	return mongoutil.InsertMany[*shopdb.PointsRefreshRecord](ctx, p.coll, record)
}
func (p *PointsRe) Take(ctx context.Context, userID string, pagination pagination.Pagination) (int64, []*shopdb.PointsRefreshRecord, error) {
	return mongoutil.FindPage[*shopdb.PointsRefreshRecord](ctx, p.coll, userID, pagination)
}
