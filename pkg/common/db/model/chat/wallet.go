package chat

import (
	"context"
	"fmt"
	"github.com/openimsdk/chat/pkg/common/db/dbutil"
	shopdb "github.com/openimsdk/chat/pkg/common/db/table/chat"
	"github.com/openimsdk/tools/db/mongoutil"
	"github.com/openimsdk/tools/db/pagination"
	"github.com/openimsdk/tools/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
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

// 通过userType来获取不同类型人员的钱包信息 具体参阅用户类型 用户类别 1:个人,10,个人推荐人一级,11,个人推荐人二级,20,商家推荐人1级,21,商家推荐人二级,30,商户,40,股东县区,41,股东....,50,平台
// 然后通过此来进行数据查询
func (w *Wallet) GetAllPointsKeepingBySystem(ctx context.Context, userType int) ([]*shopdb.Wallet, float32, error) {
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

	db_wallets := make(chan []*shopdb.Wallet)
	err_chan := make(chan error)
	var chatdbwallets []*shopdb.Wallet
	var TotalPointsKeeping float32
	//这里开始操作并发,然后进行数据计算
	go func() { //查询出来,
		var page int32 = 0
		var showNum int32 = 50
		_, cusor, err := mongoutil.FindPage[*shopdb.Wallet](ctx, w.coll, filter, &dbutil.MyPagination{PageNumber: page, ShowNumber: showNum})
		db_wallets <- cusor
		err_chan <- err

		time.Sleep(time.Millisecond * 200)
		page++

	}()
	select {
	case err := <-err_chan:
		if err != nil {
			close(db_wallets)
			close(err_chan)
			return nil, 0, errs.Wrap(err)
		}

	case wallets := <-db_wallets:
		for _, wallet := range wallets {
			fmt.Println(wallet.UserID)
			TotalPointsKeeping += wallet.PointsKeeping
		}
		chatdbwallets = append(chatdbwallets, wallets...)
		return chatdbwallets, TotalPointsKeeping, nil
	}
	return nil, 0, nil
	//cusor, err := w.coll.Find(ctx, filter)
	//if err != nil {
	//	fmt.Println(err)
	//	w.GetAllPointsKeepingBySystem(ctx, userType)
	//}
	//defer cusor.Close(ctx)
	//var TotalPointsKeeping float32
	//for cusor.Next(ctx) {
	//	var wallet shopdb.Wallet
	//	if err := cusor.Decode(&wallet); err != nil {
	//		return nil, 0, errs.Wrap(err)
	//	}
	//	TotalPointsKeeping += wallet.PointsKeeping
	//}
	//if err := cusor.Err(); err != nil {
	//	return nil, 0, errs.Wrap(err)
	//}
	//var wallets []shopdb.Wallet
	//err1 := cusor.All(ctx, &wallets)
	//if err1 != nil {
	//	return nil, 0, errs.Wrap(err1)
	//}
	//return wallets, TotalPointsKeeping, nil
}

// 接上面的内容 通过页面自动实现数据查询,然后进行并发操作.
func (w *Wallet) GetPaginationlPointsKeepingBySystem(ctx context.Context, userType int, pagination pagination.Pagination) (int64, []*shopdb.Wallet, error) {
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
	fmt.Println(filter)
	return mongoutil.FindPage[*shopdb.Wallet](ctx, w.coll, filter, pagination)
}

func (w *Wallet) UpdatePointsKeepingForSystem(ctx context.Context, userType int) (bool, error) {
	return false, nil
}
