package chat

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	chatdb "github.com/openimsdk/chat/pkg/common/db/table/chat"
	"github.com/openimsdk/chat/pkg/protocol/chat"
	"github.com/robfig/cron/v3"
	"math"
	"strconv"
	"time"
)

//用户钱包,用户交易,用户订单,用户积分每日0点刷新等等操作.
//func (o *chatSvr) UserOrder (ctx context.Context, req *UserOrderReq) (*UserOrderResp, error) {

func (o *chatSvr) GetProducts(ctx context.Context, req *chat.UserIDOrUUIdAndPagination) (res *chat.ProductList, err error) {
	fmt.Println("GetProducts", req.Useridoruuid, req.GetPagination())
	if req.Useridoruuid != "" {
		aa, products, err := o.Database.GetProducts(ctx, req.Useridoruuid, req.GetPagination())
		if err != nil {
			return nil, nil
		}
		fmt.Println(aa)
		var producttss []*chat.ProductInfo
		var product *chat.ProductInfo
		for index, v := range products {
			fmt.Println(index, v)
			product.UUid = v.UUid
			product.BrandId = int32(v.BrandId)
			product.ProductCategoryId = int32(v.ProductCategoryId)
			product.FreightTemplateId = int32(v.FreightTemplateId)
			product.ProductAttributeId = int32(v.ProductAttributeId)
			product.Name = v.Name
			product.Pic = v.Pic
			product.ProductSn = v.ProductSn
			product.DeleteStatus = int32(v.DeleteStatus)
			product.PublishStatus = int32(v.PublishStatus)
			product.NewStatus = int32(v.NewStatus)
			product.RecommandStatus = int32(v.RecommandStatus)
			product.VerifyStatus = int32(v.VerifyStatus)
			product.Sort = int32(v.Sort)
			product.Sale = int32(v.Sale)
			product.Price = v.Price
			product.GiftGrowth = int32(v.GiftGrowth)
			product.GiftPoint = int32(v.GiftPoint)
			product.UsePointLimit = int32(v.UsePointLimit)
			product.SubTitle = v.SubTitle
			product.OriginalPrice = v.OriginalPrice
			product.Stock = int32(v.Stock)
			product.LowStock = int32(v.LowStock)
			product.Unit = v.Unit
			product.Weight = v.Weight
			product.PreviewStatus = int32(v.PreviewStatus)
			product.ServiceIds = v.ServiceIds
			product.Keywords = v.Keywords
			product.Note = v.Note
			product.AlbumPics = v.AlbumPics
			product.DetailTitle = v.DetailTitle
			product.DetailDesc = v.DetailDesc
			product.DetailHtml = v.DetailHtml
			product.PromotionStartTime = v.PromotionStartTime
			product.PromotionEndTime = v.PromotionEndTime
			product.PromotionPerLimit = int32(v.PromotionPerLimit)
			product.PromotionType = int32(v.PromotionType)
			product.BrandName = v.BrandName
			product.ProductCategoryName = v.ProductCategoryName
			product.Description = v.Description
			producttss = append(producttss, product)
		}
		return &chat.ProductList{ProductList: producttss}, nil
	}
	return nil, nil
}
func (o *chatSvr) GetProductForUUID(ctx context.Context, req *chat.UserIDOrUUId) (res *chat.ProductInfo, err error) {
	fmt.Println("GetProductForUUID", req.Useridoruuid)
	if req.Useridoruuid != "" {
		product, err := o.Database.GetProductForuuid(ctx, req.Useridoruuid)
		if err != nil {
			return nil, nil
		}
		return &chat.ProductInfo{
			UUid:                product.UUid,
			BrandId:             int32(product.BrandId),
			ProductCategoryId:   int32(product.ProductCategoryId),
			FreightTemplateId:   int32(product.FreightTemplateId),
			ProductAttributeId:  int32(product.ProductAttributeId),
			Name:                product.Name,
			Pic:                 product.Pic,
			ProductSn:           product.ProductSn,
			DeleteStatus:        int32(product.DeleteStatus),
			PublishStatus:       int32(product.PublishStatus),
			NewStatus:           int32(product.NewStatus),
			RecommandStatus:     int32(product.RecommandStatus),
			VerifyStatus:        int32(product.VerifyStatus),
			Sort:                int32(product.Sort),
			Sale:                int32(product.Sale),
			Price:               product.Price,
			GiftGrowth:          int32(product.GiftGrowth),
			GiftPoint:           int32(product.GiftPoint),
			UsePointLimit:       int32(product.UsePointLimit),
			SubTitle:            product.SubTitle,
			OriginalPrice:       product.OriginalPrice,
			Stock:               int32(product.Stock),
			LowStock:            int32(product.LowStock),
			Unit:                product.Unit,
			Weight:              product.Weight,
			PreviewStatus:       int32(product.PreviewStatus),
			ServiceIds:          product.ServiceIds,
			Keywords:            product.Keywords,
			Note:                product.Note,
			AlbumPics:           product.AlbumPics,
			DetailTitle:         product.DetailTitle,
			DetailDesc:          product.DetailDesc,
			DetailHtml:          product.DetailHtml,
			PromotionStartTime:  product.PromotionStartTime,
			PromotionEndTime:    product.PromotionEndTime,
			PromotionPerLimit:   int32(product.PromotionPerLimit),
			PromotionType:       int32(product.PromotionType),
			BrandName:           product.BrandName,
			ProductCategoryName: product.ProductCategoryName,
			Description:         product.Description,
		}, nil
	}
	return nil, nil
}
func (o *chatSvr) CreateProduct(ctx context.Context, req *chat.ProductInfo) (res *chat.ChatIsOk, err error) {
	var product chatdb.ProductAbttri
	product.UUid = req.UUid
	fmt.Println("UUID", req.UUid)
	product.BrandId = int(req.BrandId)
	product.ProductCategoryId = int(req.ProductCategoryId)
	product.FreightTemplateId = int(req.FreightTemplateId)
	product.ProductAttributeId = int(req.ProductAttributeId)
	product.Name = req.Name
	product.Pic = req.Pic
	product.ProductSn = req.ProductSn
	product.DeleteStatus = int(req.DeleteStatus)
	product.PublishStatus = int(req.PublishStatus)
	product.NewStatus = int(req.NewStatus)
	product.RecommandStatus = int(req.RecommandStatus)
	product.VerifyStatus = int(req.VerifyStatus)
	product.Sort = int(req.Sort)
	product.Sale = int(req.Sale)
	product.Price = req.Price
	product.GiftGrowth = int(req.GiftGrowth)
	product.GiftPoint = int(req.GiftPoint)
	product.UsePointLimit = int(req.UsePointLimit)
	product.SubTitle = req.SubTitle
	product.OriginalPrice = req.OriginalPrice
	product.Stock = int(req.Stock)
	product.LowStock = int(req.LowStock)
	product.Unit = req.Unit
	product.Weight = req.Weight
	product.PreviewStatus = int(req.PreviewStatus)
	product.ServiceIds = req.ServiceIds
	product.Keywords = req.Keywords
	product.Note = req.Note
	product.AlbumPics = req.AlbumPics
	product.DetailTitle = req.DetailTitle
	product.DetailDesc = req.DetailDesc
	product.DetailHtml = req.DetailHtml
	product.PromotionStartTime = req.PromotionStartTime
	product.PromotionEndTime = req.PromotionEndTime
	product.PromotionPerLimit = int(req.PromotionPerLimit)
	product.PromotionType = int(req.PromotionType)
	product.BrandName = req.BrandName
	product.ProductCategoryName = req.ProductCategoryName
	product.Description = req.Description
	err = o.Database.CreateProduct(ctx, &product)
	if err != nil {
		return &chat.ChatIsOk{IsOk: false}, nil
	}
	return &chat.ChatIsOk{IsOk: true}, nil
}
func (o *chatSvr) UpdateProduct(ctx context.Context, req *chat.UpdateDataReq) (res *chat.ChatIsOk, err error) {
	fmt.Println("UpdateProduct", req.Data)
	value := ""
	for _, v := range req.GetData() {
		value = v
		fmt.Println(value)
	}
	var reslut map[string]any
	err0 := json.Unmarshal([]byte(value), &reslut)
	if err0 != nil {
		fmt.Println("Unmarshal err:", err0)
	}
	err1 := o.Database.UpdateProduct(ctx, req.Useridoruuid, reslut)
	if err1 != nil {
		fmt.Println("UpdateProduct err:", err1)
		return &chat.ChatIsOk{
			IsOk: false,
		}, err1
	}
	return &chat.ChatIsOk{
		IsOk: true,
	}, err1
}

// 订单部分
func (o *chatSvr) CreateShopOrder(ctx context.Context, req *chat.ShopOrder) (res *chat.ChatIsOk, err error) {
	var shopOrder chatdb.ShopOrder
	var payment chatdb.PayAmount
	payment.AlipayAmount = float32(req.PayAmount.AlipayAmount)
	payment.WechatAmount = float32(req.PayAmount.WechatAmount)
	payment.VoucherAmount = float32(req.PayAmount.VoucherAmount)
	payment.UnionpayAmount = float32(req.PayAmount.UnionpayAmount)
	shopOrder.UUId = req.UUId                //     string    `json:"uuid" bson:"uuid"`
	shopOrder.UserId = req.UserId            //     string    `json:"userId" bson:"user_id"`
	shopOrder.OrderType = int(req.OrderType) //     int       `json:"orderType" bson:"order_type"`
	shopOrder.GoodsId = req.GoodsId          //     string    `json:"goodsId" bson:"goods_id"`
	shopOrder.MerchantId = req.MerchantId    //    string    `json:"merchantId" bson:"merchant_id"`
	shopOrder.Status = int(req.Status)       //     int       `json:"status" bson:"status"`
	shopOrder.Amount = float32(req.Amount)   //     float32   `json:"amount" bson:"amount"`
	shopOrder.PayAmount = payment
	shopOrder.ProfitRate = int(req.ProfitRate)
	magrate := (payment.AlipayAmount + payment.WechatAmount + payment.VoucherAmount + payment.UnionpayAmount) * float32(req.ProfitRate) / 100
	shopOrder.MarginRate = float32(math.Trunc(float64(magrate*100)) / 100) // PayAmount `json:"payAmount" bson:"pay_amount"`
	shopOrder.CreateTime = req.CreateTime                                  // int64     `json:"createTime" bson:"create_time"`
	shopOrder.PayTime = req.PayTime                                        //     int64     `json:"payTime" bson:"pay_time"`
	shopOrder.FinishTime = req.FinishTime                                  //    int64     `json:"finishTime" bson:"finish_time"`

	lastencry, err0 := o.Database.GetByUserIdForLAST(ctx, req.UserId)
	if err0 != nil {
		fmt.Println("GetByUserIdForLAST err:", err0)
	}
	shopOrder.Encryption = HmacSha256ToHex("ZWL",
		req.UserId+
			fmt.Sprintf("%v", req.CreateTime)+
			fmt.Sprintf("%v", req.FinishTime)+
			fmt.Sprintf("%d", req.Amount)+
			lastencry.Encryption,
	)
	err1 := o.Database.CreateOrder(ctx, req.UserId, &shopOrder)
	if err1 != nil {
		fmt.Println("CreateOrder err:", err1)
		return &chat.ChatIsOk{
			IsOk: false,
		}, err1
	}
	return &chat.ChatIsOk{IsOk: true}, nil
}
func (o *chatSvr) GetShopOrderForUserUUid(ctx context.Context, req *chat.UserIDOrUUId) (res *chat.ShopOrder, err error) {
	oderuuid, err := o.Database.GetOrderForuuid(ctx, req.Useridoruuid)
	if err != nil {
		fmt.Println("GetOrderForuuid err:", err)
		return nil, err
	}
	return &chat.ShopOrder{
		UUId:       oderuuid.UUId,
		UserId:     oderuuid.UserId,
		OrderType:  int32(oderuuid.OrderType),
		GoodsId:    oderuuid.GoodsId,
		MerchantId: oderuuid.MerchantId,
		Status:     int32(oderuuid.Status),
		Amount:     float64(oderuuid.Amount),
		PayAmount: &chat.PayAmount{
			AlipayAmount:   float64(oderuuid.PayAmount.AlipayAmount),
			WechatAmount:   float64(oderuuid.PayAmount.WechatAmount),
			VoucherAmount:  float64(oderuuid.PayAmount.VoucherAmount),
			UnionpayAmount: float64(oderuuid.PayAmount.UnionpayAmount),
		},
		ProfitRate: int32(oderuuid.ProfitRate),
		CreateTime: oderuuid.CreateTime,
		PayTime:    oderuuid.PayTime,
		FinishTime: oderuuid.FinishTime,
		Encryption: oderuuid.Encryption,
	}, nil
}

// 可以是id也可以是uuid,也可以是merchantid或者是goodsid
func (o *chatSvr) GetShopOrders(ctx context.Context, req *chat.UserIDOrUUIdAndPagination) (res *chat.ShopOrderListRes, err error) {
	_, orders, err := o.Database.GetOrderForUserid(ctx, req.Useridoruuid, req.GetPagination())
	if err != nil {
		fmt.Println("GetOrderForUserid err:", err)
		return nil, err
	}
	var orderlist []*chat.ShopOrder
	var order chat.ShopOrder
	for _, v := range orders {
		order.UUId = v.UUId
		order.UserId = v.UserId
		order.OrderType = int32(v.OrderType)
		order.GoodsId = v.GoodsId
		order.MerchantId = v.MerchantId
		order.Status = int32(v.Status)
		order.Amount = float64(v.Amount)
		order.PayAmount = &chat.PayAmount{
			AlipayAmount:   float64(v.PayAmount.AlipayAmount),
			WechatAmount:   float64(v.PayAmount.WechatAmount),
			VoucherAmount:  float64(v.PayAmount.VoucherAmount),
			UnionpayAmount: float64(v.PayAmount.UnionpayAmount),
		}
		order.ProfitRate = int32(v.ProfitRate)
		order.CreateTime = v.CreateTime
		order.PayTime = v.PayTime
		order.FinishTime = v.FinishTime
		order.Encryption = v.Encryption
		orderlist = append(orderlist, &order)
	}
	return &chat.ShopOrderListRes{
		ShopOrderList: orderlist,
	}, nil
}
func (o *chatSvr) GetShopOrderForStatus(ctx context.Context, req *chat.ShopOrderStatus) (res *chat.ShopOrderListRes, err error) {
	_, orders, err := o.Database.GetOrderForStatus(ctx, int(req.OrderType), int(req.Status), req.GetPagination())
	if err != nil {
		fmt.Println("GetOrderForStatus err:", err)
		return nil, err
	}
	var order chat.ShopOrder
	var orderlist []*chat.ShopOrder
	for _, v := range orders {
		order.UUId = v.UUId
		order.UserId = v.UserId
		order.OrderType = int32(v.OrderType)
		order.GoodsId = v.GoodsId
		order.MerchantId = v.MerchantId
		order.Status = int32(v.Status)
		order.Amount = float64(v.Amount)
		order.PayAmount = &chat.PayAmount{
			AlipayAmount:   float64(v.PayAmount.AlipayAmount),
			WechatAmount:   float64(v.PayAmount.WechatAmount),
			VoucherAmount:  float64(v.PayAmount.VoucherAmount),
			UnionpayAmount: float64(v.PayAmount.UnionpayAmount),
		}
		order.ProfitRate = int32(v.ProfitRate)
		order.CreateTime = v.CreateTime
		order.PayTime = v.PayTime
		order.FinishTime = v.FinishTime
		order.Encryption = v.Encryption
		orderlist = append(orderlist, &order)
	}
	return &chat.ShopOrderListRes{
		ShopOrderList: orderlist,
	}, nil
}

func (o *chatSvr) GetShopOrderForAmout(ctx context.Context, req *chat.ShopOrderAmount) (res *chat.ShopOrderListRes, err error) {
	_, orders, err := o.Database.GetOrderForAmount(ctx, float32(req.Minamount), float32(req.Maxamount), req.GetPagination())
	if err != nil {
		fmt.Println("GetOrderForAmout err:", err)
		return nil, err
	}
	var order chat.ShopOrder
	var orderlist []*chat.ShopOrder
	for _, v := range orders {
		order.UUId = v.UUId
		order.UserId = v.UserId
		order.OrderType = int32(v.OrderType)
		order.GoodsId = v.GoodsId
		order.MerchantId = v.MerchantId
		order.Status = int32(v.Status)
		order.Amount = float64(v.Amount)
		order.PayAmount = &chat.PayAmount{
			AlipayAmount:   float64(v.PayAmount.AlipayAmount),
			WechatAmount:   float64(v.PayAmount.WechatAmount),
			VoucherAmount:  float64(v.PayAmount.VoucherAmount),
			UnionpayAmount: float64(v.PayAmount.UnionpayAmount),
		}
		order.ProfitRate = int32(v.ProfitRate)
		order.CreateTime = v.CreateTime
		order.PayTime = v.PayTime
		order.FinishTime = v.FinishTime
		order.Encryption = v.Encryption
		orderlist = append(orderlist, &order)
	}
	return &chat.ShopOrderListRes{
		ShopOrderList: orderlist,
	}, nil
}

// 积分自动刷新系统
func (o *chatSvr) CreatePointAutoRefresh(ctx context.Context, req *chat.PointAutoRefresh) (res *chat.ChatIsOk, err error) {
	var pointAutoRefresh *chatdb.PointsRefreshRecord
	pointAutoRefresh, err0 := o.Database.GetPointsRefreshRecordForLast(ctx, req.UserId)
	if err0 != nil {
		fmt.Println("GetPointsRefreshRecordForLast err:", err0)
	}
	pointAutoRefresh.UserID = req.UserId
	pointAutoRefresh.TotalPoints = req.TotalPoints
	pointAutoRefresh.Operator = int(req.Operator)
	pointAutoRefresh.RefreshTime = req.RefreshTime
	pointAutoRefresh.Points = float32(req.Points)
	pointAutoRefresh.RefreshVoucher = float32(req.RefreshVoucher)
	pointAutoRefresh.Note = req.Note
	pointAutoRefresh.Encryption = HmacSha256ToHex("ZWL",
		req.UserId+
			fmt.Sprintf("%v", req.TotalPoints)+
			fmt.Sprintf("%v", req.Operator)+
			fmt.Sprintf("%f", req.Points)+
			fmt.Sprintf("%f", req.RefreshVoucher)+
			pointAutoRefresh.Encryption,
	)
	err1 := o.Database.CreatePointsRefreshRecord(ctx, pointAutoRefresh)
	if err1 != nil {
		fmt.Println("CreatePointAutoRefresh err:", err1)
		return &chat.ChatIsOk{IsOk: false}, nil
	}
	return &chat.ChatIsOk{IsOk: true}, nil
}

func (o *chatSvr) GetPointAutoRefresh(ctx context.Context, req *chat.UserIDOrUUIdAndPagination) (res *chat.PointsAutoRefreshListRes, err error) {
	_, pointAutoRefreshList, err := o.Database.GetPointsRefreshRecord(ctx, req.Useridoruuid, req.GetPagination())
	if err != nil {
		fmt.Println("GetPointAutoRefresh err:", err)
		return nil, err
	}
	var pointAutoRefresh chat.PointAutoRefresh
	var pointAutoRefreshListRes []*chat.PointAutoRefresh
	for _, v := range pointAutoRefreshList {
		pointAutoRefresh.UserId = v.UserID
		pointAutoRefresh.TotalPoints = v.TotalPoints
		pointAutoRefresh.Operator = int32(int(v.Operator))
		pointAutoRefresh.RefreshTime = v.RefreshTime
		pointAutoRefresh.Points = float64(v.Points)
		pointAutoRefresh.RefreshVoucher = float64(v.RefreshVoucher)
		pointAutoRefresh.Note = v.Note
		pointAutoRefresh.Encryption = v.Encryption
		pointAutoRefreshListRes = append(pointAutoRefreshListRes, &pointAutoRefresh)
		return &chat.PointsAutoRefreshListRes{
			PointAutoRefreshList: pointAutoRefreshListRes}, nil
	}
	return nil, nil
}

// 钱包
func (o *chatSvr) GetWallet(ctx context.Context, req *chat.UserIDOrUUId) (res *chat.Wallet, err error) {
	wallet, err := o.Database.GetWalletByUserID(ctx, req.Useridoruuid)
	if err != nil {
		fmt.Println("GetWallet err:", err)
		return nil, err
	}
	return &chat.Wallet{
		UserId:        wallet.UserID,
		PointsKeeping: float64(wallet.PointsKeeping),
		Voucher:       float64(wallet.Voucher),
	}, nil
}

func (o *chatSvr) UpdateWallet(ctx context.Context, req *chat.UpdateDataReq) (res *chat.ChatIsOk, err error) {
	value := ""
	for _, v := range req.GetData() {
		value = v
		fmt.Println(value)
	}
	var reslut map[string]any
	err1 := json.Unmarshal([]byte(value), &reslut)
	if err1 != nil {
		fmt.Println("UpdateWallrt err:", err1)
	}
	isok, err2 := o.Database.UpdateWallet(ctx, req.Useridoruuid, reslut)
	if err2 != nil {
		fmt.Println("UpdateWallrt err:", err2)
		return &chat.ChatIsOk{IsOk: isok}, nil
	}
	return &chat.ChatIsOk{IsOk: isok}, nil
}

func (o *chatSvr) UpdateWalletBysystem(ctx context.Context, req *chat.UserIDOrUUId) (res *chat.ChatIsOk, err error) {
	if req.Useridoruuid != "9999999999" {
		//1.先查询昨天总利润
		//1.1设置定时器,每天夜间23点59分59秒执行一次.
		corn := cron.New(cron.WithSeconds())
		corn.AddFunc("59 59 23 * * *", func() {
			timemicro := time.Now().Unix() - 86399
			fmt.Println("定时任务开始执行", timemicro)
			//1.2获取当前系统时间.然后给与系统去查询这段时间的总订单及总利润.
			allOrderMarginRate, err := o.Database.GetOrderMarginRateForSystem(ctx, timemicro)
			if err != nil {
				fmt.Println("GetOrderMarginRateForSystem err:", err)
				o.UpdateWalletBysystem(ctx, req)
				return
			}
			/*当前存在的bug:现在数据库少,可以一次性获取所有数据,然后再对数据库进行批量操作,
			但是现在数据库太大,一次性获取数据会导致内存溢出,所以需要分批获取数据,然后进行批量操作.
			*/
			allocated_OrderMarginRate := allOrderMarginRate * 0.8 //商户及个人可分配利润
			ltd_OrderMarginRate := allOrderMarginRate * 0.1       //企业利润
			fmt.Println("今日商户及个人可分配利润:", allocated_OrderMarginRate)
			fmt.Println("今日平台公司利润:", ltd_OrderMarginRate)
			allconsumerOrder, allconsumer_points, err := o.Database.GetWalletBySystem(ctx, req.Useridoruuid, 0)
			if err != nil {
				fmt.Println("allconsumer_points err:", err)
			}
			fmt.Println(allconsumer_points)
			allmerchantOrder, allmerchant_points, err := o.Database.GetWalletBySystem(ctx, req.Useridoruuid, 1)
			if err != nil {
				fmt.Println("allmerchant_points err:", err)
			}
			fmt.Println(allmerchant_points)
			//point_radix_string :=fmt.Sprintf("%.2f", all_points / allocated_OrderMarginRate)
			//计算每一个积分的价值.
			consumer_point_radix := truncateFloat32(allconsumer_points/allocated_OrderMarginRate*0.6, 2)
			fmt.Println("个人积分分红基数", consumer_point_radix)
			//var shop_consumerOrder []chatdb.Wallet
			//var shop_consumer chatdb.Wallet
			for _, v := range allconsumerOrder {
				//shop_consumer.UserID = v.UserID
				today_consumer_points := consumer_point_radix * v.PointsKeeping
				isok, err2 := o.Database.UpdateWallet(ctx, v.UserID,
					map[string]any{
						"points_keeping": today_consumer_points,
						"voucher":        v.Voucher + today_consumer_points})
				if err2 != nil {
					fmt.Println("个人用户失败!UpdateWallrt err:", err2, "当前账号", v.UserID, "更新失败", "points_keeping", v.PointsKeeping-today_consumer_points, "voucher", v.Voucher+today_consumer_points)
				}
				cu, err := o.Database.GetPointsRefreshRecordForLast(ctx, v.UserID)
				if err != nil {
					fmt.Println("GetPointsRefreshRecordForLast err:", err)
				}
				err1 := o.Database.CreatePointsRefreshRecord(ctx, &chatdb.PointsRefreshRecord{
					UserID:         v.UserID,
					TotalPoints:    toint64(v.AllPoints),
					Operator:       0,
					RefreshTime:    time.Now().Unix(),
					Points:         today_consumer_points,
					RefreshVoucher: today_consumer_points,
					Note:           "系统给予个人积分" + fmt.Sprintf("%s", today_consumer_points),
					Encryption: HmacSha256ToHex("ZWL",
						v.UserID+
							fmt.Sprintf("%v", v.AllPoints)+
							fmt.Sprintf("%v", 0)+
							fmt.Sprintf("%f", today_consumer_points)+
							fmt.Sprintf("%f", today_consumer_points)+
							cu.Encryption)})
				fmt.Println("个人用户成功!UpdateWallrt err:", err1)
				fmt.Println(isok)
			}
			merchant_point_radix := truncateFloat32(allmerchant_points/allocated_OrderMarginRate*0.4, 2)
			fmt.Println("商户积分分红基数", merchant_point_radix)
			for _, v := range allmerchantOrder {
				merchant_points := merchant_point_radix * v.PointsKeeping
				isok, err2 := o.Database.UpdateWallet(ctx, v.UserID, map[string]any{
					"points_keeping": v.PointsKeeping - merchant_points,
					"voucher":        v.Voucher + merchant_points,
				})
				if err2 != nil {
					fmt.Println("商户失败!UpdateWallrt err:", err2, "当前账号", v.UserID, "更新失败", "points_keeping", v.PointsKeeping-merchant_points, "voucher", v.Voucher+merchant_points)
				}
				cu, err := o.Database.GetPointsRefreshRecordForLast(ctx, v.UserID)
				if err != nil {
					fmt.Println("GetPointsRefreshRecordForLast err:", err)
				}
				err1 := o.Database.CreatePointsRefreshRecord(ctx, &chatdb.PointsRefreshRecord{
					UserID:         v.UserID,
					TotalPoints:    toint64(v.AllPoints),
					Operator:       0,
					RefreshTime:    time.Now().Unix(),
					Points:         merchant_point_radix,
					RefreshVoucher: merchant_point_radix,
					Note:           "系统给予商户积分" + fmt.Sprintf("%s", merchant_point_radix),
					Encryption: HmacSha256ToHex("ZWL",
						v.UserID+
							fmt.Sprintf("%v", v.AllPoints)+
							fmt.Sprintf("%v", 0)+
							fmt.Sprintf("%f", merchant_point_radix)+
							fmt.Sprintf("%f", merchant_point_radix)+
							cu.Encryption)})
				fmt.Println("商户积分成功!UpdateWallrt err:", err1)
				fmt.Println(isok)
			}

		})
		corn.Start()
		//2.总利润减去20%.
		//3.查询需要昨天内所有剩余的积分数
		//4.个人消费池子总利润/个人消费池总剩余积分数=个人积分价值,商户消费池子总利润/商户消费池剩余积分数=商户积分价值 计算出当前积分价值
		//5.个人积分价值*个人剩余积分=个人今天获得的抵扣券,商户积分价值*商户剩余积分=商户今天获得的抵扣券
		//6.更新个人钱包与商户钱包.个人剩余积分减去今天获得的抵扣券, 个人消费池子剩余积分减去今天获得的抵扣券, 商户剩余积分减去今天获得的抵扣券.
	}
	return &chat.ChatIsOk{IsOk: false}, nil
}

func HmacSha256ToHex(key string, data string) string {
	return hex.EncodeToString(HmacSha256(key, data))
}

func HmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

func truncateFloat32(value float32, places int) float32 {
	// 计算乘以的因子
	multiplier := float32(1)
	for i := 0; i < places; i++ {
		multiplier *= 10
	}

	// 先乘以因子，然后取整，再除以因子，进行截断
	return float32(int32(value*multiplier)) / multiplier
}

func toint64(value string) int64 {
	intvalue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println(err)
		toint64(value)
	}
	return intvalue
}
