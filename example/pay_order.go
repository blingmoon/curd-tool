package example

import "time"

type PayOrderPO struct {
	ID              int64      `gorm:"column:id" json:"id"`
	UID             string     `gorm:"column:uid" json:"uid"`
	OpenID          string     `gorm:"column:open_id" json:"open_id"`
	UserIP          string     `gorm:"column:user_ip" json:"user_ip"`
	AppID           string     `gorm:"column:app_id" json:"app_id"`
	Platform        string     `gorm:"column:platform" json:"platform"`
	Project         string     `gorm:"column:project" json:"project"`
	TradeNo         string     `gorm:"column:trade_no" json:"trade_no"`
	MID             string     `gorm:"column:mid" json:"mid"`
	PayType         string     `gorm:"column:pay_type" json:"pay_type"`
	GoodsID         int        `gorm:"column:goods_id" json:"goods_id"`
	GoodsName       string     `gorm:"column:goods_name" json:"goods_name"`
	GoodsType       int        `gorm:"column:goods_type" json:"goods_type"`
	PayAmount       int        `gorm:"column:pay_amount" json:"pay_amount"`
	RefundAmount    int        `gorm:"column:refund_amount" json:"refund_amount"`
	RefundTime      string     `gorm:"column:refund_time" json:"refund_time"`
	Status          int        `gorm:"column:status" json:"status"`
	Remark          string     `gorm:"column:remark" json:"remark"`
	DeviceType      string     `gorm:"column:device_type" json:"device_type"`
	PlayletID       int        `gorm:"column:playlet_id" json:"playlet_id"`
	LinkID          string     `gorm:"column:link_id" json:"link_id"`
	DistributorID   int        `gorm:"column:distributor_id" json:"distributor_id"`
	PromoterID      int        `gorm:"column:promoter_id" json:"promoter_id"`
	PromotionPlanID int        `gorm:"column:promotion_plan_id" json:"promotion_plan_id"`
	AdID            string     `gorm:"column:ad_id" json:"ad_id"`
	OrderTime       *time.Time `gorm:"column:order_time" json:"order_time"`
	PayTime         *time.Time `gorm:"column:pay_time" json:"pay_time"`
	PayDate         *time.Time `gorm:"column:pay_date" json:"pay_date"`
	CreateTime      *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime      *time.Time `gorm:"column:update_time" json:"update_time"`
	Extra           string     `gorm:"column:extra" json:"extra"`
}

func (PayOrderPO) TableName() string {
	return "pay_order"
}
