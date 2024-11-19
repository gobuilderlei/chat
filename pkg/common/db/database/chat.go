// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package database

import (
	"context"
	"errors"
	"time"

	"github.com/openimsdk/tools/db/mongoutil"
	"github.com/openimsdk/tools/db/pagination"
	"github.com/openimsdk/tools/db/tx"

	"github.com/openimsdk/chat/pkg/common/constant"
	admindb "github.com/openimsdk/chat/pkg/common/db/model/admin"
	"github.com/openimsdk/chat/pkg/common/db/model/chat"
	"github.com/openimsdk/chat/pkg/common/db/table/admin"
	chatdb "github.com/openimsdk/chat/pkg/common/db/table/chat"
)

type ChatDatabaseInterface interface {
	GetUser(ctx context.Context, userID string) (account *chatdb.Account, err error)
	UpdateUseInfo(ctx context.Context, userID string, attribute map[string]any, updateCred, delCred []*chatdb.Credential) (err error)
	FindAttribute(ctx context.Context, userIDs []string) ([]*chatdb.Attribute, error)
	FindAttributeByAccount(ctx context.Context, accounts []string) ([]*chatdb.Attribute, error)
	TakeAttributeByPhone(ctx context.Context, areaCode string, phoneNumber string) (*chatdb.Attribute, error)
	TakeAttributeByEmail(ctx context.Context, Email string) (*chatdb.Attribute, error)
	TakeAttributeByAccount(ctx context.Context, account string) (*chatdb.Attribute, error)
	TakeAttributeByUserID(ctx context.Context, userID string) (*chatdb.Attribute, error)
	TakeAccount(ctx context.Context, userID string) (*chatdb.Account, error)
	TakeCredentialByAccount(ctx context.Context, account string) (*chatdb.Credential, error)
	TakeCredentialsByUserID(ctx context.Context, userID string) ([]*chatdb.Credential, error)
	TakeLastVerifyCode(ctx context.Context, account string) (*chatdb.VerifyCode, error)
	Search(ctx context.Context, normalUser int32, keyword string, gender int32, pagination pagination.Pagination) (int64, []*chatdb.Attribute, error)
	SearchUser(ctx context.Context, keyword string, userIDs []string, genders []int32, pagination pagination.Pagination) (int64, []*chatdb.Attribute, error)
	CountVerifyCodeRange(ctx context.Context, account string, start time.Time, end time.Time) (int64, error)
	AddVerifyCode(ctx context.Context, verifyCode *chatdb.VerifyCode, fn func() error) error
	UpdateVerifyCodeIncrCount(ctx context.Context, id string) error
	DelVerifyCode(ctx context.Context, id string) error
	RegisterUser(ctx context.Context, register *chatdb.Register, account *chatdb.Account, attribute *chatdb.Attribute, credentials []*chatdb.Credential) error
	LoginRecord(ctx context.Context, record *chatdb.UserLoginRecord, verifyCodeID *string) error
	UpdatePassword(ctx context.Context, userID string, password string) error
	UpdatePasswordAndDeleteVerifyCode(ctx context.Context, userID string, password string, codeID string) error
	NewUserCountTotal(ctx context.Context, before *time.Time) (int64, error)
	UserLoginCountTotal(ctx context.Context, before *time.Time) (int64, error)
	UserLoginCountRangeEverydayTotal(ctx context.Context, start *time.Time, end *time.Time) (map[string]int64, int64, error)
	DelUserAccount(ctx context.Context, userIDs []string) error

	//商品部分
	GetProducts(ctx context.Context, userid string, pagination pagination.Pagination) (aa int64, products []*chatdb.ProductAbttri, err error)
	GetProductForuuid(ctx context.Context, uuid string) (*chatdb.ProductAbttri, error)
	//GetProductsForuseid(ctx context.Context, uuid string, pagination pagination.Pagination) (int64, []*chatdb.ProductAbttri, error)
	CreateProduct(ctx context.Context, product ...*chatdb.ProductAbttri) error
	UpdateProduct(ctx context.Context, uuid string, data map[string]any) error
	//购物车部分

	//订单部分
	CreateOrder(ctx context.Context, userid string, order ...*chatdb.ShopOrder) error
	GetOrders(ctx context.Context, Userid string, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error)
	GetOrderForuuid(ctx context.Context, uuid string) (*chatdb.ShopOrder, error)
	GetOrderForUserid(ctx context.Context, userid string, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error)
	GetByUserIdForLAST(ctx context.Context, userid string) (*chatdb.ShopOrder, error)
	GetOrderForMerchantId(ctx context.Context, merchantid string, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error)
	GetOrderForStatus(ctx context.Context, ordertype, status int, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error)
	GetOrderForGoodsId(ctx context.Context, goodsId string, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error)
	GetOrderForAmount(ctx context.Context, minAmount, maxAmount float32, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error)
	//积分操作部分
	CreatePointsRefreshRecord(ctx context.Context, record ...*chatdb.PointsRefreshRecord) error
	GetPointsRefreshRecord(ctx context.Context, userid string, pagination pagination.Pagination) (int64, []*chatdb.PointsRefreshRecord, error)

	//钱包部分
	CreateWallet(ctx context.Context, wallet ...*chatdb.Wallet) error
	GetWalletByUserID(ctx context.Context, userid string) (*chatdb.Wallet, error)
	UpdateWallet(ctx context.Context, userId string, data map[string]any) (bool, error)
}

func NewChatDatabase(cli *mongoutil.Client) (ChatDatabaseInterface, error) {
	register, err := chat.NewRegister(cli.GetDB())
	if err != nil {
		return nil, err
	}
	account, err := chat.NewAccount(cli.GetDB())
	if err != nil {
		return nil, err
	}
	attribute, err := chat.NewAttribute(cli.GetDB())
	if err != nil {
		return nil, err
	}
	credential, err := chat.NewCredential(cli.GetDB())
	if err != nil {
		return nil, err
	}
	userLoginRecord, err := chat.NewUserLoginRecord(cli.GetDB())
	if err != nil {
		return nil, err
	}
	verifyCode, err := chat.NewVerifyCode(cli.GetDB())
	if err != nil {
		return nil, err
	}
	forbiddenAccount, err := admindb.NewForbiddenAccount(cli.GetDB())
	if err != nil {
		return nil, err
	}
	goods, err := chat.NewGoods(cli.GetDB())
	if err != nil {
		return nil, err
	}
	order, err := chat.NewOrder(cli.GetDB())
	if err != nil {
		return nil, err
	}
	points, err := chat.NewPoints(cli.GetDB())
	if err != nil {
		return nil, err
	}
	wallet, err := chat.NewWallet(cli.GetDB())
	if err != nil {
		return nil, err
	}
	subject, err := chat.NewSbuject(cli.GetDB())
	if err != nil {
		return nil, err
	}
	return &ChatDatabase{
		tx:               cli.GetTx(),
		register:         register,
		account:          account,
		attribute:        attribute,
		credential:       credential,
		userLoginRecord:  userLoginRecord,
		verifyCode:       verifyCode,
		forbiddenAccount: forbiddenAccount,
		goods:            goods,
		order:            order,
		points:           points,
		wallet:           wallet,
		subject:          subject,
	}, nil
}

type ChatDatabase struct {
	tx               tx.Tx
	register         chatdb.RegisterInterface
	account          chatdb.AccountInterface
	attribute        chatdb.AttributeInterface
	credential       chatdb.CredentialInterface
	userLoginRecord  chatdb.UserLoginRecordInterface
	verifyCode       chatdb.VerifyCodeInterface
	goods            chatdb.ProductInterface
	order            chatdb.ShopOrderInterface
	points           chatdb.PointsRefreshRecordInterface
	subject          chatdb.SubjectProductInterface
	wallet           chatdb.WalletInterface
	forbiddenAccount admin.ForbiddenAccountInterface
}

func (o *ChatDatabase) GetUser(ctx context.Context, userID string) (account *chatdb.Account, err error) {
	return o.account.Take(ctx, userID)
}

func (o *ChatDatabase) UpdateUseInfo(ctx context.Context, userID string, attribute map[string]any, updateCred, delCred []*chatdb.Credential) (err error) {
	return o.tx.Transaction(ctx, func(ctx context.Context) error {
		if err = o.attribute.Update(ctx, userID, attribute); err != nil {
			return err
		}
		for _, credential := range updateCred {
			if err = o.credential.CreateOrUpdateAccount(ctx, credential); err != nil {
				return err
			}
		}
		if err = o.credential.DeleteByUserIDType(ctx, delCred...); err != nil {
			return err
		}
		return nil
	})
}

func (o *ChatDatabase) FindAttribute(ctx context.Context, userIDs []string) ([]*chatdb.Attribute, error) {
	return o.attribute.Find(ctx, userIDs)
}

func (o *ChatDatabase) FindAttributeByAccount(ctx context.Context, accounts []string) ([]*chatdb.Attribute, error) {
	return o.attribute.FindAccount(ctx, accounts)
}

func (o *ChatDatabase) TakeAttributeByPhone(ctx context.Context, areaCode string, phoneNumber string) (*chatdb.Attribute, error) {
	return o.attribute.TakePhone(ctx, areaCode, phoneNumber)
}

func (o *ChatDatabase) TakeAttributeByEmail(ctx context.Context, email string) (*chatdb.Attribute, error) {
	return o.attribute.TakeEmail(ctx, email)
}

func (o *ChatDatabase) TakeAttributeByAccount(ctx context.Context, account string) (*chatdb.Attribute, error) {
	return o.attribute.TakeAccount(ctx, account)
}

func (o *ChatDatabase) TakeAttributeByUserID(ctx context.Context, userID string) (*chatdb.Attribute, error) {
	return o.attribute.Take(ctx, userID)
}

func (o *ChatDatabase) TakeLastVerifyCode(ctx context.Context, account string) (*chatdb.VerifyCode, error) {
	return o.verifyCode.TakeLast(ctx, account)
}

func (o *ChatDatabase) TakeAccount(ctx context.Context, userID string) (*chatdb.Account, error) {
	return o.account.Take(ctx, userID)
}

func (o *ChatDatabase) TakeCredentialByAccount(ctx context.Context, account string) (*chatdb.Credential, error) {
	return o.credential.TakeAccount(ctx, account)
}

func (o *ChatDatabase) TakeCredentialsByUserID(ctx context.Context, userID string) ([]*chatdb.Credential, error) {
	return o.credential.Find(ctx, userID)
}

func (o *ChatDatabase) Search(ctx context.Context, normalUser int32, keyword string, genders int32, pagination pagination.Pagination) (total int64, attributes []*chatdb.Attribute, err error) {
	var forbiddenIDs []string
	if int(normalUser) == constant.NormalUser {
		forbiddenIDs, err = o.forbiddenAccount.FindAllIDs(ctx)
		if err != nil {
			return 0, nil, err
		}
	}
	total, totalUser, err := o.attribute.SearchNormalUser(ctx, keyword, forbiddenIDs, genders, pagination)
	if err != nil {
		return 0, nil, err
	}
	return total, totalUser, nil
}

func (o *ChatDatabase) SearchUser(ctx context.Context, keyword string, userIDs []string, genders []int32, pagination pagination.Pagination) (int64, []*chatdb.Attribute, error) {
	return o.attribute.SearchUser(ctx, keyword, userIDs, genders, pagination)
}

func (o *ChatDatabase) CountVerifyCodeRange(ctx context.Context, account string, start time.Time, end time.Time) (int64, error) {
	return o.verifyCode.RangeNum(ctx, account, start, end)
}

func (o *ChatDatabase) AddVerifyCode(ctx context.Context, verifyCode *chatdb.VerifyCode, fn func() error) error {
	return o.tx.Transaction(ctx, func(ctx context.Context) error {
		if err := o.verifyCode.Add(ctx, []*chatdb.VerifyCode{verifyCode}); err != nil {
			return err
		}
		if fn != nil {
			return fn()
		}
		return nil
	})
}

func (o *ChatDatabase) UpdateVerifyCodeIncrCount(ctx context.Context, id string) error {
	return o.verifyCode.Incr(ctx, id)
}

func (o *ChatDatabase) DelVerifyCode(ctx context.Context, id string) error {
	return o.verifyCode.Delete(ctx, id)
}

func (o *ChatDatabase) RegisterUser(ctx context.Context, register *chatdb.Register, account *chatdb.Account, attribute *chatdb.Attribute, credentials []*chatdb.Credential) error {
	return o.tx.Transaction(ctx, func(ctx context.Context) error {
		if err := o.register.Create(ctx, register); err != nil {
			return err
		}
		if err := o.account.Create(ctx, account); err != nil {
			return err
		}
		if err := o.attribute.Create(ctx, attribute); err != nil {
			return err
		}
		if err := o.credential.Create(ctx, credentials...); err != nil {
			return err
		}
		return nil
	})
}

func (o *ChatDatabase) LoginRecord(ctx context.Context, record *chatdb.UserLoginRecord, verifyCodeID *string) error {
	return o.tx.Transaction(ctx, func(ctx context.Context) error {
		if err := o.userLoginRecord.Create(ctx, record); err != nil {
			return err
		}
		if verifyCodeID != nil {
			if err := o.verifyCode.Delete(ctx, *verifyCodeID); err != nil {
				return err
			}
		}
		return nil
	})
}

func (o *ChatDatabase) UpdatePassword(ctx context.Context, userID string, password string) error {
	return o.account.UpdatePassword(ctx, userID, password)
}

func (o *ChatDatabase) UpdatePasswordAndDeleteVerifyCode(ctx context.Context, userID string, password string, codeID string) error {
	return o.tx.Transaction(ctx, func(ctx context.Context) error {
		if err := o.account.UpdatePassword(ctx, userID, password); err != nil {
			return err
		}
		if codeID == "" {
			return nil
		}
		if err := o.verifyCode.Delete(ctx, codeID); err != nil {
			return err
		}
		return nil
	})
}

func (o *ChatDatabase) NewUserCountTotal(ctx context.Context, before *time.Time) (int64, error) {
	return o.register.CountTotal(ctx, before)
}

func (o *ChatDatabase) UserLoginCountTotal(ctx context.Context, before *time.Time) (int64, error) {
	return o.userLoginRecord.CountTotal(ctx, before)
}

func (o *ChatDatabase) UserLoginCountRangeEverydayTotal(ctx context.Context, start *time.Time, end *time.Time) (map[string]int64, int64, error) {
	return o.userLoginRecord.CountRangeEverydayTotal(ctx, start, end)
}

func (o *ChatDatabase) DelUserAccount(ctx context.Context, userIDs []string) error {
	return o.tx.Transaction(ctx, func(ctx context.Context) error {
		if err := o.register.Delete(ctx, userIDs); err != nil {
			return err
		}
		if err := o.account.Delete(ctx, userIDs); err != nil {
			return err
		}
		if err := o.attribute.Delete(ctx, userIDs); err != nil {
			return err
		}
		return nil
	})
}

// 商品部分
func (o *ChatDatabase) GetProducts(ctx context.Context, userid string, pagination pagination.Pagination) (aa int64, products []*chatdb.ProductAbttri, err error) {
	return o.goods.GetProducts(ctx, userid, pagination)
}
func (o *ChatDatabase) GetProductForuuid(ctx context.Context, uuid string) (*chatdb.ProductAbttri, error) {
	return o.goods.GetProduct(ctx, uuid)
}

//	func (o *ChatDatabase) GetProductsForuuid(ctx context.Context, uuid string, pagination pagination.Pagination) (int64, []*chatdb.ProductAbttri, error) {
//		return o.goods.GetProductsForuuid(ctx, uuid, pagination)
//	}
func (o *ChatDatabase) CreateProduct(ctx context.Context, product ...*chatdb.ProductAbttri) error {
	return o.goods.CreateProduct(ctx, product...)
}
func (o *ChatDatabase) UpdateProduct(ctx context.Context, uuid string, data map[string]any) error {
	return o.goods.UpdateProduct(ctx, uuid, data)
}

//购物车部分

// 订单部分
func (o *ChatDatabase) CreateOrder(ctx context.Context, userid string, order ...*chatdb.ShopOrder) error {
	if len(order) == 0 {
		return errors.New("order is nil")
	}
	return o.order.Create(ctx, userid, order...)
}
func (o *ChatDatabase) GetOrders(ctx context.Context, Userid string, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error) {
	return 0, nil, nil
}
func (o *ChatDatabase) GetOrderForuuid(ctx context.Context, uuid string) (*chatdb.ShopOrder, error) {
	return nil, nil
}
func (o *ChatDatabase) GetOrderForUserid(ctx context.Context, userid string, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error) {
	return o.order.GetByUserId(ctx, userid, pagination)
}
func (o *ChatDatabase) GetByUserIdForLAST(ctx context.Context, userid string) (*chatdb.ShopOrder, error) {
	return o.order.GetByUserIdForLast(ctx, userid)
}
func (o *ChatDatabase) GetOrderForMerchantId(ctx context.Context, merchantid string, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error) {
	return o.order.GetByMerchantId(ctx, merchantid, pagination)
}
func (o *ChatDatabase) GetOrderForStatus(ctx context.Context, ordertype, status int, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error) {
	return o.order.GetByStatus(ctx, ordertype, status, pagination)
}
func (o *ChatDatabase) GetOrderForGoodsId(ctx context.Context, goodsId string, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error) {
	return o.order.GetByGoodsId(ctx, goodsId, pagination)
}
func (o *ChatDatabase) GetOrderForAmount(ctx context.Context, minAmount, maxAmount float32, pagination pagination.Pagination) (int64, []*chatdb.ShopOrder, error) {
	return o.order.GetByAmount(ctx, minAmount, maxAmount, pagination)
}

// 积分操作部分
func (o *ChatDatabase) CreatePointsRefreshRecord(ctx context.Context, record ...*chatdb.PointsRefreshRecord) error {
	return o.points.Create(ctx, record...)
}
func (o *ChatDatabase) GetPointsRefreshRecord(ctx context.Context, userid string, pagination pagination.Pagination) (int64, []*chatdb.PointsRefreshRecord, error) {
	return o.points.Take(ctx, userid, pagination)
}

// 钱包部分
func (o *ChatDatabase) CreateWallet(ctx context.Context, wallet ...*chatdb.Wallet) error {
	return o.wallet.Create(ctx, wallet...)
}
func (o *ChatDatabase) GetWalletByUserID(ctx context.Context, userid string) (*chatdb.Wallet, error) {
	return o.wallet.GetByUserID(ctx, userid)
}
func (o *ChatDatabase) UpdateWallet(ctx context.Context, userId string, data map[string]any) (bool, error) {
	if o.wallet.Update(ctx, userId, data) != error(nil) {
		return false, error(nil)
	}
	return true, o.wallet.Update(ctx, userId, data)
}
