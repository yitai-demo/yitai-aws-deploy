package model

type Medal struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Gift struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Coupon struct {
	Id                 string `json:"id"`
	GiftId             string `json:"giftId"`
	GiftName           string `json:"giftName"`
	OwnerId            string `json:"ownerId"`
	CollectedAtUnixSec int64  `json:"collectedAtUnixSec"`
	Collected          bool   `json:"collected"`
	Source             int    `json:"source"`
	Sequence           int    `json:"sequence"`
	RedeemedAtUnixSec  int64  `json:"redeemedAtUnixSec"`
	Redeemed           bool   `json:"redeemed"`
	ExtraInfo          string `json:"extraInfo"`
}

type Egg struct {
	Id                 string  `json:"id"`
	Type               string  `json:"type"`
	CouponId           string  `json:"couponId"`
	OwnerId            string  `json:"ownerId"`
	CollectedAtUnixSec int64   `json:"collectedAtUnixSec"`
	Collected          bool    `json:"collected"`
	BatchId            string  `json:"batchId"`
	ServerId           string  `json:"serverId"`
	MapId              string  `json:"mapId"`
	LocationX          float32 `json:"locationX"`
	LocationY          float32 `json:"locationY"`
}

// Medal
type CollectMedalReq struct {
	MedalName string `json:"medalName"`
}

type CollectMedalRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	IsSuccessful bool   `json:"data"`
}

// Coupon
type AccessSecretRoomReq struct {
	AccessKey string `json:"accessKey"`
}

type TakeLuckyDrawRsp struct {
	ErrorCode    int     `json:"errcode"`
	ErrorMessage string  `json:"errmsg"`
	Coupon       *Coupon `json:"data"`
}

type RedeemCouponReq struct {
	CouponId string `json:"couponId"`
}

type RedeemCouponRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	IsSuccessful bool   `json:"data"`
}

// Egg
type CollectEggReq struct {
	EggId string `json:"eggId"`
}

type CollectEggRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	Egg          *Egg   `json:"data"`
}

type GetUncollectedEggsRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	Eggs         []*Egg `json:"data"`
}

// Query
type GetMyBagSummaryData struct {
	Medals  []*Medal  `json:"medals"`
	Eggs    []*Egg    `json:"eggs"`
	Coupons []*Coupon `json:"coupons"`
	HasBell bool      `json:"hasBell"`
}

type GetMyBagSummaryRsp struct {
	ErrorCode    int                  `json:"errcode"`
	ErrorMessage string               `json:"errmsg"`
	Data         *GetMyBagSummaryData `json:"data"`
}

// --------------------
// Query
// --------------------
type CheckSecretRoomAccessRsp struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
	Data         bool   `json:"data"`
}

type MedalIssuedCount struct {
	Ranking     int    `json:"ranking"`
	Name        string `json:"name"`
	IssuedCount int    `json:"issuedCount"`
}

type GetMedalIssuedCountRsp struct {
	ErrorCode    int                 `json:"errcode"`
	ErrorMessage string              `json:"errmsg"`
	Data         []*MedalIssuedCount `json:"data"`
}

type UserTitleCount struct {
	Title string `json:"title"`
	Count int    `json:"count"`
}

type GetUserTitleCountRsp struct {
	ErrorCode    int               `json:"errcode"`
	ErrorMessage string            `json:"errmsg"`
	Data         []*UserTitleCount `json:"data"`
}
