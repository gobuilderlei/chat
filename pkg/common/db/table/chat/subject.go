package chat

import "context"

//专题商品

type SubjectProduct struct {
	UUid            string  `json:"uuid" bson:"uuid"`               //商品ID
	CategoryID      int64   `json:"category_id" bson:"category_id"` //分类ID
	Tile            string  `json:"tile" bson:"tile"`               //标题
	Pic             string  `json:"pic" bson:"pic"`
	ProductCount    int     `json:"productCount" bson:"productCount"`
	RecommendStatus int     `json:"recommendStatus" bson:"recommendStatus"` //推荐状态
	CollectCount    int     `json:"collectCount" bson:"collectCount"`
	ReadCount       int     `json:"readCount" bson:"readCount"`
	CommentCount    int     `json:"commentCount" bson:"commentCount"`
	Price           float64 `json:"price" bson:"price"` //价格
	AlbumPics       string  `json:"albumPics" bson:"albumPics"`
	Description     string  `json:"description" bson:"description"`   //描述
	ShowStatus      int     `json:"showStatus" bson:"showStatus"`     //展示状态 0:下架 1:上架 2:预览 3:审核中 4:审核失败
	Content         string  `json:"content" bson:"content"`           //内容
	ForwardCount    int     `json:"forwardCount" bson:"forwardCount"` //转发数
	CategoryName    string  `json:"categoryName" bson:"categoryName"`
}

func (SubjectProduct) TableName() string {
	return "shop_subjectProduct"
}

type SubjectProductInterface interface {
	Create(ctx context.Context, product *SubjectProduct) error
	Take(ctx context.Context, UUid string) (*SubjectProduct, error)
	Update(ctx context.Context, UUid string, data map[string]any) error
}
