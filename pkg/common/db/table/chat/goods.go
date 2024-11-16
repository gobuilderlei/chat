package chat

import (
	"context"
	"github.com/openimsdk/tools/db/pagination"
)

// 商品信息
// 品牌列表
type BrandList struct {
	ID                  int    `json:"brand_id" bson:"brand_id"`
	Name                string `json:"name" bson:"name"`
	FirstLetter         string `json:"firstLetter" bson:"firstLetter"`
	Sort                int    `json:"sort" bson:"sort"`
	FactoryStatus       string `json:"factoryStatus" bson:"factoryStatus"`
	ShowStatus          string `json:"showStatus" bson:"showStatus"`
	ProductCount        int    `json:"productCount" bson:"productCount"`
	ProductCommentCount int    `json:"productCommentCount" bson:"productCommentCount"`
	Logo                string `json:"logo" bson:"logo"`     //brand logo url
	BigPic              string `json:"bigPic" bson:"bigPic"` //brand big picture url
}

func (BrandList) TableName() string {
	return "shop_brand"
}

type BrandInterface interface {
	Create(ctx context.Context, brand ...*BrandList) error
	Take(ctx context.Context, Userid string, pagination pagination.Pagination) (int64, brandlist []*BrandList, err error)
	Updater(ctx context.Context, ID int, data map[string]any) error
	Delete(ctx context.Context, ID int) error
}

// 抢购商品:带有倒计时的商品列表
type HomeFlashPromotion struct {
	UUid              string  `json:"uuid" bson:"uuid"`
	StartTime         string  `json:"startTime" bson:"startTime"`
	EndTime           string  `json:"endTime" bson:"endTime"`
	NextStartTime     string  `json:"nextStartTime" bson:"nextStartTime"`
	NextEndTime       string  `json:"nextEndTime" bson:"nextEndTime"`
	ProductgoosIdList []int32 `json:"productgoosIdList" bson:"productgoosIdList"` //写入商品ID列表
	//ProductList   []int        `json:"productList" bson:"productList"`
}

func (HomeFlashPromotion) TableName() string {
	return "shop_home_flash_promotion"
}

type HomeFlashPromotionInterface interface {
	Create(ctx context.Context, home ...*HomeFlashPromotion) error
	Take(ctx context.Context, userid string, pagination pagination.Pagination) (int64, []*HomeFlashPromotion, error)
	Update(ctx context.Context, uuid string, data map[string]any) error
	Delete(ctx context.Context, uuid string) error
}

type ProductAbttri struct {
	//ProductID          int     `json:"product_id" bson:"product_id"`
	UUid string `json:"uuid" bson:"uuid"` //商品uuid
	//GoodsId            int     `json:"goodsId" bson:"goods_id"`                      //商品id
	BrandId             int     `json:"brandId" bson:"brandId"`                       //品牌id
	ProductCategoryId   int     `json:"productCategoryId" bson:"productCategoryId"`   //商品分类id
	FreightTemplateId   int     `json:"freightTemplateId" bson:"freightTemplateId"`   //运费模板id
	ProductAttributeId  int     `json:"productAttributeId" bson:"productAttributeId"` //商品属性id
	Name                string  `json:"name" bson:"name"`                             //商品名称
	Pic                 string  `json:"pic" bson:"pic"`                               //商品主图
	ProductSn           string  `json:"productSn" bson:"productSn"`                   //商品编号
	DeleteStatus        int     `json:"deleteStatus" bson:"deleteStatus"`             //删除状态 0为正常,1为删除
	PublishStatus       int     `json:"publishStatus" bson:"publishStatus"`           //发布状态 0为下架,1为上架
	NewStatus           int     `json:"newStatus" bson:"newStatus"`                   //是否是新平 1为新品 0为不是新品 2为精品 3为热销 4为抢购
	RecommandStatus     int     `json:"recommandStatus" bson:"recommandStatus"`       //推荐状态 0为不推荐 1为推荐
	VerifyStatus        int     `json:"verifyStatus" bson:"verifyStatus"`             //审核状态 0为未审核 1为审核通过 2为审核中 3为审核失败
	Sort                int     `json:"sort" bson:"sort"`                             //排序
	Sale                int     `json:"sale" bson:"sale"`                             //销量
	Price               float64 `json:"price" bson:"price"`                           //价格
	GiftGrowth          int     `json:"giftGrowth" bson:"giftGrowth"`                 //赠送成长值
	GiftPoint           int     `json:"giftPoint" bson:"giftPoint"`                   //赠送积分
	UsePointLimit       int     `json:"usePointLimit" bson:"usePointLimit"`           //限制积分使用
	SubTitle            string  `json:"subTitle" bson:"subTitle"`                     //副标题
	OriginalPrice       float64 `json:"originalPrice" bson:"originalPrice"`           //原价
	Stock               int     `json:"stock" bson:"stock"`                           //库存
	LowStock            int     `json:"lowStock" bson:"lowStock"`                     //库存预警值
	Unit                string  `json:"unit" bson:"unit"`                             //商品单位
	Weight              float64 `json:"weight" bson:"weight"`                         //商品重量
	PreviewStatus       int     `json:"previewStatus" bson:"previewStatus"`           //是否为预览商品 0为不是 1为是
	ServiceIds          string  `json:"serviceIds" bson:"serviceIds"`                 //服务项目ids
	Keywords            string  `json:"keywords" bson:"keywords"`                     //关键字
	Note                string  `json:"note" bson:"note"`                             //备注
	AlbumPics           string  `json:"albumPics" bson:"albumPics"`                   //图册图片
	DetailTitle         string  `json:"detailTitle" bson:"detailTitle"`               //详情标题
	DetailDesc          string  `json:"detailDesc" bson:"detailDesc"`
	DetailHtml          string  `json:"detailHtml" bson:"detailHtml"`
	PromotionStartTime  string  `json:"promotionStartTime" bson:"promotionStartTime"`   //促销开始时间
	PromotionEndTime    string  `json:"promotionEndTime" bson:"promotionEndTime"`       //促销结束时间
	PromotionPerLimit   int     `json:"promotionPerLimit" bson:"promotionPerLimit"`     //每人限购数量 0 		不限制 其他数值为具体限制人数
	PromotionType       int     `json:"promotionType" bson:"promotionType"`             //促销类型 1为打折 2为满减 3为阶梯折扣 4为满折 5为满赠 6为指定折扣 7为限时折扣 8为会员价 9为阶梯价格 10为满减送 11为满折送 12为满赠送 13为指定折扣送 14为限时折扣送 15为会员价送
	BrandName           string  `json:"brandName" bson:"brandName"`                     //品牌名称
	ProductCategoryName string  `json:"productCategoryName" bson:"productCategoryName"` //商品分类名称
	Description         string  `json:"description" bson:"description"`                 //商品描述
}

func (ProductAbttri) TableName() string {
	return "shop_product_abttri"
}

type ProductInterface interface {
	GetProduct(ctx context.Context, uuid string) (*ProductAbttri, error)
	GetProducts(ctx context.Context, userID string, pagination pagination.Pagination) (int64, []*ProductAbttri, error)
	GetProductForuuid(ctx context.Context, uuid string) (*ProductAbttri, error)
	GetProductsForuuid(ctx context.Context, uuid string, pagination pagination.Pagination) (int64, []*ProductAbttri, error)
	CreateProduct(ctx context.Context, product ...*ProductAbttri) error
	UpdateProduct(ctx context.Context, uuid string, data map[string]any) error
}
